//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"gosky/app/services/user"
)

var SuperSet = wire.NewSet(
	//u
	user.NewUserService,
)

// NOTE 如果在 cmd/job 中需要，在这里暴露 XXXService
// For cmd/job/**

func UserService() *user.UserService {
	wire.Build(SuperSet)
	return &user.UserService{}
}
