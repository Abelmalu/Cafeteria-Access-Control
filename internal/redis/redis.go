package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✔ Redis connected successfully!")
}
