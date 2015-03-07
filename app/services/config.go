// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package services

import (
	"fmt"
	"github.com/golang/glog"
	. "github.com/iMax-pp/go-flexget/app/common"
	"net/http"
)

// Command to retrieve FlexGet configuration
const getConfigCmd = "cat .flexget/config.yml"
const CACHE_CONFIG_KEY = "config"

// '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	if data, exist := Cache().Get(CACHE_CONFIG_KEY); exist {
		glog.Info("Retrieve FlexGet config from cache")
		fmt.Fprint(w, data.(string))
	} else {
		glog.Info("Retrieve FlexGet config from server")
		if body, err := ExecSSHCmd(getConfigCmd); err != nil {
			glog.Error("Error retrieving FlexGet config: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			Cache().Add(CACHE_CONFIG_KEY, body, 0)
			fmt.Fprint(w, body)
		}
	}
}
