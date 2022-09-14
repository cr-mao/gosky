package routers

import (
	"github.com/gin-gonic/gin"
	"gosky/infra/errcode"
	"gosky/infra/response"

	"gosky/app/http/middlewares"
	"gosky/infra/app"
	"gosky/infra/conf"
)

// 404处理
func setup404Handler(r *gin.Engine) {
	// 添加 Get 请求路路由
	r.NoRoute(func(c *gin.Context) {
		response.ErrorAbort(c, errcode.ErrCodes.ErrNotFound)
	})
}

//全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),      //自定义请求响应中间件
		middlewares.Recovery(),    //panic   错误 拦截处理
		middlewares.HeaderCheck(), //头校验，参数是否传了，签名等

		//middleware.RateLimit(),        //请求限流
		//middleware.MetheusPathCount(), //请求方法 统计基数 监控
	)
}

func NewRouter() *gin.Engine {
	if app.IsLocal() && conf.GetBool("app.debug", false) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	registerGlobalMiddleWare(router)
	setup404Handler(router)
	RegisterAPIRoutes(router)
	return router
}
