// provider.go
package main

import (
	"context"

	"github.com/google/go-github/v33/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description: "The GitHub token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"configstash_record": resourceRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client, nil
}
