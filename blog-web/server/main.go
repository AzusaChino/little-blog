package main

import (
	"context"
	"github.com/kataras/iris/v12"
	. "little-blog/handler"
	"log"
	"net/http"
	"time"
)

/**
 * 博客启动主方法
 */
func main() {
	// default contains logger
	app := iris.Default()

	// 允许跨域
	app.Use(corsHandler)
	// graceful shutdown
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		_ = app.Shutdown(ctx)
	})

	process(app, &ArticleHandler{}, &ArticleDetailHandler{}, &CommentListHandler{})
	_ = app.Build()

	srv := &http.Server{Handler: app, Addr: ":8080"}
	log.Default().Println("Start a server listening on http://localhost:8080")
	_ = srv.ListenAndServe()

}

/**
  跨域控件
*/
func corsHandler(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == iris.MethodOptions {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.Next()
}

/**
  handler统一控制
*/
func process(app *iris.Application, handlers ...Handler) {
	if len(handlers) <= 1 {
		return
	}
	for _, handler := range handlers {
		app.Handle(handler.Method(), handler.Path(), handler.HandlerFunc)
	}
}
