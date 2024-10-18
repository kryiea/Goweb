package frame

// IGroup 路由分组接口
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

// Group 实现IGroup接口
type Group struct {
	core   *Core
	prefix string
}

// 初始化 Group
func NewGroup(core *Core, prefix string) IGroup {
	return &Group{core, prefix}
}

// Get 实现IGroup接口
func (g *Group) Get(path string, handler ControllerHandler) {
	g.core.Get(g.prefix+path, handler)
}

// Post 实现IGroup接口
func (g *Group) Post(path string, handler ControllerHandler) {
	g.core.Post(g.prefix+path, handler)
}

// Put 实现IGroup接口
func (g *Group) Put(path string, handler ControllerHandler) {
	g.core.Put(g.prefix+path, handler)
}

// Delete 实现IGroup接口
func (g *Group) Delete(path string, handler ControllerHandler) {
	g.core.Delete(g.prefix+path, handler)
}
