package main

import (
	"context"
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

type Config struct {
	Comment   string `mapstructure:"comment"`
	SendToUi  bool   `mapstructure:"ui"`
	Bubble    bool   `mapstructure:"bubble_text"`
	PackerSay bool   `mapstructure:"packer_say"`
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

	if p.config.PackerSay && p.config.Bubble {
		return fmt.Errorf("Can't have both packer_say and bubble_text options set.")
	}

	return nil
}

func (p *CommentProvisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator) error {
	if p.config.SendToUi {
		if p.config.Bubble {
			myFigure := figure.NewFigure(p.config.Comment, "", false)
			ui.Say(myFigure.String())
		} else if p.config.PackerSay {
			// CreatePackerFriend is being imported from happy_packy.go
			packyText, err := CreatePackerFriend(p.config.Comment)
			if err != nil {
				return err
			}
			ui.Say(packyText)
		} else {
			ui.Say(p.config.Comment)
		}

	}

	return nil
}
