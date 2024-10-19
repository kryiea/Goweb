package middleware

import (
	"Goweb/frame"
	"fmt"
)

func Test1() frame.ControllerHandler {
	return func(c *frame.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() frame.ControllerHandler {
	return func(c *frame.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}
func Test3() frame.ControllerHandler {
	return func(c *frame.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
