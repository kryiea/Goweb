package main

import (
	"Goweb/frame"
	"Goweb/frame/middleware"
)

// 注册路由规则
func registerRouter(core *frame.Core) {
	// 需求1+2: HTTP方法+静态路由匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3: 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4: 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
