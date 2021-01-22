package entity

import "time"

type BaseEntity struct {
	Id         string `gorm:"primary_key"`
	CreateUser string
	CreateTime time.Time
	UpdateUser string
	UpdateTime time.Time
	IsDelete   int
}
