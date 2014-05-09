//
// server bin
//
// Gets called by Google if the spreadsheets are updated
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package main

import (
	"encoding/json"
	"fmt"
	"github.com/dothiv/translations-updater/config"
	"github.com/dothiv/translations-updater/server"
	"net/http"
	"os"
)

func main() {
	cfg_file := "example/config.json"
	if len(os.Args) > 1 {
		cfg_file = os.Args[1]
	}
	cfg, err := os.Open(cfg_file)
	if err != nil {
		os.Stdout.WriteString("Must provide config")
		os.Exit(1)
	}

	decoder := json.NewDecoder(cfg)
	configuration := config.Configuration{}
	cfg_err := decoder.Decode(&configuration)

	if cfg_err != nil {
		os.Stdout.WriteString("Invalid config!\n")
		os.Stdout.WriteString(cfg_err.Error())
		os.Exit(1)
	}

	hookHandler := server.NewHookHandler(configuration)

	http.Handle("/hook", hookHandler)
	os.Stdout.WriteString(fmt.Sprintf("Starting server at port %d ...\n", configuration.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
