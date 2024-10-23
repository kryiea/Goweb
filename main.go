package main

import (
	"context"
	"github.com/kryiea/GoWeb/frame"
	"github.com/kryiea/GoWeb/frame/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := frame.NewCore()

	core.Use(middleware.Recovery())

	// 注册路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8080",
	}

	// 启动服务的 goroutine
	go func() {
		log.Println("Listening on " + server.Addr)
		server.ListenAndServe()
	}()

	// 当前的 goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监听信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 等待关闭信号
	<-quit

	// 接收到信号，关闭服务
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
