package frame

// NewInstance 实例化
type NewInstance func(...interface{}) (interface{}, error)

// serviceProvider 服务提供者
type ServiceProvider interface {
	// 注册服务到容器中
	Register(Container) NewInstance

	// 启动服务,可以在服务启动时执行一些初始化操作：基础配置等等。
	// 如果返回error，表示整个服务实例化失败。
	Boot(Container) error

	// 判断服务是否延迟加载, false 为立即加载
	IsDefer() bool

	// 获取服务提供者的参数列表，定义传递给NewInstance的参数列表
	Params(Container) []interface{}

	// 获取服务提供者的凭证
	Name() string
}
