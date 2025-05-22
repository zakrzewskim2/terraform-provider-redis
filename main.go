// main.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"example.com/terraform-provider-redis/redis"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: redis.Provider,
	})
}