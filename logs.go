// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"net/http"
)

// Command to retrieve FlexGet logs (only keep 100 last lines)
const getLogsCmd = "tail -100 .flexget/flexget.log"
const CACHE_LOGS_KEY = "logs"

// '/api/logs' request handler. Store FlexGet data in cache
func LogsHandler() (int, string) {
	data, exist := fgCache.Get(CACHE_LOGS_KEY)
	if exist {
		logger.Debug("Retrieve FlexGet logs from cache")
		return http.StatusOK, data.(string)
	} else {
		logger.Debug("Retrieve FlexGet logs from server")
		if body, err := ExecSSHCmd(getLogsCmd); err != nil {
			return http.StatusInternalServerError, err.Error()
		} else {
			fgCache.Add(CACHE_LOGS_KEY, body, 0)
			return http.StatusOK, body
		}
	}
}
