package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type BroadcastCache interface {
	IsActive(postID uint) (bool, bool)
	SetInActive(postID uint)
}

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(rdb *redis.Client) *RedisCache {
	return &RedisCache{rdb: rdb}
}

func key(postID uint) string {
	return fmt.Sprintf("broadcast:%d:active", postID)
}

func (c *RedisCache) IsActive(postID uint) (bool, bool) {
	ctx := context.Background()
	val, err := c.rdb.Get(ctx, key(postID)).Result()
	if err == redis.Nil {
		return false, false
	}

	if err != nil {
		return false, false
	}
	return val == "1", true
}

func (c *RedisCache) SetInActive(postID uint) {
	ctx := context.Background()
	c.rdb.Set(ctx, key(postID), "0", 24*time.Hour)

}
