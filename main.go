// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	utils "github.com/iMax-pp/go-utils"
	cache "github.com/robfig/go-cache"
	"net/http"
	"strconv"
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
