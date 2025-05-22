// redis/config.go
package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	addr := d.Get("address").(string)
	pass := d.Get("password").(string)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return &RedisClient{Client: rdb}, nil
}