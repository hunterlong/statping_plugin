package main

import (
	"fmt"
	"github.com/hunterlong/statup/plugin"
)

func (p pkg) GetInfo() plugin.Info {
	return Plugin.Info
}

func (p pkg) OnLoad() {
	fmt.Println("Statup Example Plugin Loaded")
	plugin.CreateSettingsTable(p.Info)
}

func (p pkg) Form() string {
	return "none"
}

func (p pkg) OnInstall() {
	fmt.Println("Statup Example Plugin installing!")
}

func (p pkg) OnUninstall() {
	fmt.Println("Statup Example Plugin uninstalling!")
}

func (p pkg) OnSave(data map[string]string) {
	//example := data["example"]
	//webhook := data["webhook"]
	//awesome := data["awesome"]

	//_, err := plugin.DB.Exec("UPDATE plugin_example SET example=$1, webhook=$2, awesome=$3;", example, webhook, awesome)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("Statup Example Plugin received updated data! ", data)
}

func (p pkg) OnNewService() {
	fmt.Println("Statup Example Plugin received a new service!")
}

func (p pkg) OnHit(s *plugin.Service) {
	//fmt.Println("Statup Example Plugin received a hit! ", s)
}

func (p pkg) OnFailure(s *plugin.Service) {
	//fmt.Println("Statup Example Plugin received a failure! ", s)
}

func (p pkg) OnSettingsSaved() {
	fmt.Println("Statup Example Plugin received a save on settings!")
}

func (p pkg) OnNewUser() {
	fmt.Println("Statup Example Plugin received a new user!")
}

func (p *pkg) OnShutdown() {
	fmt.Println("Statup Example Plugin received a shutdown event!")
}

func (p pkg) OnBeforeRequest() {
	fmt.Println("Statup Example Plugin received before service request!")
}

func (p pkg) OnAfterRequest() {
	fmt.Println("Statup Example Plugin received after service request!")
}
