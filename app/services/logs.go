// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package services

import (
	"fmt"
	"github.com/golang/glog"
	common "github.com/iMax-pp/go-flexget/app/common"
	"net/http"
)

const (
	// Command to retrieve FlexGet logs (only keep 100 last lines)
	getLogsCmd   = "tail -100 .flexget/flexget.log"
	cacheLogsKey = "logs"
)

// LogsHandler '/api/logs' request handler. Store FlexGet data in cache
func LogsHandler(w http.ResponseWriter, req *http.Request) {
	if data, exist := common.Cache().Get(cacheLogsKey); exist {
		glog.Info("Retrieve FlexGet logs from cache")
		fmt.Fprint(w, data.(string))
	} else {
		glog.Info("Retrieve FlexGet logs from server")
		if body, err := common.ExecSSHCmd(getLogsCmd); err != nil {
			glog.Error("Error retrieving FlexGet logs: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			common.Cache().Add(cacheLogsKey, body, 0)
			fmt.Fprint(w, body)
		}
	}
}
