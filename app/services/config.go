// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package services

import (
	"fmt"
	"github.com/golang/glog"
	common "github.com/iMax-pp/go-flexget/app/common"
	"net/http"
)

const (
	// Command to retrieve FlexGet configuration
	getConfigCmd   = "cat .flexget/config.yml"
	cacheConfigKey = "config"
)

// ConfigHandler '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	if data, exist := common.Cache().Get(cacheConfigKey); exist {
		glog.Info("Retrieve FlexGet config from cache")
		fmt.Fprint(w, data.(string))
	} else {
		glog.Info("Retrieve FlexGet config from server")
		if body, err := common.ExecSSHCmd(getConfigCmd); err != nil {
			glog.Error("Error retrieving FlexGet config: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			common.Cache().Add(cacheConfigKey, body, 0)
			fmt.Fprint(w, body)
		}
	}
}
