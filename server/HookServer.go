package server

import (
	"fmt"
	"net/http"
)

type Configuration struct {
	Sites map[string]interface{} `json:"sites"`
}

func (c *Configuration) GetSiteUri(site string) (uri string, err error) {
	// m := c.sites.(map[string]interface{})
	for k, v := range c.Sites {
		switch vv := v.(type) {
		case string:
			if k == site {
				uri = vv
				return
			}
		}
	}
	err = fmt.Errorf("Site not configured: %s", site)
	return
}

type HookServer struct {
	Config Configuration
}

func (e *HookServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO: get site from request
	site := "charity"
	uri, err := e.Config.GetSiteUri(site)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte("ERROR\n"))
	}
	go func(uri string) {
		l, csverr := NewCsvFileLoader(uri)
		if csverr != nil {
			return
		}
		entries, readerr := l.Reader.ReadAll()
		if readerr != nil {
			return
		}
		for _, v := range entries {
			print("+++++\n")
			for kk, vv := range v {
				print(string(kk) + ":" + vv + "\n")
			}
		}

	}(uri)
	response.Write([]byte("Fetching strings from " + uri + ". Bye.\n"))
}

func NewHookServer(cfg Configuration) (hs *HookServer) {
	hs = new(HookServer)
	hs.Config = cfg
	return
}
