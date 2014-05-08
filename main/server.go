package main

import (
	"encoding/json"
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
	configuration := server.Configuration{}
	cfg_err := decoder.Decode(&configuration)

	if cfg_err != nil {
		os.Stdout.WriteString("Invalid config!\n")
		os.Stdout.WriteString(cfg_err.Error())
		os.Exit(1)
	}

	server := server.NewHookServer(configuration)

	http.Handle("/hook", server)
	http.ListenAndServe(":8080", nil)
}
