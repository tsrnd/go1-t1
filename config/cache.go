package config

import (
	"os"

	"goweb1/services/cache"
	"goweb1/services/cache/redis"
)

// Cache func
func Cache() cache.Cache {
	return redis.Connect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
}
