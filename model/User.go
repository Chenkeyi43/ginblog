package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username,omitempty"`
	Password string `gorm:"type:varchar(200);not null" json:"password,omitempty"`
	Role     int    `gorm:"type:int;not null" json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)

	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 更新查询（更新用户名的时候要查下）
func CheckUpUser(id int, username string) int {
	var user User
	db.Select("id,username").Where("id = ?", id).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
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
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).Limit(1).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
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

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 修改密码
func ChangePassword(id int, data *User) int {
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// hook 对创建用户更新用户行为加密
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	//	  创建角色后的钩子，密码加密和权限控制
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// 使用bcrypt 加密密码
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

// 后台登入验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCSE
}

// 前台登入
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswoedErr error
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_ART_NOT_EXIST
	}
	PasswoedErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswoedErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCSE
}
