// main.go
package main

import (
	provider "github.com/HappyPathway/terraform-provider-configstash"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
