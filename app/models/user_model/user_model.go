//Package user 模型
package user_model

import (
	"gorm.io/gorm"
)

type User struct {
	//fix
	UserId int32 `gorm:"column:user_id;primaryKey;autoIncrement;" json:"user_id,omitempty"`
	// user_id
	// 用户唯一标志 即token
	Guid string `gorm:"column:guid;index;" json:"guid"`
	// 0正常1禁用
	ForbiddenStatus int8 `gorm:"column:forbidden_status" json:"forbidden_status"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Create(tx *gorm.DB) (int32, error) {
	err := tx.Create(&user).Error
	return user.UserId, err
}

func (user *User) Save(tx *gorm.DB) (rowsAffected int64) {
	result := tx.Save(&user)
	return result.RowsAffected
}

func (user *User) Delete(tx *gorm.DB) (rowsAffected int64) {
	result := tx.Delete(&user)
	return result.RowsAffected
}
