package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SetUserStatusRequest struct {
	Status int64 `valid:"status" json:"status"`
}

func SetUserStatus(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"status": []string{"numeric_between:0,1"},
	}
	messages := govalidator.MapData{
		"status": []string{
			"numeric_between: value Illegal",
		},
	}
	return validate(data, rules, messages)
}
