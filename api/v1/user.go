package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询用户是否存在
func GetUser(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	//	todo 添加用户
	var data model.User
	// 绑定 request body 到结构体中
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	//model.GetUserS()
	// 从Get 请求中获取相关参数
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	users, total := model.GetUserS(username, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCSE,
		"data":    users,
		"total":   total,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DelUser(c *gin.Context) {

}
