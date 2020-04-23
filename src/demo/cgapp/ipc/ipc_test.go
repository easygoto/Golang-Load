package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
	BaseServer
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{Body: fmt.Sprintf("method:%s,params:%s", method, params)}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("Client1", "params1")
	resp2, _ := client1.Call("Client2", "params2")
	if resp1.Body != "method:Client1,params:params1" || resp2.Body != "method:Client2,params:params2" {
		t.Error("Client. Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
	}
	client1.Close()
	client2.Close()
}
