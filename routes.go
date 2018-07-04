package main

import (
	"fmt"
	"github.com/hunterlong/statup/plugin"
	"net/http"
)

// Create an array of custom HTTP Routes
// this route will allow the Statup server to route "http://domain.com/hello"
// to my "CustomInfoHandler" Handler.
func (p *pkg) Routes() []plugin.Routing {
	return []plugin.Routing{{
		URL:     "hello",
		Method:  "GET",
		Handler: CustomInfoHandler,
	}}
}

// This is the HTTP handler for the '/hello' URL created above
func CustomInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Oh Wow!!! I can create plugins with database interactions and HTTP Handlerz!?!?! OH BOY!!!")
}
