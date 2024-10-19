package frame

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler // 当前请求的hanler链条
	index          int                 // 当前请求调用到调用链的哪个节点，初始值应该为 -1，每次调用都会自增 1

	// 超时标志位
	hasTimeout bool
	// 写保护锁
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		hasTimeout:     false,
		index:          -1,
	}
}

// Next 方法每调用一次，就将这个控制器链路的调用控制器，往后移动一步。
func (c *Context) Next() error {
	c.index++
	if c.index < len(c.handlers) {
		if err := c.handlers[c.index](c); err != nil {
			return err
		}
	}
	return nil
}

// base functions
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.WriterMux()
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponseWriter() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// base context
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// query url
func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			return vals[length-1]
		}
	}
	return def
}
func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			intVal, err := strconv.Atoi(vals[length-1])
			if err == nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		// 类型转换
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

// post
func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			intval, err := strconv.Atoi(vals[length-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		length := len(vals)
		if length > 0 {
			return vals[length-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

// application/json post
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := io.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}

		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}

	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// response
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}

func (c *Context) SetHandlers(handlers []ControllerHandler) {
	// TODO
	c.handlers = handlers
}
