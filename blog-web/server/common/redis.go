package common

import "github.com/go-redis/redis/v8"

const RedisUrl = "redis://localhost:6379/"

func GetRedisClient() *redis.Client {
	opt, err := redis.ParseURL(RedisUrl)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}
