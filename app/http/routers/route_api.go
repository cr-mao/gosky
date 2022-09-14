package routers

import (
	"gosky/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"

	"gosky/app/http/controllers/v1"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	routerV1 := r.Group("/v1")
	{
		loginController := new(v1.LoginController)
		routerV1.POST("/login", loginController.Login)

		appController := new(v1.AppController)
		routerV1.POST("/time_sync", appController.TimeSync)

		authGroup := routerV1.Use(middlewares.Auth())
		{
			userController := new(v1.UserController)
			authGroup.POST("/set_user_status", userController.SetUserStatus)
		}
	}

}
