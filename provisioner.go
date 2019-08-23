package main

import (
	"context"

	"github.com/common-nighthawk/go-figure"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Comment  string `mapstructure:"comment"`
	SendToUi bool   `mapstructure:"ui"`
	Fancy    bool   `mapstructure:"fancy"`

	ctx interpolate.Context
}

type CommentProvisioner struct {
	config Config
}

func (p *CommentProvisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)
	if err != nil {
		return err
	}

	return nil
}

func (p *CommentProvisioner) Provision(_ context.Context, ui packer.Ui, comm packer.Communicator) error {
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
