// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	utils "github.com/iMax-pp/go-utils"
	"net/http"
)

var logger *utils.Logger
var props map[string]string

func main() {
	// Init Logger
	logger, _ = utils.NewLoggerFromConfig("logging.properties")
	defer logger.Close()
	// Init Application properties
	props, _ = utils.LoadConfig("application.properties")

	// Serve static content
	http.Handle("/", http.FileServer(http.Dir("views")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// Serve dynamic content
	http.Handle("/api/logs", http.HandlerFunc(LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(ConfigHandler))

	// Up and listening
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		logger.Fatal("ListenAndServe:", err)
	}
}

// Command to retrieve FlexGet logs (only keep 100 last lines)
var getLogsCmd = "tail -100 .flexget/flexget.log"

func LogsHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("LogsHandler")
	body, err := ExecSSHCmd(getLogsCmd)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprint(w, body)
	logger.TraceEnd("LogsHandler")
}

// Command to retrieve FlexGet configuration
var getConfigCmd = "cat .flexget/config.yml"

func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("ConfigHandler")
	body, err := ExecSSHCmd(getConfigCmd)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprint(w, body)
	logger.TraceEnd("ConfigHandler")
}
