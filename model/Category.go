package model

// 文章类别
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}
