package handler

import (
	"github.com/kataras/iris/v12"
	. "little-blog/common"
	. "little-blog/model"
	"net/http"
	"time"
)

// ArticleHandler 列表控制器
type ArticleHandler struct {
}

// ArticleDetailHandler 详情控制器
type ArticleDetailHandler struct {
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
	return ""
}

func (_ *ArticleDetailHandler) Method() string {
	return http.MethodGet
}

func (_ *ArticleDetailHandler) HandlerFunc(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, _ = ctx.JSON(*fetchArticleDetail(id))
}

func (_ *ArticleDetailHandler) Path() string {
	return "/{id}"
}

func fetchArticleList(pageNum, pageSize int) *RestResponse {
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()

	// 5秒超时
	timer := time.NewTimer(5 * time.Second)
	ch := make(chan []Article)

	var articles []Article
	go func() {

		// 查找已发布且未删除的文章
		db.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("publish_state = ? AND is_delete = ?", 1, 0).Find(&articles)
		ch <- articles
	}()

	select {
	case articles = <-ch:
		return Ok(articles)
	case <-timer.C:
		return Error(http.StatusRequestTimeout, "查询超时")
	}

}

func fetchArticleDetail(id string) *RestResponse {
	db, cleanFunc, _ := GetDb()
	defer cleanFunc()
	var articleDetail ArticleDetail

	db.Where("id = ?", id).Find(&articleDetail)

	return Ok(articleDetail)
}
