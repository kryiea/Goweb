package gin

import (
	"context"
	"fmt"

	"github.com/kryiea/GoWeb/frame"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

func (engine *Engine) Bind(provide frame.ServiceProvider) error {
	return engine.container.Bind(provide)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 实现 container 的几个封装
// 实现 make 的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 实现 mustMake 的封装
func (ctx *Context) MustMake(key string) interface{} {
	if ctx.container == nil {
		fmt.Println("ctx.container is nil")
		return nil
	}
	return ctx.container.MustMake(key)
}

// 实现 makenew 的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
