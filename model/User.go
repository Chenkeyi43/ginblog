package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username,omitempty"`
	Password string `gorm:"type:varchar(200);not null" json:"password,omitempty"`
	Role     int    `gorm:"type:int;not null" json:"role,omitempty"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	result := db.Select("id").Where("username = ?", name).First(&user)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUserS(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role,create_at").Where(
			"username Like ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

		db.Model(&users).Where("username Like ?", username+"%").Count(&total)

		return users, total
	}
	// 传空返回
	return users, 0

}
