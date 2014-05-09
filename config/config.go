//
// Configuration
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package config

import (
	"fmt"
)

type Site struct {
	Source  string    `json:"source"`
	Targets []*Target `json:"targets"`
}

type Target struct {
	Code   string `json:"code"`
	Val    string `json:"val"`
	Target string `json:"target"`
}

type Configuration struct {
	Port  int             `json:"port"`
	Sites map[string]Site `json:"sites"`
}

func (c *Configuration) GetSite(ident string) (site Site, err error) {
	for k, v := range c.Sites {
		if k == ident {
			site = v
			return
		}
	}
	err = fmt.Errorf("Site not configured: %s", ident)
	return
}
