package model

import "gorm.io/gorm"

// 文章详细
type Article struct {
	gorm.Model
	Titile       string   `json:"titile" grom:"type:varchar(100);not null"`
	Category     Category `json:"category" gorm:"foreignkey:Cid"`
	Cid          int      `json:"cid" gorm:"type:int"`
	Desc         string   `json:"desc" c"`
	Content      string   `json:"content" gorm:"type:longtext"`
	Img          string   `json:"img" grom:"type:varchar(100)""`
	CommentCount int      `json:"comment_count" gorm:"type:int;not null;default 0"`
	ReadCount    int      `json:"read_count" gorm:"type:int;not null;default 0"`
}
