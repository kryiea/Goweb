package app

import (
	"errors"

	"github.com/kryiea/GoWeb/frame"
)

type GoWebAppProvider struct {
	BaseFolder string
}

func (G GoWebAppProvider) Params(container frame.Container) []interface{} {
	return []interface{}{
		container,
		G.BaseFolder,
	}
}

func (G GoWebAppProvider) NewGoWebApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	conainer := params[0].(frame.Container)
	baseFolder := params[1].(string)
	return &GoWebApp{
		container:  conainer,
		baseFolder: baseFolder,
	}, nil
}
