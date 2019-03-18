package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/samueldumont/terraform-provider-rabbitmq/rabbitmq"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rabbitmq.Provider})
}
