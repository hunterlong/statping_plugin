package main

import (
	"github.com/hunterlong/statup/plugin"
)

type pkg plugin.PluginInfo

var (
	Plugin  pkg
	VERSION string
)

func init() {

	Plugin = pkg{
		Info: plugin.Info{
			Name: "example",
			Description: "this is an example plugin for Statup Status Page application.",
			Form: []plugin.FormElement{{
				Name:        "example",
				Description: "fill out a cool example",
				SQLValue:    "example",
				SQLType:     "text",
			}, {
				Name:        "webhook url",
				Description: "the webhook url can be here",
				SQLValue:    "webhook",
				SQLType:     "text",
			}, {
				Name:        "awesome testing",
				SQLValue:    "awesome",
				SQLType:     "text",
			}},
		},
	}

}
