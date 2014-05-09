//
// Handles calls to by webhooks
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package server

import (
	"fmt"
	"github.com/dothiv/translations-updater/command"
	"github.com/dothiv/translations-updater/config"
	"net/http"
	"os"
)

type HookHandler struct {
	Config config.Configuration
}

func (e *HookHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO: get site from request
	ident := "charity"
	site, err := e.Config.GetSite(ident)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte("ERROR\n"))
	}

	go func(site *config.Site) {
		c := command.NewImportCommand(site)
		err, errorStrings := c.Exec()
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			for _, v := range errorStrings {
				os.Stderr.WriteString(v.Msg)
			}
		}
	}(&site)

	fmt.Fprintf(response, "Fetching strings for site '%s' from source '%s'.\n", ident, site.Source)
}

func NewHookHandler(cfg config.Configuration) (hs *HookHandler) {
	hs = new(HookHandler)
	hs.Config = cfg
	return
}
