package trial

import (
    "encoding/json"
    "fmt"
    "github.com/go-redis/redis/v8"
    _ "github.com/kataras/iris/v12"
    . "little-blog/common"
)

// demonstrating redis usage
func _fetchArticleDetail(id string) interface{} {
    client := GetRedisClient()
    client.Set(Ctx, "article_detail_"+id, "", 0)
    val, err := client.Get(Ctx, "").Result()
    switch {
    case err == redis.Nil:
        panic("not found")
    case err != nil:
        panic(err)
    case val == "":
        fmt.Println("value is empty")
    }

    var articleDetail interface{}

    err = json.Unmarshal([]byte(val), &articleDetail)
    if err == nil {
        return articleDetail
    }
    db, cleanFunc, _ := GetDb()
    defer cleanFunc()

    db.Where("id = ?", id).Find(&articleDetail)

    return articleDetail
}
