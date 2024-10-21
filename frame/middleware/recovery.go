package middleware

import "Goweb/frame"

// recovery 机制，将协程中的函数异常进行捕获
func Recovery() frame.ControllerHandler {
	// 使用函数回调
	return func(c *frame.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()

		return nil
	}
}
