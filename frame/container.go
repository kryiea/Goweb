package frame

import (
	"errors"
	"fmt"
	"sync"
)

// Container 接口 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回 error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// GowebContainer 是一个服务容器的实现，它实现了 Container 接口
type GowebContainer struct {
	// 强制要求 GowebContainer 实现 Container 接口
	Container

	// providers 存储注册的服务提供者，key 为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key 为字符串凭证
	instances map[string]interface{}
	// 读写锁 lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// Bind 方法用于将服务提供者绑定到容器中
func (Goweb *GowebContainer) Bind(provider ServiceProvider) error {
	// 加锁，保证线程安全
	Goweb.lock.Lock()
	defer Goweb.lock.Unlock()
	// 获取服务提供者的名称，作为键
	key := provider.Name()

	// 注册并记录服务提供者
	Goweb.providers[key] = provider

	// 检查服务提供者是否延迟加载
	if provider.IsDefer() {
		// 如果是延迟加载，调用 Boot 方法进行初始化
		if err := provider.Boot(Goweb); err != nil {
			return err
		}
	}

	// 获取服务提供者的参数列表
	params := provider.Params(Goweb)
	// 获取服务提供者的注册方法
	method := provider.Register(Goweb)
	// 调用注册方法实例化服务
	instance, err := method(params...)

	// 如果实例化失败，返回错误
	if err != nil {
		return errors.New(err.Error())
	}
	// 将实例化的服务保存到容器中
	Goweb.instances[key] = instance

	return nil
}

// Make 方式调用内部的 make 实现
func (Goweb *GowebContainer) Make(key string) (interface{}, error) {
	return Goweb.make(key, nil, false)
}

func (Goweb *GowebContainer) MustMake(key string) interface{} {
	serv, err := Goweb.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

// MakeNew 方式使用内部的 make 初始化
func (Goweb *GowebContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return Goweb.make(key, params, true)
}

// findServiceProvider 方法用于在 GowebContainer 中查找指定 key 的服务提供者
func (Goweb *GowebContainer) findServiceProvider(key string) ServiceProvider {
	// 加读锁，保证在查找过程中，其他 goroutine 不能修改 providers 映射
	Goweb.lock.RLock()
	// 查找结束后，释放读锁
	defer Goweb.lock.RUnlock()
	// 尝试从 providers 映射中获取服务提供者
	if sp, ok := Goweb.providers[key]; ok {
		// 如果找到，返回服务提供者
		return sp
	}
	// 如果未找到，返回 nil
	return nil
}
func (Goweb *GowebContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(Goweb); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(Goweb)
	}
	method := sp.Register(Goweb)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// make 函数用于创建或获取指定服务的实例
// forceNew 表示是否强制创建新的实例
func (Goweb *GowebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	// 加读锁
	Goweb.lock.RLock()
	defer Goweb.lock.RUnlock()

	// 查找服务提供者
	sp := Goweb.findServiceProvider(key)
	// 如果未找到服务提供者，则返回错误
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	// 如果 forceNew 为 true，则创建新的实例
	if forceNew {
		return Goweb.newInstance(sp, params)
	}

	// 检查实例是否已经存在
	if ins, ok := Goweb.instances[key]; ok {
		return ins, nil
	}

	// 创建新的实例
	inst, err := Goweb.newInstance(sp, nil)
	// 如果创建实例时发生错误，则返回错误
	if err != nil {
		return nil, err
	}

	// 将新创建的实例保存到容器中
	Goweb.instances[key] = inst
	return inst, nil
}

func NewGowebContainer() *GowebContainer {
	return &GowebContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (hade *GowebContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range hade.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}
