package main

import (
	"fmt"
	"github.com/hunterlong/statup/plugin"
	"net/http"
)

func (p *pkg) Routes() []plugin.Routing {
	fmt.Println("Statup Example Plugin routing!")
	return []plugin.Routing{{
		URL:     "hello",
		Method:  "GET",
		Handler: CustomInfoHandler,
	}}
}

func CustomInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is totally custom!")
}
