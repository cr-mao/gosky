// Package app 应用信息
package app

import (
	"gosky/infra/conf"
	"time"
)

func IsLocal() bool {
	return conf.Get("app.env") == "local"
}

func IsProduction() bool {
	return conf.Get("app.env") == "production"
}

func IsTesting() bool {
	return conf.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	location, _ := time.LoadLocation(conf.GetString("app.timezone"))
	return time.Now().In(location)
}
