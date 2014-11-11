// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"net/http"
)

// Command to retrieve FlexGet configuration
var getConfigCmd = "cat .flexget/config.yml"

// '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("ConfigHandler")

	var body string
	data, exist := fgCache.Get("config")
	if exist {
		logger.Debug("Retrieve FlexGet config from cache")
		body = data.(string)
	} else {
		logger.Debug("Retrieve FlexGet config from server")
		var err error
		if body, err = ExecSSHCmd(getConfigCmd); err != nil {
			http.Error(w, err.Error(), 500)
		}
		fgCache.Add("config", body, 0)
	}

	fmt.Fprint(w, body)
	logger.TraceEnd("ConfigHandler")
}
