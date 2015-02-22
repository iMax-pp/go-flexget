// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"github.com/go-martini/martini"
	utils "github.com/iMax-pp/go-utils"
	"github.com/martini-contrib/render"
	cache "github.com/robfig/go-cache"
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

	m := martini.Classic()
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	m.Get("/api/status", StatusHandler)
	m.Get("/api/logs", LogsHandler)
	m.Get("/api/config", ConfigHandler)

	// Up and listening
	logger.Info("Will start listening on port", props[CONF_SERVER_PORT])
	m.RunOnAddr(":" + props[CONF_SERVER_PORT])
}
