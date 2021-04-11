package common

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// redis://<user>:<password>@<host>:<port>/<db_number>
const redisUrl = "redis://180.76.169.35:6379/"

var c *redis.Client
var Ctx context.Context

func init() {
	doInitRedis()
}

func GetRedisClient() *redis.Client {
	if c == nil {
		doInitRedis()
	}
	return c
}

func doInitRedis() {
	var lock = NewSpinLock()
	lock.Lock()
	defer lock.Unlock()

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	c = redis.NewClient(opt)
	Ctx = context.TODO()
}
