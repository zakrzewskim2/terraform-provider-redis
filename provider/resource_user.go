// redis/resource_user.go
package provider

import (
	"context"
	"fmt"
	"strings"

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
				Type:     schema.TypeString,
				Required: true,
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

func buildACLString(d *schema.ResourceData) string {
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

	return strings.Join(acl, " ")
}

func resourceRedisUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Get("username").(string)
	acl := buildACLString(d)

	_, err := client.Do(ctx, "ACL", "SETUSER", user, acl).Result()
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
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceRedisUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*RedisClient).Client
	ctx := context.Background()

	user := d.Id()
	acl := buildACLString(d)

	_, err := client.Do(ctx, "ACL", "SETUSER", user, acl).Result()
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
