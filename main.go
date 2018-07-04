package main

import (
	"github.com/hunterlong/statup/plugin"
)

type pkg plugin.PluginInfo

var (
	Plugin  pkg
	VERSION string
)

func main() {

}

func init() {
	Plugin = pkg{
		Info: plugin.Info{
			Name:        "Example Plugin",
			Description: "This is an example plugin for Statup Status Page application. It can be implemented pretty quick!",
		},
	}
}
