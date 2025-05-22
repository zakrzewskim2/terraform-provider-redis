// redis/provider.go
package redis

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Sensitive: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"redis_user": resourceRedisUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}