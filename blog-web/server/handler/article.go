package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	. "little-blog/common"
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
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()

	var articles []Article

	// 查找已发布且未删除的文章
	db.Where("publish_state = ? AND is_delete = ?", 1, 0).Find(&articles)

	return articles
}

func fetchArticleDetail(id string) ArticleDetail {
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()
	var articleDetail ArticleDetail

	db.Where("id = ?", id).Find(&articleDetail)

	return articleDetail
}

func _fetchArticleDetail(id string) ArticleDetail {
	ctx := context.Background()
	client := GetRedisClient()
	val, err := client.Get(ctx, "").Result()
	switch {
	case err == redis.Nil:
		panic("not found")
	case err != nil:
		panic(err)
	case val == "":
		fmt.Println("value is empty")
	}

	var articleDetail ArticleDetail

	err = json.Unmarshal([]byte(val), &articleDetail)
	if err == nil {
		return articleDetail
	}
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()

	db.Where("id = ?", id).Find(&articleDetail)

	return articleDetail
}
