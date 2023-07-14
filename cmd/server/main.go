// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"notification/internal/server"
)

func main() {
	// 创建服务端实例
	s := server.NewServer()

	// 注册处理函数
	http.HandleFunc("/", s.HandleRequest)

	// 监听端口并处理请求
	log.Fatal(http.ListenAndServe(":8080", nil))
}
