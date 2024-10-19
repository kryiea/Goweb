package frame

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 1. 执行业务逻辑前的预操作： 初始化超时context
		duration, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(duration)

		// 2. 执行业务逻辑
		go func() {
			// 捕获异常
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行
			c.Next()
			// 完成
			finish <- struct{}{}
		}()

		// 3. 等待超时或者业务逻辑执行完成
		select {
		case <-finish:
			fmt.Println("finish")
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-duration.Done():
			c.SetHasTimeout()
			c.responseWriter.Write([]byte("timeout"))
		}
		return nil
	}
}
