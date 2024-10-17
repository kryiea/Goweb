package frame

import (
	"log"
	"net/http"
)

// 框架核心结构
type Core struct {
	router map[string]ControllerHandler
}

// 初始化核心结构
func NewCore() *Core {
	return &Core{
		router: make(map[string]ControllerHandler),
	}
}

// 框架核心结构，实现handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("Core.serveHTTP")

	ctx := NewContext(request, response)

	// 写死
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")
	//
	router(ctx)
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}
