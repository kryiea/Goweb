package main

import (
	"Goweb/frame"
	"context"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *frame.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	duration, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		// 异常结束
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		// 正常结束
		fmt.Println("finish")
	case <-duration.Done():
		// 超时结束
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()

		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
