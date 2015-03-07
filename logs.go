// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

// Command to retrieve FlexGet logs (only keep 100 last lines)
const getLogsCmd = "tail -100 .flexget/flexget.log"
const CACHE_LOGS_KEY = "logs"

// '/api/logs' request handler. Store FlexGet data in cache
func LogsHandler(w http.ResponseWriter, req *http.Request) {
	if data, exist := fgCache.Get(CACHE_LOGS_KEY); exist {
		glog.Info("Retrieve FlexGet logs from cache")
		fmt.Fprint(w, data.(string))
	} else {
		glog.Info("Retrieve FlexGet logs from server")
		if body, err := ExecSSHCmd(getLogsCmd); err != nil {
			glog.Error("Error retrieving FlexGet logs: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fgCache.Add(CACHE_LOGS_KEY, body, 0)
			fmt.Fprint(w, body)
		}
	}
}
