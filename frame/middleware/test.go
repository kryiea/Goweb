package middleware

import (
	"fmt"

	"github.com/kryiea/GoWeb/frame/gin"
)

// Test1 是一个中间件函数，它在处理请求之前和之后执行一些操作
func Test1() gin.HandlerFunc {
    // 使用函数回调
    return func(c *gin.Context) {
        fmt.Println("middleware pre test1")
        c.Next()
        fmt.Println("middleware post test1")
    }
}

// Test2 是一个中间件函数，它在处理请求之前和之后执行一些操作
func Test2() gin.HandlerFunc {
    // 使用函数回调
    return func(c *gin.Context) {
        fmt.Println("middleware pre test2")
        c.Next()
        fmt.Println("middleware post test2")
    }
}

// Test3 是一个中间件函数，它在处理请求之前和之后执行一些操作
func Test3() gin.HandlerFunc {
    // 使用函数回调
    return func(c *gin.Context) {
        fmt.Println("middleware pre test3")
        c.Next()
        fmt.Println("middleware post test3")
    }
}
