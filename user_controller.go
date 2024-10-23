package main

import "github.com/kryiea/GoWeb/frame"

func UserLoginController(c *frame.Context) error {
	// 打印控制器名字
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
