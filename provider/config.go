package provider

import (
	"context"
	"fmt"

	"github.com/mediocregopher/radix/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(string)
	database := d.Get("database").(string)

	var diags diag.Diagnostics

	cfg := radix.PoolConfig{
		Dialer: radix.Dialer{
			SelectDB: database,
		},
	}

	client, err := cfg.New(ctx, "tcp", fmt.Sprintf("%s:%s", hostname, port))
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}