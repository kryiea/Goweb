package frame

// IGroup 路由接口
type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)
	// 嵌套中间件
	Use(middlewares ...ControllerHandler)

	// 实现嵌套group
	Group(string) IGroup
}

// Group 实现 IGroup 接口
type Group struct {
	core        *Core               // core结构
	parent      *Group              // 指向上一个Group
	prefix      string              // Group 的通用前缀
	middlewares []ControllerHandler // 存放中间件
}

// 初始化 Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		prefix:      prefix,
		parent:      nil,
		middlewares: make([]ControllerHandler, 0),
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}

// 实现Post方法
func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

// 实现Put方法
func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

// 实现Delete方法
func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

// 取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 获取某个 group 的 middleware
// 这里就是获取除了 Get/Post/Put/Delete 之外设置的 middleware
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.middlewares == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现Group方法
func (g *Group) Group(rui string) IGroup {
	cgroup := NewGroup(g.core, rui)
	cgroup.parent = g
	return cgroup
}

// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}
