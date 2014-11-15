// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"net/http"
)

// Command to retrieve FlexGet configuration
const getConfigCmd = "cat .flexget/config.yml"
const CACHE_CONFIG_KEY = "config"

// '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("ConfigHandler")

	var body string
	data, exist := fgCache.Get(CACHE_CONFIG_KEY)
	if exist {
		logger.Debug("Retrieve FlexGet config from cache")
		body = data.(string)
	} else {
		logger.Debug("Retrieve FlexGet config from server")
		var err error
		if body, err = ExecSSHCmd(getConfigCmd); err != nil {
			http.Error(w, err.Error(), 500)
		}
		fgCache.Add(CACHE_CONFIG_KEY, body, 0)
	}

	fmt.Fprint(w, body)
	logger.TraceEnd("ConfigHandler")
}
