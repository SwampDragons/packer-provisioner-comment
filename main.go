package main

import (
	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterProvisioner(new(CommentProvisioner))
	server.Serve()
}
