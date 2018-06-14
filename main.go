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
			Name:        "Example Plugin",
			Description: "This is an example plugin for Statup Status Page application. It can be implemented pretty quick!",
			Form: []*plugin.FormElement{{
				Name:        "example",
				Description: "insert a cool example that will show below an input",
				InputName:   "example",
				InputType:   "text",
				Value:       "example value here",
			}, {
				Name:      "awesome testing",
				InputName: "awesome",
				InputType: "text",
				Value:     "55123",
			}},
		},
	}

}
