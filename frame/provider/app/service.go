package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/kryiea/GoWeb/frame"
	"github.com/kryiea/GoWeb/frame/util"
)

// NewGoWebApp 代表Goweb框架的app实现
type GoWebApp struct {
	container  frame.Container //服务容器
	baseFolder string          // 基础目录
}

// version 表示版本号
func (g GoWebApp) Version() string {
	return "0.0.1"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (g GoWebApp) BaseFolder() string {
	// 检查服务提供者有无设置基础目录
	if g.baseFolder != "" {
		return g.baseFolder
	}
	// 如果没有设置，判断命令行参数中是否有设置
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数, 默认为当前路径")
	flag.Parse()

	if baseFolder != "" {
		return baseFolder
	}
	// 如果命令行参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (g GoWebApp) ConfigFolder() string {
	return filepath.Join(g.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (g GoWebApp) LogFolder() string {
	return filepath.Join(g.StorageFolder(), "log")
}

func (g GoWebApp) gttpFolder() string {
	return filepath.Join(g.BaseFolder(), "gttp")
}

func (g GoWebApp) ConsoleFolder() string {
	return filepath.Join(g.BaseFolder(), "console")
}

func (g GoWebApp) StorageFolder() string {
	return filepath.Join(g.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (g GoWebApp) ProviderFolder() string {
	return filepath.Join(g.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (g GoWebApp) MiddlewareFolder() string {
	return filepath.Join(g.gttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (g GoWebApp) CommandFolder() string {
	return filepath.Join(g.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (g GoWebApp) RuntimeFolder() string {
	return filepath.Join(g.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (g GoWebApp) TestFolder() string {
	return filepath.Join(g.BaseFolder(), "test")
}

// NewgadeApp 初始化gadeApp
func NewgadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(frame.Container)
	baseFolder := params[1].(string)
	return &GoWebApp{baseFolder: baseFolder, container: container}, nil
}
