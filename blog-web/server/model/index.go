package model

import "time"

// baseEntity 实例父类
type baseEntity struct {
	Id         string `gorm:"primary_key"`
	CreateUser string
	CreateTime time.Time
	UpdateUser string
	UpdateTime time.Time
	IsDelete   int
}
