package v1

import (
	"github.com/gin-gonic/gin"
	"gosky/app/services/user"
	"gosky/infra/errcode"

	"gosky/infra/response"
)

// LoginController 登录控制器
type LoginController struct {
}

// login
func (cc *LoginController) Login(c *gin.Context) {
	guid := c.GetHeader("guid")
	userService := user.NewUserService()
	userAllInfo, err := userService.GetUserInfoByGuidOrCreate(c, guid)
	if err != nil {
		response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer)
		return
	}
	response.Success(c, errcode.ErrCodes.ErrNo, userAllInfo)
	return
}
