package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"little-blog/controller"
	"net/http"
	"time"
)

/**
 * 博客启动主方法
 */
func main() {
	app := iris.Default()
	app.Use(loggerMiddleware)

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		_ = app.Shutdown(ctx)
	})

	var baseApi = "/api/v1"

	app.Get(baseApi+"/article", controller.FetchArticleList)

	_ = app.Build()

	srv := &http.Server{Handler: app, Addr: ":8080"}
	println("Start a server listening on http://localhost:8080")
	_ = srv.ListenAndServe()

}

func loggerMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Run before %s", ctx.Path())
	ctx.Next()
}
