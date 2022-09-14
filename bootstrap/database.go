package bootstrap

import (
	"fmt"
	"gosky/infra/console"
	"sync"
	"time"

	"gosky/infra/conf"
	"gosky/infra/db"
	"gosky/infra/logger"
)

var dbOnce sync.Once

// SetupDB 初始化数据库
func SetupDB() {
	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
			conf.Get("database.default.username"),
			conf.Get("database.default.password"),
			conf.Get("database.default.host"),
			conf.Get("database.default.port"),
			conf.Get("database.default.database"),
			conf.Get("database.default.charset"),
			conf.Get("app.timezone"),
		)

		err := db.InitMysqlClientWithOptions(conf.Get("database.default.database"), dsn,
			logger.NewGormLogger(),
			db.WithMaxOpenConn(conf.GetInt("database.default.max_open_connections")),
			db.WithMaxIdleConn(conf.GetInt("database.default.max_idle_connections")),
			db.WithEnableSqlLog(conf.GetBool("database.default.enable_sql_log")),
			db.WithPrepareStmt(false),
			db.WithSlowLogMillisecond(conf.GetInt64("database.default.slow_log_millisecond")),
			db.WithConnMaxLifeSecond(time.Duration(conf.GetInt("database.default.max_life_seconds"))),
		)
		if err != nil {
			panic(fmt.Sprintf("数据库连接错误%s", err.Error()))
		}
		console.Success("mysql connect success:" + dsn)
	})
}
