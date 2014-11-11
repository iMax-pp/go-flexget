// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"net/http"
)

// Command to retrieve FlexGet logs (only keep 100 last lines)
var getLogsCmd = "tail -100 .flexget/flexget.log"

// '/api/logs' request handler. Store FlexGet data in cache
func LogsHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("LogsHandler")

	var body string
	data, exist := fgCache.Get("logs")
	if exist {
		logger.Debug("Retrieve FlexGet logs from cache")
		body = data.(string)
	} else {
		logger.Debug("Retrieve FlexGet logs from server")
		var err error
		if body, err = ExecSSHCmd(getLogsCmd); err != nil {
			http.Error(w, err.Error(), 500)
		}
		fgCache.Add("logs", body, 0)
	}

	fmt.Fprint(w, body)
	logger.TraceEnd("LogsHandler")
}
