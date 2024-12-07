// 服务接口说明文件

package demo

const Key = "Goweb:demo"

// demo 服务接口
type Service interface {
	GetFoo() Foo
}

// demo 服务接口定义的一个数据结构
type Foo struct {
	Name string
}
