package study

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

// 哈希算法
func TestString(t *testing.T) {
	testString := "123123"
	md5Inst := md5.New()
	_, _ = md5Inst.Write([]byte(testString))
	result := md5Inst.Sum([]byte(""))
	_, _ = fmt.Printf("%x\n", result)

	sha1Inst := sha1.New()
	_, _ = sha1Inst.Write([]byte(testString))
	result = sha1Inst.Sum([]byte(""))
	_, _ = fmt.Printf("%x\n", result)

	sha256Inst := sha256.New()
	_, _ = sha256Inst.Write([]byte(testString))
	result = sha256Inst.Sum([]byte(""))
	_, _ = fmt.Printf("%x\n", result)

	testFile := "./test.txt"
	file, _ := os.Open(testFile)
	md5h := md5.New()
	_, _ = io.Copy(md5h, file)
	_, _ = fmt.Printf("%x %s\n", md5h.Sum([]byte("")), testFile)

	sha1h := sha1.New()
	_, _ = io.Copy(sha1h, file)
	_, _ = fmt.Printf("%x %s\n", sha1h.Sum([]byte("")), testFile)
}

// 数字证书
func TestCert(t *testing.T) {
	max := new(big.Int).Lsh(big.NewInt(1), 128) // max := 1 << 128
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Demo Co."},
		OrganizationalUnit: []string{"test"},
		CommonName:         "Go lang Test",
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		// KeyUsage 与 ExtKeyUsage 用来表明该证书是用作服务器认证的
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	// 生成一对具有指定字位数的 RSA 密钥
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // DER 格式
	certFile, _ := os.Create("cert.pem")
	_ = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	_ = certFile.Close()
	defer certFile.Close()
	keyFile, _ := os.Create("key.pem")
	_ = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	_ = keyFile.Close()
	defer keyFile.Close()
}

// RSA公钥私钥产生
func TestGenRsaKey(t *testing.T) {
	pk, _ := rsa.GenerateKey(rand.Reader, 1024)
	derBytes := x509.MarshalPKCS1PublicKey(&pk.PublicKey)
	pubFile, _ := os.Create("public.pem")
	_ = pem.Encode(pubFile, &pem.Block{Type: "PUBLIC KEY", Bytes: derBytes})
	defer pubFile.Close()

	pvtFile, _ := os.Create("private.pem")
	_ = pem.Encode(pvtFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	defer pvtFile.Close()
}

// Https
func TestHttpsServer(t *testing.T) {
	http.HandleFunc("/", rootHandler)
	_ = myListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
}

// Sftp
func TestSftpServer(t *testing.T) {
	handler := http.FileServer(http.Dir("."))
	_ = http.ListenAndServeTLS(":8022", "cert.pem", "key.pem", handler)
}

// Echo Server
func TestEchoServer(t *testing.T) {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
	service := "127.0.0.1:7999"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}

// Echo Client
func TestEchoClient(t *testing.T) {
	conn, err := tls.Dial("tcp", "127.0.0.1:7999", nil)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
	state := conn.ConnectionState()
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)
	message := "Hello\n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", message, n)
	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}

func myListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) (err error) {
	config := &tls.Config{
		Rand:       rand.Reader,
		Time:       time.Now,
		NextProtos: []string{"http/1.1"},
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = myLoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	tlsListener := tls.NewListener(conn, config)
	return http.Serve(tlsListener, handler)
}

func myLoadX509KeyPair(certFile string, keyFile string) (cert tls.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return
	}
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = errors.New("crypto/tls: failed to parse certificate PEM data")
		return
	}
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain == nil {
		cert.Certificate = [][]byte{certDERBlock.Bytes}
	} else {
		cert.Certificate = [][]byte{certDERBlock.Bytes, certDERBlockChain.Bytes}
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		err = errors.New("crypto/tls: failed to parse key PEM data")
		return
	}
	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		err = errors.New("crypto/tls: failed to parse key")
		return
	}
	cert.PrivateKey = key
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}
	if x509Cert.PublicKeyAlgorithm != x509.RSA ||
		x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N) != 0 {
		err = errors.New("crypto/tls: private key does not match public key")
		return
	}
	return
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	msg := "Welcome"
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(msg)))
	_, _ = w.Write([]byte(msg))
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read: %s", err)
			}
			break
		}
		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)
		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}
