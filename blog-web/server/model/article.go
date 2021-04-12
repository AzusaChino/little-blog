package model

import "time"

// Article 文章实体类
type Article struct {
	baseEntity
	Topic        string
	Thumbnail    string
	PublishState int
	PublishTime  time.Time
}

// ArticleDetail 文章详情实体类
type ArticleDetail struct {
	Article
	Content string
}

// TableName gorm指定表名
func (_ *Article) TableName() string {
	return "tb_article"
}

// TableName gorm指定表名
func (_ *ArticleDetail) TableName() string {
	return "tb_article"
}
