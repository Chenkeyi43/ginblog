package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	code := model.CheckoutCategory(category.Name)
	if code == errmsg.SUCCSE {
		model.CreateCate(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100

	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCate(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCSE,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})
}

// 查询分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCateInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckoutCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.EditCate(id, &data)
	}
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
