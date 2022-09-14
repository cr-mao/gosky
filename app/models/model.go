// Package models 模型通用属性和方法
package models

// CommonTimestampsField 时间戳
//type CommonTimestampsField struct {
//	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
//	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
//}

type UserAllInfo struct {
	IsNew int8 `json:"is_new"`
	// 用户唯一标志 即token
	Guid string `json:"guid"`
	// 0正常1禁用
	ForbiddenStatus int8 `json:"forbidden_status"`
}
