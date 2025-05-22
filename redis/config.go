// redis/config.go
package redis

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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