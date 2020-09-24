package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// redisRepository to store redis client
type redisRepository struct {
	client *redis.Client
}

// newRedisClient returns a redis client
func newRedisClient(redisURL string) (*redis.Client, error) {
	ctx := context.Background()
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
