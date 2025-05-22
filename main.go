package main

import (
	"context"
	"flag"
	"log"

	"github.com/zakrzewskim2/terraform-provider-redis/redis"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)


var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New(version)}

	plugin.Serve(opts)
}