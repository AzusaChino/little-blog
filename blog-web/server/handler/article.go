package handler

import (
	"github.com/kataras/iris/v12"
	. "little-blog/entity"
	"net/http"
	"time"
)

// 文章实体类
type Article struct {
	BaseEntity
	Topic        string
	Thumbnail    string
	PublishState int
	PublishTime  time.Time
}

// 文章详情实体类
type ArticleDetail struct {
	Article
	Content string
}

// 列表控制器
type ArticleHandler struct {
}

// 详情控制器
type ArticleDetailHandler struct {
}

// gorm指定表名
func (_ *Article) TableName() string {
	return "tb_article"
}

// gorm指定表名
func (_ *ArticleDetail) TableName() string {
	return "tb_article"
}

func (_ *ArticleHandler) Method() string {
	return http.MethodGet
}

func (_ *ArticleHandler) HandlerFunc(ctx iris.Context) {
	_, _ = ctx.JSON(fetchArticleList())
}

func (_ *ArticleHandler) Path() string {
	return "api/v1/article"
}

func (_ *ArticleDetailHandler) Method() string {
	return http.MethodGet
}

func (_ *ArticleDetailHandler) HandlerFunc(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, _ = ctx.JSON(fetchArticleDetail(id))
}

func (_ *ArticleDetailHandler) Path() string {
	return "api/v1/article/{id}"
}

func fetchArticleList() []Article {
	db := GetDb()
	var articles []Article

	// 查找已发布且未删除的文章
	db.Where("publish_state = ? AND is_delete = ?", 1, 0).Find(&articles)

	return articles
}

func fetchArticleDetail(id string) ArticleDetail {
	db := GetDb()
	var articleDetail ArticleDetail

	db.Where("id = ?", id).Find(&articleDetail)

	return articleDetail
}
