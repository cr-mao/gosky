package user

import (
	"context"

	"gosky/app/models"
	"gosky/app/models/user_model"
	"gosky/infra/db"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

//用户初始化
func (s *UserService) UserInit(ctx context.Context, guid string) error {
	var (
		err error
	)
	userModel := &user_model.User{
		Guid: guid,
	}
	tx := db.Tx(ctx, "demo")
	_, err = userModel.Create(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//根据guid 获得用户信息
func (s *UserService) GetUserInfoByGuid(ctx context.Context, guid string) (*models.UserAllInfo, error) {
	var (
		err     error
		userRes *user_model.User
	)
	session := db.Session(ctx, "demo")
	userRes, err = user_model.GetByGuid(session, guid)
	if err != nil {
		return nil, err
	}
	return &models.UserAllInfo{
		Guid:            userRes.Guid,
		ForbiddenStatus: userRes.ForbiddenStatus,
	}, nil
}

//查询guid ，没查到创建 并返回 用户信息
func (s *UserService) GetUserInfoByGuidOrCreate(ctx context.Context, guid string) (*models.UserAllInfo, error) {
	var (
		err error
	)
	session := db.Session(ctx, "demo")
	isExistUser := user_model.IsExistByGuid(session, guid)
	if !isExistUser {
		err = s.UserInit(ctx, guid)
		if err != nil {
			return nil, err
		}
	}
	userAllInfo, err := s.GetUserInfoByGuid(ctx, guid)
	if !isExistUser {
		userAllInfo.IsNew = 1
	}
	return userAllInfo, err
}

func (s *UserService) IsExistByGuid(ctx context.Context, guid string) bool {
	session := db.Session(ctx, "demo")
	return user_model.IsExistByGuid(session, guid)
}
