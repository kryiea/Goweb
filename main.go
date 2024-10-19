package main

import (
	"Goweb/frame"
	"Goweb/frame/middleware"
	"log"
	"net/http"
)

func main() {
	core := frame.NewCore()
	// core中使用 use 注册中间件
	core.Use(
		middleware.Test1(),
		middleware.Test2(),
	)

	//  在group中批量设置路由
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())

	// 注册路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8080",
	}

	log.Println("Listening on " + server.Addr)
	server.ListenAndServe()
}
