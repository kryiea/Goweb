## GoWeb 框架项目
- Go 1.22.1


## 扣腚日记

### 10.21
1. context 支持链式调用：`c.SetStatus(500).Json("time out")`
1. 在 context 这个数据结构中，封装和实现“读取请求数据”和“封装返回数据”的方法
   1. 抽象出两个接口 `IReuqest`、`IResponse`，`context` 实现这两个接口。
2. 更优雅的关闭`server`进程：监听`os.signal` + `net/http` 提供的 `server.Shutdown`


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




