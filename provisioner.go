package main

import (
	"context"

	"github.com/common-nighthawk/go-figure"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

type Config struct {
	Comment  string `mapstructure:"comment"`
	SendToUi bool   `mapstructure:"ui"`
	Fancy    bool   `mapstructure:"fancy"`
}

type CommentProvisioner struct {
	config Config
}

func (p *CommentProvisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate: true,
	}, raws...)
	if err != nil {
		return err
	}

	return nil
}

func (p *CommentProvisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator) error {
	if p.config.SendToUi {
		if p.config.Fancy {
			myFigure := figure.NewFigure(p.config.Comment, "", false)
			ui.Say(myFigure.String())
		} else {
			ui.Say(p.config.Comment)
		}
	}

	return nil
}
