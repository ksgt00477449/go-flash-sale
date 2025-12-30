package cache

import (
	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	redis redis.UniversalClient
}

func NewUserCache(client redis.UniversalClient) *UserCache {
	return &UserCache{
		redis: client,
	}
}
