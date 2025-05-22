package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/mediocregopher/radix/v4"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"hostname": {
					Description: "Server hostname. Can be specified with the `REDISDB_HOSTNAME` environment variable. Defaults to `127.0.0.1`",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("REDISDB_HOSTNAME", "127.0.0.1"),
				},
				"port": {
					Description: "Server port. Can be specified with the `REDISDB_PORT` environment variable. Defaults to `6379`",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("REDISDB_PORT", "6379"),
				},
				"database": {
					Description: "Database number. Can be specified with the `REDISDB_DATABASE` environment variable. Defaults to `0`",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("REDISDB_DATABASE", "0"),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
			"redis_user": resourceRedisUser(),
			},
			ConfigureContextFunc: providerConfigure,
		}
	}
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