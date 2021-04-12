package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"little-blog/router"
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

	router.ApplyRouter(app)
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
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if ctx.Request().Method == iris.MethodOptions {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.Next()
}
