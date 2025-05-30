package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				"password": {
					Description: "Password for Redis authentication. Can be set via `REDISDB_PASSWORD`.",
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("REDISDB_PASSWORD", nil),
					Sensitive:   true,
				},
				"tls": {
					Description: "Enable TLS (true/false).",
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
			"redis_user": resourceRedisUser(),
			},
			ConfigureContextFunc: providerConfigure,
		}
	}
}