package main

import (
	"Goweb/frame"
	"log"
	"net/http"
)

func main() {
	core := frame.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8080",
	}

	log.Println("Listening on " + server.Addr)
	server.ListenAndServe()
}
