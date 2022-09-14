package apitests

import (
	"gosky/bootstrap"
	"gosky/config"
)

func init() {
	// 加载 config 目录下的配置信息
	config.Initialize()
	bootstrap.Bootstrap("local")
}
