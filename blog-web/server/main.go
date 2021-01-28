package main

import (
	"context"
	"github.com/kataras/iris/v12"
	. "little-blog/handler"
	"net/http"
	"time"
)

/**
 * 博客启动主方法
 */
func main() {
	app := iris.Default()

	// 日志
	app.Use(loggerHandler)
	// 允许跨域
	app.Use(corsHandler)
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		_ = app.Shutdown(ctx)
	})

	process(app, &ArticleHandler{}, &ArticleDetailHandler{}, &CommentHandler{})
	_ = app.Build()

	srv := &http.Server{Handler: app, Addr: ":8080"}
	println("Start a server listening on http://localhost:8080")
	_ = srv.ListenAndServe()

}

func loggerHandler(ctx iris.Context) {
	ctx.Application().Logger().Infof("Run before %s", ctx.Path())
	ctx.Next()
}

func corsHandler(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func process(app *iris.Application, handlers ...Handler) {
	if len(handlers) <= 1 {
		return
	}
	for _, handler := range handlers {
		app.Handle(handler.Method(), handler.Path(), handler.HandlerFunc)
	}
}
