//go:generate mapstructure-to-hcl2 -type PostProcessorConfig

package main

import (
	"context"
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type PostProcessorConfig struct {
	Comment   string `mapstructure:"comment"`
	SendToUi  bool   `mapstructure:"ui"`
	Bubble    bool   `mapstructure:"bubble_text"`
	PackerSay bool   `mapstructure:"packer_say"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config PostProcessorConfig
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)
	if err != nil {
		return err
	}

	if p.config.PackerSay && p.config.Bubble {
		return fmt.Errorf("Can't have both packer_say and bubble_text options set.")
	}

	return nil
}

func (p *PostProcessor) PostProcess(_ context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {
	comment, err := interpolate.Render(p.config.Comment, &p.config.ctx)
	if err != nil {
		return artifact, true, false, fmt.Errorf("Error interpolating comment: %s", err)
	}

	if p.config.SendToUi {
		if p.config.Bubble {
			myFigure := figure.NewFigure(comment, "", false)
			ui.Say(myFigure.String())
		} else if p.config.PackerSay {
			// CreatePackerFriend is defined in happy_packy.go
			packyText, err := CreatePackerFriend(comment)
			if err != nil {
				return artifact, true, false, err
			}
			ui.Say(packyText)
		} else {
			ui.Say(comment)
		}

	}

	return artifact, true, false, nil
}
