package bootstrap

import (
	"gosky/infra/conf"
)

func Bootstrap(env string) {
	// 配置初始化，依赖命令行 --env 参数
	conf.InitConfig(env)
	// 初始化 Logger
	SetupLogger()
	// 初始化数据库
	SetupDB()
	// 初始化 Redis
	//SetupRedis()
}
