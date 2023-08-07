package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		//	分类模块的路由接口

		//	文章模块的路由接口
	}
	route := r.Group("api/v1")
	{
		route.POST("user/add", v1.AddUser)
	}
	// 定义一个没有默认中间件的路由
	//r := gin.New()

	_ = r.Run(utils.HttpPort)
}
