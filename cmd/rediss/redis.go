package rediss

import "github.com/redis/go-redis/v9"

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "my-redis-container:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
