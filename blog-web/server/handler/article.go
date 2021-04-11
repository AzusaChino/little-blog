package handler

import (
	"github.com/kataras/iris/v12"
	. "little-blog/common"
	"net/http"
	"time"
)

// Article 文章实体类
type Article struct {
	BaseEntity
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

// ArticleHandler 列表控制器
type ArticleHandler struct {
}

// ArticleDetailHandler 详情控制器
type ArticleDetailHandler struct {
}

// TableName gorm指定表名
func (_ *Article) TableName() string {
	return "tb_article"
}

// TableName gorm指定表名
func (_ *ArticleDetail) TableName() string {
	return "tb_article"
}

func (_ *ArticleHandler) Method() string {
	return http.MethodGet
}

// HandlerFunc pagination search
func (_ *ArticleHandler) HandlerFunc(ctx iris.Context) {
	// not found => default 0 || return 0, v.notFound(reflect.Int)
	pageNum, _ := ctx.Params().GetInt("pageNum")
	pageSize, _ := ctx.Params().GetInt("pageSize")

	_, _ = ctx.JSON(*fetchArticleList(pageNum, pageSize))
}

func (_ *ArticleHandler) Path() string {
	return "api/v1/article"
}

func (_ *ArticleDetailHandler) Method() string {
	return http.MethodGet
}

func (_ *ArticleDetailHandler) HandlerFunc(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, _ = ctx.JSON(*fetchArticleDetail(id))
}

func (_ *ArticleDetailHandler) Path() string {
	return "api/v1/article/{id}"
}

func fetchArticleList(pageNum, pageSize int) *RestResponse {
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()

	var articles []Article

	// 查找已发布且未删除的文章
	db.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("publish_state = ? AND is_delete = ?", 1, 0).Find(&articles)

	return Ok(articles)
}

func fetchArticleDetail(id string) *RestResponse {
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()
	var articleDetail ArticleDetail

	db.Where("id = ?", id).Find(&articleDetail)

	return Ok(articleDetail)
}
