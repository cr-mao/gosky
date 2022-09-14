package middlewares

import (
	"github.com/gin-gonic/gin"
	"gosky/infra/errcode"
	"gosky/infra/response"
)

func HeaderCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		guid := c.GetHeader("guid")
		if len(guid) == 0 {
			response.ErrorAbort(c, errcode.ErrCodes.ErrParams, "guid required")
			return
		}
		c.Next()
	}

}
