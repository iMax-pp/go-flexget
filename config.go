// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("ConfigHandler")

	body, err := ioutil.ReadFile(props["flexget.path"] + "/config.yml")
	if err != nil {
		logger.Error("ConfigHandler", err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(body))
	logger.TraceEnd("ConfigHandler")
}
