package main

import (
	"github.com/hashicorp/packer/packer/plugin"
)

// Assume this implements packer.Builder
type Builder struct{}

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterProvisioner(new(Provisioner))
	server.Serve()
}
