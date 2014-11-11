// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"fmt"
	utils "github.com/iMax-pp/go-utils"
	cache "github.com/robfig/go-cache"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	logger  *utils.Logger
	props   map[string]string
	fgCache *cache.Cache
)

func main() {
	// Init Logger
	logger, _ = utils.NewLoggerFromConfig("logging.properties")
	defer logger.Close()
	// Init Application properties
	props, _ = utils.LoadConfig("application.properties")
	// Init FlexGet Cache
	expir, _ := strconv.Atoi(props["cache.expiration"])
	clean, _ := strconv.Atoi(props["cache.cleanup"])
	fgCache = cache.New(time.Duration(expir)*time.Second, time.Duration(clean)*time.Second)

	// Serve static content
	http.Handle("/", http.FileServer(http.Dir("views")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// Serve dynamic content
	http.Handle("/api/status", http.HandlerFunc(StatusHandler))
	http.Handle("/api/logs", http.HandlerFunc(LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(ConfigHandler))

	// Up and listening
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		logger.Fatal("ListenAndServe:", err)
	}
}

// Command to retrieve FlexGet
var getStatusCmd = "ps | grep flexget | grep -v grep"

// '/api/status' request handler.
func StatusHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("StatusHandler")

	body, err := ExecSSHCmd(getStatusCmd)
	if err != nil && !strings.Contains(err.Error(), "Process exited with: 1") {
		http.Error(w, err.Error(), 500)
	}
	status := body != ""

	fmt.Fprint(w, status)
	logger.TraceEnd("StatusHandler")
}

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

// Command to retrieve FlexGet configuration
var getConfigCmd = "cat .flexget/config.yml"

// '/api/config' request handler. Store FlexGet data in cache
func ConfigHandler(w http.ResponseWriter, req *http.Request) {
	logger.TraceBegin("ConfigHandler")

	var body string
	data, exist := fgCache.Get("config")
	if exist {
		logger.Debug("Retrieve FlexGet config from cache")
		body = data.(string)
	} else {
		logger.Debug("Retrieve FlexGet config from server")
		var err error
		if body, err = ExecSSHCmd(getConfigCmd); err != nil {
			http.Error(w, err.Error(), 500)
		}
		fgCache.Add("config", body, 0)
	}

	fmt.Fprint(w, body)
	logger.TraceEnd("ConfigHandler")
}
