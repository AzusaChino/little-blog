package routers

import "github.com/kataras/iris/v12"
import . "little-blog/handler"

// init all routers by router group

func Router(app *iris.Application) {
	v1 := app.Party("/api/v1/article")
	{
		articleHandler := &ArticleHandler{}
		v1.Handle(articleHandler.Method(), articleHandler.Path(), articleHandler.HandlerFunc)
	}
}
