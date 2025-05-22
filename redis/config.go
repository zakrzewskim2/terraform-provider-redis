package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/mediocregopher/radix/v4"
)

type RedisClient struct {
	Client *redis.Client
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	addr := d.Get("address").(string)
	password := d.Get("password").(string)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Ping to verify connection
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to connect to Redis: %w", err))
	}

	return &RedisClient{Client: rdb}, diags
}
