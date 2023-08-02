package model

import "gorm.io/gorm"

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username,omitempty"`
	Password string `gorm:"type:varchar(200);not null" json:"password,omitempty"`
	Role     int    `gorm:"type:int;not null" json:"role,omitempty"`
}
