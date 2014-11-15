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

const (
	// Configuration files
	LOG_PROPS_FILE = "logging.properties"
	APP_PROPS_FILE = "application.properties"
	// Configuration keys
	CONF_SERVER_PORT    = "server.port"
	CONF_FG_SSH_SERVER  = "flexget.ssh.server"
	CONF_FG_SSH_USER    = "flexget.ssh.user"
	CONF_FG_SSH_PRIVKEY = "flexget.ssh.privatekey"
	CONF_CACHE_EXPIR    = "cache.expiration"
	CONF_CACHE_CLEAN    = "cache.cleanup"
)

func main() {
	// Init Logger
	logger, _ = utils.NewLoggerFromConfig(LOG_PROPS_FILE)
	defer logger.Close()
	// Init Application properties
	props, _ = utils.LoadConfig(APP_PROPS_FILE)
	// Init FlexGet Cache
	expir, _ := strconv.Atoi(props[CONF_CACHE_EXPIR])
	clean, _ := strconv.Atoi(props[CONF_CACHE_CLEAN])
	fgCache = cache.New(time.Duration(expir)*time.Second, time.Duration(clean)*time.Second)

	// Serve static content
	http.Handle("/", http.FileServer(http.Dir("views")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// Serve dynamic content
	http.Handle("/api/status", http.HandlerFunc(StatusHandler))
	http.Handle("/api/logs", http.HandlerFunc(LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(ConfigHandler))

	// Up and listening
	err := http.ListenAndServe(":"+props[CONF_SERVER_PORT], nil)
	if err != nil {
		logger.Fatal("ListenAndServe:", err)
	}
}
