package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	. "little-blog/entity"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Article struct {
	BaseEntity
	Topic        string
	Thumbnail    string
	Content      string
	PublishState int
	PublishTime  time.Time
}

func FetchArticleList(ctx iris.Context) {
	var list = []Article{
		{
			BaseEntity: BaseEntity{
				Id:         "1234",
				CreateUser: "az",
				CreateTime: time.Now(),
			},
			Topic:        "haha",
			Thumbnail:    "http://azusahin.cn/a.jpg",
			Content:      "Go 语言 的 数组 的初始化，即数组定义的时候给数组赋初值，一共可分为四种方法",
			PublishState: 0,
			PublishTime:  time.Now(),
		},
		{
			BaseEntity: BaseEntity{
				Id:         "1234",
				CreateUser: "az",
				CreateTime: time.Now(),
			},
			Topic:        "haha",
			Thumbnail:    "http://azusahin.cn/a.jpg",
			Content:      "Go 语言 的 数组 的初始化，即数组定义的时候给数组赋初值，一共可分为四种方法",
			PublishState: 0,
			PublishTime:  time.Now(),
		},
	}
	_, _ = ctx.JSON(list)
}

func fetchArticleList() []Article {
	mysqlDialector := mysql.Dialector{}
	mysqlConfig := gorm.Config{}
	db, err := gorm.Open(mysqlDialector, &mysqlConfig)
	if err != nil {
		panic(err)
	}
	var article Article

	fmt.Println(db.Having(&article))

	return nil
}