package study

import (
	"fmt"
	"net/http"
	"testing"
)

// 开启一个 http 服务
func TestHttpServer(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "<h1>Welcome! %s!</h1>", request.FormValue("name"))
	})
	_ = http.ListenAndServe(":8888", nil)
}
