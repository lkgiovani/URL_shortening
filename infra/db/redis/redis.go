package redis

import (
	"context"
	"time"
	"url_shortening/infra/config/environment"


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

	opt.PoolSize = 10
	opt.MinIdleConns = 5
	opt.MaxRetries = 3
	opt.DialTimeout = 5 * time.Second
	opt.ReadTimeout = 3 * time.Second
	opt.WriteTimeout = 3 * time.Second

	return &Redis{Client: redis.NewClient(opt)}, nil
}

func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(context.Background(), key, value, expiration).Err()
}
