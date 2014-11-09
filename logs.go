// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LogsHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("LogsHandler")

	body, err := ioutil.ReadFile(props["flexget.path"] + "/flexget.log")
	if err != nil {
		logger.Error("LogsHandler", err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(body))
	logger.TraceEnd("LogsHandler")
}
