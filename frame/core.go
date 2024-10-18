package frame

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router map[string]*Tree // 路由树
}

// 初始化核心结构
func NewCore() *Core {
	// 初始化路由
	router := make(map[string]*Tree)
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router}
}

// 框架核心结构，实现 handler 接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装cont
	ctx := NewContext(request, response)

	// 尝试匹配路由
	router := c.FindRouteByRequest(ctx.request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner errorr")
		return
	}

}

// 路由注册函数，method大小写不敏感
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
