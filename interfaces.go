package main

import (
	"fmt"
	"github.com/hunterlong/statup/plugin"
	"upper.io/db.v3/lib/sqlbuilder"
)

func (p pkg) GetInfo() plugin.Info {
	return Plugin.Info
}

func (p pkg) OnLoad(db sqlbuilder.Database) {
	fmt.Println("=============================================================")
	fmt.Printf("  Statup Example Plugin Loaded using %v database\n", db.Name())
	fmt.Println("=============================================================")
}

func (p pkg) OnSave(data map[string]interface{}) {
	example := data["example"]
	webhook := data["webhook"]
	awesome := data["awesome"]
	fmt.Println("Statup Example Plugin received updated data! ", example, webhook, awesome)
}

func (p pkg) OnNewService(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin received a new service!")
}

func (p pkg) OnSuccess(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin received a hit! ")
	fmt.Println("Name:    ", data["Name"])
	fmt.Println("Domain:  ", data["Domain"])
	fmt.Println("Method:  ", data["Method"])
	fmt.Println("Latency: ", data["Latency"])
}

func (p pkg) OnFailure(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin received a failure! ")
	fmt.Println("Name:    ", data["Name"])
	fmt.Println("Domain:  ", data["Domain"])
	fmt.Println("Method:  ", data["Method"])
	fmt.Println("Latency: ", data["Latency"])
}

func (p pkg) OnSettingsSaved(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin received a save on settings!")
}

func (p pkg) OnNewUser(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin received a new user!")
}

func (p pkg) OnInstall(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin installing!")
}

func (p pkg) OnUninstall(data map[string]interface{}) {
	fmt.Println("Statup Example Plugin uninstalling!")
}

func (p pkg) OnBeforeRequest(map[string]interface{}) {
	fmt.Println("Statup Example Plugin received before service request!")
}

func (p pkg) OnAfterRequest(map[string]interface{}) {
	fmt.Println("Statup Example Plugin received after service request!")
}

func (p *pkg) OnShutdown() {
	fmt.Println("Statup Example Plugin received a shutdown event!")
}
