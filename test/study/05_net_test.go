package study

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"
)

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"msg":  "Welcome",
		"name": request.FormValue("name"),
		"url":  html.EscapeString(request.RequestURI)})
	_, _ = fmt.Fprintf(writer, string(jsonBytes))
}

// 开启一个 http 服务
func TestHttpServer(t *testing.T) {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
