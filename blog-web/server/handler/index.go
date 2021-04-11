package handler

import "github.com/kataras/iris/v12"

/**
 * URL控制器
 */
type Handler interface {
	Method() string
	Path() string
	HandlerFunc(ctx iris.Context)
}
