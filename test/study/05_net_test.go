package study

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
)

// TCP server
func Test(t *testing.T) {
	listener, _ := net.Listen("tcp", ":7800")
	for {
		conn0, _ := listener.Accept()

		go func(conn net.Conn) {
			defer conn.Close()
			_, _ = fmt.Println("client address:", conn.RemoteAddr())
			buffer := make([]byte, 1024)
			recvLen, _ := conn.Read(buffer)
			_, _ = conn.Write([]byte("I am server, you message : " + string(buffer[:recvLen])))
			_, _ = fmt.Println("send message success")
		}(conn0)
	}
}

// 使用 tcp 写报文
func TestSocket(t *testing.T) {
	service := "www.example.com:80"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	_, _ = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	result, _ := ioutil.ReadAll(conn)
	_, _ = fmt.Println(string(result))
}

// 开启一个 http 服务
func TestHttpServer(t *testing.T) {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// 测试 http get 方法
func TestHttpClient(t *testing.T) {
	http.Client.Do()
	response, _ := http.Get("https://www.example.com/")
	defer response.Body.Close()
	_, _ = io.Copy(os.Stdout, response.Body)
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"msg":  "Welcome",
		"name": request.FormValue("name"),
		"url":  html.EscapeString(request.RequestURI)})
	_, _ = fmt.Fprintf(writer, string(jsonBytes))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
