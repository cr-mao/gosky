package config

import (
	"gosky/infra/conf"
)

func init() {

	//mysql数据库相关配置, 暂时只有一个数据库实例，可能要增加主从对应的配置
	conf.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"default": map[string]interface{}{
				// 数据库连接信息
				"host":     conf.Env("DB_HOST", "127.0.0.1"),
				"port":     conf.Env("DB_PORT", "3306"),
				"database": conf.Env("DB_DATABASE", "test"),
				"username": conf.Env("DB_USERNAME", "root"),
				"password": conf.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
				// 连接池配置
				"max_idle_connections": conf.Env("DB_MAX_IDLE_CONNECTIONS", 10),
				"max_open_connections": conf.Env("DB_MAX_OPEN_CONNECTIONS", 100),
				"max_life_seconds":     conf.Env("DB_MAX_LIFE_SECONDS", 300),
				"enable_sql_log":       conf.Env("DB_SQL_LOG", false),
				"slow_log_millisecond": 500, //毫秒
			},
		}
	})
}
