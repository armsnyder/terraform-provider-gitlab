package main

import (
	"fmt"

	"github.com/gitlabhq/terraform-provider-gitlab/gitlab"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	fmt.Println("Hello world")
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gitlab.Provider})
}
