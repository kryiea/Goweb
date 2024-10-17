package main

import "Goweb/frame"

func registerRouter(core *frame.Core) {
	core.Get("foo", FooControllerHandler)
}
