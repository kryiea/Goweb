package demo

import (
	"fmt"

	"github.com/kryiea/GoWeb/frame"
)

// 具体的接口实例
type DemoService struct {
	// 实现接口
	Service

	// 参数
	c frame.Container
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(frame.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
