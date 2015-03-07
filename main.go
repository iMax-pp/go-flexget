// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	utils "github.com/iMax-pp/go-utils"
	cache "github.com/robfig/go-cache"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
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

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go-flexget -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Init Application properties
	props, _ = utils.LoadConfig(APP_PROPS_FILE)
	// Init FlexGet Cache
	expir, _ := strconv.Atoi(props[CONF_CACHE_EXPIR])
	clean, _ := strconv.Atoi(props[CONF_CACHE_CLEAN])
	fgCache = cache.New(time.Duration(expir)*time.Second, time.Duration(clean)*time.Second)

	// Service static content
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.Handle("/api/status", http.HandlerFunc(StatusHandler))
	http.Handle("/api/logs", http.HandlerFunc(LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(ConfigHandler))

	// Up and listening
	glog.Info("Will start listening on port ", props[CONF_SERVER_PORT])
	if err := http.ListenAndServe(":"+props[CONF_SERVER_PORT], nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
