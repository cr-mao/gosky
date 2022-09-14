package config

import (
	"gosky/infra/conf"
)

func init() {
	//redis 相关配置
	conf.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			// 存缓存用
			"cache": map[string]interface{}{
				"host":     conf.Env("REDIS_HOST", "127.0.0.1"),
				"port":     conf.Env("REDIS_PORT", "6379"),
				"password": conf.Env("REDIS_PASSWORD", ""),
				"database": conf.Env("REDIS_CACHE_DB", 0),
			},
			// 数据不能丢的场景 ，如 会话信息
			"default": map[string]interface{}{
				"host":     conf.Env("REDIS_HOST", "127.0.0.1"),
				"port":     conf.Env("REDIS_PORT", "6379"),
				"password": conf.Env("REDIS_PASSWORD", ""),
				// 业务类存储使用 1 (会话 这种不能随意清楚的)
				"database": conf.Env("REDIS_DB", 1),
			},
		}
	})
}
