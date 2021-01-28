package handler

import (
	"github.com/kataras/iris/v12"
	. "little-blog/entity"
	"net/http"
)

type Comment struct {
	BaseEntity
	ArticleId string
	Pid       string
	Nickname  string
	Email     string
	Content   string
}

type CommentHandler struct {
}

func (_ *Comment) TableName() string {
	return "tb_comment"
}

func (_ *CommentHandler) Method() string {
	return http.MethodGet
}

func (_ *CommentHandler) HandlerFunc(ctx iris.Context) {
	id := ctx.Params().Get("id")
	_, _ = ctx.JSON(fetchArticleCommentList(id))
}

func (_ *CommentHandler) Path() string {
	return "api/v1/article/{id}/comment"
}

// 仅查询当前文章的第一层评论 TODO
func fetchArticleCommentList(articleId string) []Comment {
	db := GetDb()

	var comments []Comment

	db.Where("article_id = ?", articleId).Find(&comments)

	return comments
}
