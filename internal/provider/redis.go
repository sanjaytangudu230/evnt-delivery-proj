package provider

import (
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient     *redis.Client
	QueueNames      = []string{"queue1", "queue2", "queue3"}
	Destinations    = []string{"destination1"}
	DeadLetterQueue = "deadLetterQueue"
)

func InitializeRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
	})
}
