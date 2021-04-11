package common

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	rdb := GetRedisClient()
	if rdb == nil {
		t.Error("redis连接失败")
	}
}

func TestGetRedisClient2(t *testing.T) {
	rdb := GetRedisClient()
	val, err := rdb.Get(nil, "key").Result()
	switch {
	case err == redis.Nil:
		t.Error("key does not exist")
	case err != nil:
		t.Error("Get failed", err)
	case val == "":
		t.Error("value is empty")
	}
}
