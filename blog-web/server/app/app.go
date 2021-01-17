package app

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type Server struct {
	App *iris.Application
	Status bool
}

var Srv *Server

func NewServer() *Server{
	app := iris.New()
	app.Logger().SetLevel("WARN")

	iris.RegisterOnInterrupt(func() {
		db, err := gorm.DB.
	})
}