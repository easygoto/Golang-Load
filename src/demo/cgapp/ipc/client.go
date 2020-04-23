package ipc

import (
	"encoding/json"
)

type Client struct {
	serverConn chan string
}

func NewIpcClient(server *Server) *Client {
	conn := server.Connect()
	return &Client{serverConn: conn}
}

func (client *Client) Call(method, params string) (response *Response, err error) {
	request := &Request{method, params}
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	client.serverConn <- string(jsonBytes)
	str := <-client.serverConn
	_ = json.Unmarshal([]byte(str), &response)
	return
}

func (client *Client) Close() {
	client.serverConn <- "CLOSE"
}
