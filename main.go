package main

import (
	"packer-plugin-comment/comment"

	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	p := new(comment.Provisioner)
	server.RegisterProvisioner(p)
	server.Serve()
}
