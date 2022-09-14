// Package config 站点配置信息
package config

import (
	"gosky/infra/conf"
)

func init() {
	//应用相关配置
	conf.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": conf.Env("APP_NAME", "SKYGO"),
			// 当前环境，用以区分多环境， local, testing, production
			"env": conf.Env("APP_ENV", "local"),
			// 是否进入调试模式
			"debug": conf.Env("APP_DEBUG", false),
			// 应用服务host，port,url 等
			"http_host":        conf.Env("HTTP_HOST", "127.0.0.1"),
			"http_port":        conf.Env("HTTP_PORT", "8001"),
			"http_connect":     conf.Env("HTTP_CONNECT", "127.0.0.1:8001"),
			"http_connect_url": conf.Env("HTTP_CONNECT_URL", "http://127.0.0.1:8001"),
			// 加密会话、JWT 加密
			"key": conf.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 设置时区，  JWT 里会使用，日志记录里也会使用到
			"timezone": conf.Env("TIMEZONE", "UTC"),
		}
	})
}
