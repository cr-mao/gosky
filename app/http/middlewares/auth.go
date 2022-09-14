package middlewares

import (
	"github.com/gin-gonic/gin"
	"gosky/app/services/user"
	"gosky/infra/errcode"
	"gosky/infra/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		guid := c.GetHeader("guid")
		if len(guid) == 0 {
			response.ErrorAbort(c, errcode.ErrCodes.ErrParams)
		}
		userService := user.NewUserService()
		if !userService.IsExistByGuid(c, guid) {
			response.ErrorAbort(c, errcode.ErrCodes.ErrAuthenticationHeader)
			return
		}
		c.Next()
	}
}
