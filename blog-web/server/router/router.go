package router

import (
	"github.com/kataras/iris/v12"
	"little-blog/handler"
)

func ApplyRouter(app *iris.Application) {
	article := app.Party("/api/v1/article")
	{
		articleHandler := new(handler.ArticleHandler)
		articleDetailHandler := new(handler.ArticleDetailHandler)
		commentListHandler := new(handler.CommentListHandler)
		commentHandler := new(handler.CommentHandler)
		apply(article, articleHandler, articleDetailHandler, commentListHandler, commentHandler)
	}
}

func apply(party iris.Party, handlers ...handler.Handler) {
	if len(handlers) > 0 {
		for _, h := range handlers {
			party.Handle(h.Method(), h.Path(), h.HandlerFunc)
		}
	}
}
