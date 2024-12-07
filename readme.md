## GoWeb 框架项目
**GoWeb 框架目标：生产中可用且具有丰富功能模块的框架**

 开发环境要求：
- Go 1.22.1
- gin 1.7.3


## 扣腚记

### 12.6

#### 一、框架设计核心思想: 服务提供者和服务容器理论
- `服务容器`: 框架主体作为一个服务容器。
- `服务提供者`: 其他各个服务模块都作为服务提供者，在服务容器中注册自己的服务凭证和服务接口，通过服务凭证来获取具体的服务实例。

这样，功能的具体实现交给了各个服务模块，我们只需要规范服务提供者也就是服务容器中的接口协议。

每个模块服务都是一个`服务提供者（service provider）`，而我们主体框架需要承担起来的角色叫做`服务容器（service container）`，服务容器中绑定了多个接口协议，每个接口协议都由一个服务提供者提供服务。
![框架架构](/pic/框架架构.png)

#### 二、 服务提供者的设计

服务提供者提供的是“创建服务实例的方法”，服务容器提供的是“实例化服务的方法”。


> 服务提供者接口定义：

定义了五个基本能力：
   - `Name() string`：获取服务凭证的能力。
   - `Register(Container) NewInstance`：注册服务实例化方法的能力。
   - `Params(Container) []interface{}`：获取服务实例化方法参数的能力。
   - `IsDefer() bool`：控制实例化时机的能力。
   - `Boot(Container) error`：实例化预处理的能力。
- `NewInstance`函数定义了创建服务实例的方法，规范了输入输出参数。

​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​ 



#### 三、 服务容器的设计
> container 接口设计
1. 注册绑定服务提供者：通过 `Bind` 方法将服务提供者注册到容器中。
2. 获取服务实例：通过 `Make` 方法根据关键字凭证获取服务实例。
3. 确认服务绑定状态：通过 `IsBind` 方法检查某个关键字凭证是否已经绑定服务提供者。
4. 强制获取服务实例：通过 `MustMake` 方法获取服务实例，不返回错误信息，但未绑定服务提供者时会触发 `panic`。
5. 按需创建服务实例：通过 `MakeNew` 方法根据不同参数创建新的服务实例。


> GowebContainer 实现 container 接口

在实现 `GowebContainer` 数据结构时，我们使用了 `map[string]interface{}` 来存储服务实例和 `map[string]ServiceProvider` 来存储服务提供者。为了保证并发性，我们使用了读写锁机制。

具体实现中，`Bind` 方法负责将服务提供者注册到容器中，并处理实例化逻辑。`Make` 方法则根据关键字凭证获取服务实例，并考虑是否需要强制初始化实例。

#### 四、服务容器与框架结合

为了将服务容器融合进框架中，我们在 `Engine` 数据结构中增加了 `container` 字段，并在初始化 `Engine` 时同时初始化 `container`。在创建 `Context` 时，将 `container` 传递进入 `Context`。

此外，我们还为 `Engine` 封装了 `Bind` 和 `IsBind` 方法，为 `Context` 封装了 `Make`、`MustMake` 和 `MakeNew` 方法。

> 创建服务提供方

在业务目录中，我们创建了一个示例服务提供方 `DemoServiceProvider`，并定义了相应的服务接口 `Service`。在 `DemoServiceProvider` 中实现了 `Name`、`Register`、`Params`、`IsDefer` 和 `Boot` 方法。

#### 五、服务使用示例

在 `main.go` 中，我们通过 `Engine` 的 `Bind` 方法绑定了 `DemoServiceProvider`。在 `SubjectListController` 中，通过 `Context` 的 `MustMake` 方法获取了 `DemoService` 实例，并调用了其 `GetFoo` 方法。

#### 六、验证结果

最后，我们在浏览器中访问了 `/subject/list/all` 路由，成功获取到了 `Foo` 数据结构的 JSON 输出，验证了服务容器和服务提供方的正确性。




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




