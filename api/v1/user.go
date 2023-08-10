package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
func GetUserInfo(c *gin.Context) {
	// c.Param 解析路由参数
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	//model.GetUserS()
	// 查询url 请求字符串参数
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
	var data model.User
	// 获取url 参数
	id, _ := strconv.Atoi(c.Param("id"))
	// 绑定body 到结构体
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改用户密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.ChangePassword(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
