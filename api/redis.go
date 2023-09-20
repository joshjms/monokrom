package api

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}
