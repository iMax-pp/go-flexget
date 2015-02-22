// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"net/http"
)

// Command to retrieve FlexGet configuration
const getConfigCmd = "cat .flexget/config.yml"
const CACHE_CONFIG_KEY = "config"

// '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler() (int, string) {
	data, exist := fgCache.Get(CACHE_CONFIG_KEY)
	if exist {
		logger.Debug("Retrieve FlexGet config from cache")
		return http.StatusOK, data.(string)
	} else {
		logger.Debug("Retrieve FlexGet config from server")
		if body, err := ExecSSHCmd(getConfigCmd); err != nil {
			return http.StatusInternalServerError, err.Error()
		} else {
			fgCache.Add(CACHE_CONFIG_KEY, body, 0)
			return http.StatusOK, body
		}
	}
}
