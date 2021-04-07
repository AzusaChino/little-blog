package handler

import (
    "github.com/kataras/iris/v12"
    . "little-blog/common"
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

type CommentListHandler struct {
}

type CommentHandler struct {
}

func (_ *Comment) TableName() string {
    return "tb_comment"
}

func (_ *CommentListHandler) Method() string {
    return http.MethodGet
}

func (_ *CommentListHandler) HandlerFunc(ctx iris.Context) {
    id := ctx.Params().Get("id")
    _, _ = ctx.JSON(fetchArticleCommentList(id))
}

func (_ *CommentListHandler) Path() string {
    return "api/v1/article/{id}/comment"
}

func (_ *CommentHandler) Method() string {
    return http.MethodPost
}

func (_ *CommentHandler) HandlerFunc(ctx iris.Context) {
    id := ctx.Params().Get("id")
    _, _ = ctx.JSON(fetchComment(id))
}

func (_ *CommentHandler) Path() string {
    return "api/v1/comment/{id}"
}

// 仅查询当前文章的第一层评论 TODO
func fetchArticleCommentList(articleId string) []Comment {
    db, cleanFunc, _ := GetDb()
    defer cleanFunc()

    var comments []Comment

    db.Where("article_id = ?", articleId).Find(&comments)

    return comments
}

// 查询当前评论的评论
func fetchComment(commentId string) []Comment {
    db, cleanFunc, _ := GetDb()
    defer cleanFunc()

    var comments []Comment

    db.Where("pid = ?", commentId).Find(&comments)

    return comments
}
