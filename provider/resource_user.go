// redis/resource_user.go
package provider

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRedisUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisUserCreate,
		Read:   resourceRedisUserRead,
		Update: resourceRedisUserUpdate,
		Delete: resourceRedisUserDelete,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"commands": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func buildACLArgs(d *schema.ResourceData) []string {
	var acl []string

	if d.Get("enabled").(bool) {
		acl = append(acl, "on")
	} else {
		acl = append(acl, "off")
	}

	pw := d.Get("password").(string)
	acl = append(acl, fmt.Sprintf(">%s", pw))

	for _, c := range d.Get("commands").([]interface{}) {
		acl = append(acl, fmt.Sprintf("commands|%s", c.(string)))
	}

	for _, k := range d.Get("keys").([]interface{}) {
		acl = append(acl, fmt.Sprintf("keys|%s", k.(string)))
	}

	return acl
}

func stringSliceToInterfaceSlice(s []string) []interface{} {
	out := make([]interface{}, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}

func resourceRedisUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Get("username").(string)
	aclArgs := buildACLArgs(d)

	// ACL SETUSER <user> <args...>
	cmd := append([]interface{}{"ACL", "SETUSER", user}, stringSliceToInterfaceSlice(aclArgs)...)
	_, err := client.Do(ctx, cmd...).Result()
	if err != nil {
		return err
	}

	d.SetId(user)
	return resourceRedisUserRead(d, m)
}

func resourceRedisUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Id()
	_, err := client.Do(ctx, "ACL", "GETUSER", user).Result()
	if err == redis.Nil {
		d.SetId("") // user doesn't exist
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func resourceRedisUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Id()
	aclArgs := buildACLArgs(d)

	cmd := append([]interface{}{"ACL", "SETUSER", user}, stringSliceToInterfaceSlice(aclArgs)...)
	_, err := client.Do(ctx, cmd...).Result()
	if err != nil {
		return err
	}

	return resourceRedisUserRead(d, m)
}

func resourceRedisUserDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Id()
	_, err := client.Do(ctx, "ACL", "DELUSER", user).Result()
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
