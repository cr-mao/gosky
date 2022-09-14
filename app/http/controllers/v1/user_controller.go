package v1

import (
	"github.com/gin-gonic/gin"
	"gosky/app/requests"
	"gosky/infra/errcode"
	"gosky/infra/response"
)

type UserController struct {
}

func (cc *UserController) SetUserStatus(c *gin.Context) {
	var setUserStatusReq requests.SetUserStatusRequest
	if ok := requests.Validate(c, &setUserStatusReq, requests.SetUserStatus); !ok {
		return
	}
	//todo  修改数据库
	response.Success(c, errcode.ErrCodes.ErrNo, nil)
}
