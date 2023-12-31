package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

// 文章类别
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}

// 查询分类是否存在(添加分类和编辑分类的时候要判断一下)
func CheckoutCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询单个分类的信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

// 查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cates []Category
	var total int64
	err := db.Find(&cates).Limit(pageSize).Offset((pageSize - 1) * pageNum).Error
	db.Model(&cates).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
