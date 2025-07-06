package redis

import (
	"context"
	"time"
	"url_shortening/config/environment"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(config *environment.Config) (*Redis, error) {
	opt, err := redis.ParseURL(config.REDIS.Address)
	if err != nil {
		return nil, err
	}

	return &Redis{Client: redis.NewClient(opt)}, nil
}

func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(context.Background(), key, value, expiration).Err()
}
