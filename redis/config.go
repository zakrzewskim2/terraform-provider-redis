package redis

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfig() map[string]*schema.Schema {
    return map[string]*schema.Schema{
        "address": {
            Type:     schema.TypeString,
            Required: true,
        },
        "password": {
            Type:     schema.TypeString,
            Optional: true,
            Sensitive: true,
        },
    }
}
