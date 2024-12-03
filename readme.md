## GoWeb 框架项目
- Go 1.22.1
- gin 1.7.3


## 扣腚记

### 12.3
1. 引入gin1.7.3，基于gin做深度定制。
   2. 将 Gin 目录下的 go.mod 的内容放在项目目录下
   3. 将 Gin 中原有 Gin 库的引用地址，统一替换为当前项目的地址: github.com/kryiea/GoWeb/frame
2. 在gin目录新增Goweb_context.go文件，实现自定义的context。
3. 在gin目录新增Goweb_request.go文件，实现自定义的request。
4. 在gin目录新增Goweb_response.go文件，实现自定义的response。
5. 修改后通过所有gin的单元测试。
6. go build 项目成功。

``` 
➜  Goweb git:(master) ./GoWeb 
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /user/login               --> main.UserLoginController (4 handlers)
[GIN-debug] DELETE /subject/:id              --> main.SubjectDelController (4 handlers)
[GIN-debug] PUT    /subject/:id              --> main.SubjectUpdateController (4 handlers)
[GIN-debug] GET    /subject/:id              --> main.SubjectGetController (5 handlers)
[GIN-debug] GET    /subject/list/all         --> main.SubjectListController (4 handlers)
[GIN-debug] GET    /subject/info/name        --> main.SubjectNameController (4 handlers)
2024/12/03 16:13:35 Listening on :8080
2024/12/03 16:13:42 api uri start: /
2024/12/03 16:13:42 api uri end: /, cost: 1.6985e-05
```



#### 如何迁移
> gin框架具备的功能

Gin 的框架中，Context、路由、中间件，都已经有了 Gin 自己的实现

>GoWeb框架目前已经实现的功能

- Context：请求控制器，控制每个请求的超时等逻辑； 
- 路由：让请求更快寻找目标函数，并且支持通配符、分组等方式制定路由规则； 
- 中间件：能将通用逻辑转化为中间件，并串联中间件实现业务逻辑； 
- 封装：提供易用的逻辑，把 request 和 response 封装在 Context 结构中； 
- 重启：实现优雅关闭机制，让服务可以重启。



#### 迁移的关键点

- Context 方面，Gin 的实现基本和我们之前的实现是一致的。
- 之前实现的 Core 数据结构对应 Gin 中的 Engine，Group 数据结构对应 Gin 的 Group 结构，Context 数据结构对应 Gin 的 Context 数据结构。

> 关键点1

对于 Request 和 Response 的封装， Gin 的实现比较简陋。

Gin 对 Request 并没有以接口的方式，将 Request 支持哪些接口展示出来；并且在参数查询的时候，返回类型并不多。

> 关键点2

Response 中，我们的设计是带有链式调用方法的，而 Gin 中没有。



#### gin的开源许可协议
> gin框架使用MIT开源许可协议
- 允许被许可人使用、复制、修改、合并、出版发行、散布、再许可、售卖软件及软件副本。
- 唯一条件是在软件和软件副本中必须包含著作权声明和本许可声明。

> 复制形式使用 Gin 框架的话，需要在 Gin 框架每个文件的头部，增加上著作权和许可声明：

```
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
```


### 10.21
1. context 支持链式调用：`c.SetStatus(500).Json("time out")`
2. 在 context 这个数据结构中，封装和实现“读取请求数据”和“封装返回数据”的方法
   1. 抽象出两个接口 `IReuqest`、`IResponse`，`context` 实现这两个接口。
3. 更优雅的关闭`server`进程：监听`os.signal` + `net/http` 提供的 `server.Shutdown`


### 10.19
1. 链式中间件机制

### 10.18
1. route：字典树维护
2. 静态路由
3. 路由分组
4. 动态路由


### 10.17 
1. 初始化仓库
2. 封装自定义 context， 封装 请求和响应




