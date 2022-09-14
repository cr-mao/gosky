package bootstrap

import (
	"fmt"
	"gosky/infra/console"
	"strconv"
	"sync"

	"gosky/infra/conf"
	"gosky/infra/redis"
)

var redisOnce sync.Once

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redisOnce.Do(func() {
		//专门做cache用
		cacheAddress := fmt.Sprintf("%v:%v", conf.GetString("redis.cache.host"), conf.GetString("redis.cache.port"))
		redis.ConnectRedis(
			"cache",
			cacheAddress,
			conf.GetString("redis.cache.username"),
			conf.GetString("redis.cache.password"),
			conf.GetInt("redis.cache.database"),
		)
		console.Success("redis cache connect success:" + cacheAddress + " db:" + strconv.Itoa(conf.GetInt("redis.cache.database")))
		//专门做 非cache用 ，不能清空的场景 如会话
		defaultAddress := fmt.Sprintf("%v:%v", conf.GetString("redis.default.host"), conf.GetString("redis.default.port"))
		redis.ConnectRedis(
			"default",
			defaultAddress,
			conf.GetString("redis.default.username"),
			conf.GetString("redis.default.password"),
			conf.GetInt("redis.default.database"),
		)
		console.Success("redis default connect success:" + defaultAddress + " db:" + strconv.Itoa(conf.GetInt("redis.default.database")))

	})
}
