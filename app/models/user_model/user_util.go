package user_model

import (
	"gorm.io/gorm"
)

func GetByGuid(tx *gorm.DB, value string) (*User, error) {
	userModel := &User{}
	err := tx.Where("guid = ?", value).First(userModel).Error
	return userModel, err
}

func IsExistByGuid(tx *gorm.DB, guid string) bool {
	var count int64
	tx.Model(User{}).Where("guid=?", guid).Count(&count)
	return count > 0
}
