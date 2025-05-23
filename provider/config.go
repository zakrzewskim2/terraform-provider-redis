package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RedisClient struct {
	Client *redis.Client
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(string)
	databaseStr := d.Get("database").(string)
	password := d.Get("password").(string)
	useTLS := d.Get("tls").(bool)

	var diags diag.Diagnostics

	db, err := strconv.Atoi(databaseStr)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("invalid database number: %v", err))
	}

	opts := &redis.Options{
		Addr: fmt.Sprintf("%s:%s", hostname, port),
		DB:   db,
		Password: password,
	}
	
	if useTLS {
		opts.TLSConfig = &tls.Config{}
	}

	client := redis.NewClient(opts)

	// Test the connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, diag.FromErr(fmt.Errorf("unable to connect to Redis: %v", err))
	}

	return &RedisClient{Client: client}, diags
}
