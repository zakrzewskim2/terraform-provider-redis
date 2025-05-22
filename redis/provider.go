// redis/provider.go
package redis

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:        providerConfig(),
		ResourcesMap:  map[string]*schema.Resource{"redis_user": resourceRedisUser()},
		ConfigureContextFunc: providerConfigure,
	}
}