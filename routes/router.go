package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	/*
		后台管理接口
	*/
	auth := r.Group("api/v1")
	// 注册中间件
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		// 修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		//	分类模块的路由接口
		auth.GET("admin/cagtegory", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		//	文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		//	 上传文件
		auth.POST("upload", v1.Upload)
		//	 更新个人设置

		// 评论模块

	}
	/*
		前端展示页面接口
	*/
	route := r.Group("api/v1")
	{
		// 用户信息模块
		route.POST("user/add", v1.AddUser)
		route.GET("user/:id", v1.GetUserInfo)
		route.GET("users", v1.GetUsers)
		//	 文章分类信息模块
		route.GET("category", v1.GetCate)
		route.GET("category/:id", v1.GetCateInfo)

		//	 文章模块
		route.GET("article", v1.GetArt)
		route.GET("article/list/:id", v1.GetCateArt)
		route.GET("article/info/:id", v1.GetArtInfo)
	}
	// 定义一个没有默认中间件的路由
	//r := gin.New()

	_ = r.Run(utils.HttpPort)
}
