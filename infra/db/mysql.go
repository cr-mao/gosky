package db

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gosky/infra/app"
)

type ConnectionMap struct {
	mapping map[string]*gorm.DB
}

type option struct {
	MaxOpenConn        int
	MaxIdleConn        int
	ConnMaxLifeSecond  time.Duration
	PrepareStmt        bool
	SlowLogMillisecond int64
	EnableSqlLog       bool
}

type Option func(*option)

func (o *option) reset() {
	o.MaxOpenConn = 0
	o.MaxIdleConn = 0
	o.ConnMaxLifeSecond = 0
	o.PrepareStmt = false
	o.SlowLogMillisecond = DefaultSlowLogMillisecond
}

const (
	DefaultMaxOpenConn        = 100
	DefaultMaxIdleConn        = 10
	DefaultLogName            = "gorm"
	DefaultSlowLogMillisecond = 500
)

var (
	connectionMap ConnectionMap
)

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(opt *option) {
		opt.MaxOpenConn = maxOpenConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(opt *option) {
		opt.MaxIdleConn = maxIdleConn
	}
}

func WithConnMaxLifeSecond(connMaxLifeTime time.Duration) Option {
	return func(opt *option) {
		opt.ConnMaxLifeSecond = connMaxLifeTime
	}
}

func WithSlowLogMillisecond(slowLogMillisecond int64) Option {
	return func(opt *option) {
		opt.SlowLogMillisecond = slowLogMillisecond
	}
}

func WithPrepareStmt(prepareStmt bool) Option {
	return func(opt *option) {
		opt.PrepareStmt = prepareStmt
	}
}
func WithEnableSqlLog(enableSqlLog bool) Option {
	return func(opt *option) {
		opt.EnableSqlLog = enableSqlLog
	}
}

func InitMysqlClientWithOptions(clientName, dsn string, logger gormlogger.Interface, options ...Option) error {
	if len(clientName) == 0 {
		return errors.New("client name is empty")
	}
	opt := &option{}
	for _, f := range options {
		if f != nil {
			f(opt)
		}
	}
	db, err := dbConnect(dsn, logger, opt)
	if err != nil {
		panic(clientName + "数据库连接失败" + err.Error())
	}
	connectionMap.mapping = make(map[string]*gorm.DB)
	connectionMap.mapping[clientName] = db
	//本地打印sql日志
	localWriteLog(db, opt.EnableSqlLog)
	return nil
}

//获得gorm.DB 事务实例
func Tx(parent context.Context, clientName string) *gorm.DB {
	if client, ok := connectionMap.mapping[clientName]; ok {
		ctx, ok := parent.(*gin.Context)
		if !ok {
			//  非gin.Context处理
			return client.Begin().WithContext(parent)
		}
		txdb, ok := ctx.Get("tx_id")
		//同一次请求 ， 重复用用Tx方法 第二次直接用老的事务，不能再Begin了,不然不是一个事务。
		if ok {
			return txdb.(*gorm.DB)
		}
		//第一次,withContext 会调 db.Session ，Session create new db session
		res := client.Begin().WithContext(parent)
		ctx.Set("tx_id", res)
		return res
	}
	return nil
}

//获得gorm.DB session实例
func Session(parent context.Context, clientName string) *gorm.DB {
	if client, ok := connectionMap.mapping[clientName]; ok {
		ctx, ok := parent.(*gin.Context)
		if !ok {
			return client.WithContext(parent)
		}
		sessionDb, ok := ctx.Get("session_id")
		//同一次请求 ， 重复用用Tx方法 第二次直接用老的事务，不能再Begin了,不然不是一个事务。
		if ok {
			return sessionDb.(*gorm.DB)
		}
		//第一次,withContext 会调 db.Session ，Session create new db session
		res := client.WithContext(parent)
		ctx.Set("session_id", res)
		return res
	}
	return nil
}

//关闭数据库
func CloseDb(clientName string) error {
	sqlDB, err := connectionMap.mapping[clientName].DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

//关闭所有数据库
func CloseAllDb() error {
	for _, v := range connectionMap.mapping {
		sqlDB, err := v.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//连接数据库
func dbConnect(dsn string, logger gormlogger.Interface, option *option) (*gorm.DB, error) {
	if option.SlowLogMillisecond == 0 {
		option.SlowLogMillisecond = DefaultSlowLogMillisecond
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除），  每一条sql都是一个事务
		//如果没有这方面的要求，可以设置SkipDefaultTransaction为true来禁用它。
		//SkipDefaultTransaction: true,
		Logger: logger,
		//执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续执行的效率
		PrepareStmt: option.PrepareStmt,
		NamingStrategy: schema.NamingStrategy{
			//使用单数表名,默认为复数表名，即当model的结构体为User时，默认操作的表名为users
			//设置	SingularTable: true 后当model的结构体为User时，操作的表名为user
			SingularTable: true,
			//TablePrefix: "pre_", //表前缀
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed]  dsn链接: %s", dsn))
	}

	//db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	if option.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(option.MaxOpenConn)
	} else {
		sqlDB.SetMaxOpenConns(DefaultMaxOpenConn)
	}
	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	if option.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(option.MaxIdleConn)
	} else {
		sqlDB.SetMaxIdleConns(DefaultMaxIdleConn)
	}
	// 设置最大连接超时时间
	if option.ConnMaxLifeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * option.ConnMaxLifeSecond)
	}

	return db, nil
}

// 打印sql,本地环境 并且开启sql打印功能的
func localWriteLog(db *gorm.DB, enableSqlLog bool) {
	//本地环境 打印sql 到控制台
	if app.IsLocal() && enableSqlLog {
		err := db.Callback().Create().After("gorm:after_create").Register(DefaultLogName, afterLog)
		if err != nil {
			db.Logger.Info(context.Background(), "reqister after_create err:", err.Error())
		}
		err = db.Callback().Query().After("gorm:after_query").Register(DefaultLogName, afterLog)
		if err != nil {
			db.Logger.Info(context.Background(), "reqister after_query err:", err.Error())
		}
		err = db.Callback().Update().After("gorm:after_update").Register(DefaultLogName, afterLog)
		if err != nil {
			db.Logger.Info(context.Background(), "reqister after_update err:", err.Error())
		}
		err = db.Callback().Delete().After("gorm:after_delete").Register(DefaultLogName, afterLog)
		if err != nil {
			db.Logger.Info(context.Background(), "reqister after_delete err:", err.Error())
		}
	}
}

func afterLog(db *gorm.DB) {
	err := db.Error
	//ctx := db.Statement.Context
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	if err != nil {
		db.Logger.Info(context.Background(), sql, err.Error())
	} else {
		fmt.Println("[ SQL语句 ]", sql)
	}
}
