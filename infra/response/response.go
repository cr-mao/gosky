package response

import (
	"github.com/gin-gonic/gin"
	"gosky/infra/errcode"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

//失败响应函数,禁止调用下一层函数
func ErrorAbort(c *gin.Context, errCode errcode.ErrCode, msg ...string) {
	defaultMsg := errCode.Desc
	if len(msg) > 0 {
		defaultMsg = msg[0]
	}
	c.AbortWithStatusJSON(errCode.HTTPCode, Response{
		ErrorCode: errCode.Code,
		Msg:       defaultMsg,
		Data:      nil,
	})
}

//成功响应函数
func Success(c *gin.Context, errCode errcode.ErrCode, data interface{}) {
	c.JSON(errCode.HTTPCode, Response{
		ErrorCode: errCode.Code,
		Msg:       errCode.Desc,
		Data:      data,
	})
}

//失败响应函数
func Error(c *gin.Context, errCode errcode.ErrCode, msg ...string) {
	defaultMsg := errCode.Desc
	if len(msg) > 0 {
		defaultMsg = msg[0]
	}
	c.JSON(errCode.HTTPCode, Response{
		ErrorCode: errCode.Code,
		Msg:       defaultMsg,
		Data:      nil,
	})
}
