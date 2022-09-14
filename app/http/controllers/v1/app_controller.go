package v1

import (
	"time"

	"github.com/gin-gonic/gin"

	"gosky/infra/errcode"
	"gosky/infra/helpers"
	"gosky/infra/response"
)

type AppController struct {
}

// 客户端时间同步
func (cc *AppController) TimeSync(c *gin.Context) {
	var result = map[string]string{
		"utc_time": time.Now().Format(helpers.CSTLayout),
	}
	response.Success(c, errcode.ErrCodes.ErrNo, result)
}
