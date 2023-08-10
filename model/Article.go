package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

/*
reload与Joins的区别
Preload方法是用来加载关联字段（belongTo、many2many、hasOne、hasMany）的数据的。
gorm中的Joins方法仅适用的查询，无法加载关联字段内容。并且，gorm原生的方法只支持一对一关系(has one, belongs to)
*/

// 文章详细
type Article struct {
	gorm.Model
	Titile       string   `json:"titile" grom:"type:varchar(100);not null"`
	Category     Category `json:"category" gorm:"foreignkey:Cid"`
	Cid          int      `json:"cid" gorm:"type:int"`
	Desc         string   `json:"desc" gorm:"type:varchar(200)"`
	Content      string   `json:"content" gorm:"type:longtext"`
	Img          string   `json:"img" grom:"type:varchar(100)""`
	CommentCount int      `json:"comment_count" gorm:"type:int;not null;default 0"`
	ReadCount    int      `json:"read_count" gorm:"type:int;not null;default 0"`
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

/*
	查询分类下的所有文章

parm:

	id: 文章类别ID
	pageSize: 每页显示数量
	pageNum: 哪一页

return:

	[]Article : 文章列表
	int:   状态码
	int64: 数量
*/
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	// 预加载分类表：后台处理数据库一对多的关系时，当然希望“多”的这部分部分数据乖乖的存放在数组中
	// 查找 cateartlist 时预加载 category
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid = ?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Count(&total)
	return cateArtList, errmsg.SUCCSE, total
}

// 查询单个文章内容
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	//	更新阅读次数
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err := db.Select("article.id, title, img, " +
		"created_at, updated_at, `desc`, comment_count, " +
		"read_count, category.name").Limit(
		pageSize).Offset((pageNum - 1) * pageSize).Order(
		"Created_At DESC").Joins("Category").Find(
		&articleList).Error

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// 根据文章标题搜索
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64

	err := db.Select("article.id,title, img, created_at, "+
		"updated_at, `desc`, comment_count, read_count, Category.name").Order(
		"Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	//	 计数
	db.Model(&articleList).Where("title Like ?", title+"%").Count(&total)

	return articleList, errmsg.SUCCSE, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Titile
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
