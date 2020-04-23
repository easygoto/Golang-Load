package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type BaseServer interface {
	Name() string
	Handle(method, params string) *Response
}

type Server struct {
	BaseServer
}

func NewIpcServer(server BaseServer) *Server {
	return &Server{server}
}

// 服务端连接之后把这个连接暴露出去
func (server *Server) Connect() (session chan string) {
	session = make(chan string, 0)
	go func(tSession chan string) {
		for {
			tRequest := <-tSession
			if tRequest == "CLOSE" {
				break
			}

			var request Request
			err := json.Unmarshal([]byte(tRequest), &request)
			if err != nil {
				_, _ = fmt.Println("Invalid request format:", tRequest)
			}
			response := server.Handle(request.Method, request.Params)
			b, err := json.Marshal(response)
			tSession <- string(b)
		}
		_, _ = fmt.Println("Session closed.")
	}(session)
	_, _ = fmt.Println("A new session has been created successfully.")
	return
}
