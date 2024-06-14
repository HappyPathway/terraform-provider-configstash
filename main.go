package main

import (
	"context"

	"github.com/google/go-github/v33/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"golang.org/x/oauth2"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repo": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: d.Get("token").(string)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client, err := github.NewEnterpriseClient(d.Get("url").(string), d.Get("url").(string), tc)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
