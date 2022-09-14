package bootstrap

import (
	"gosky/infra/conf"
	"gosky/infra/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		conf.GetString("log.filename"),
		conf.GetInt("log.max_size"),
		conf.GetInt("log.max_backup"),
		conf.GetInt("log.max_age"),
		conf.GetBool("log.compress"),
		conf.GetString("log.type"),
		conf.GetString("log.level"),
	)
}
