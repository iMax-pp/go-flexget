// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package app

import (
	"github.com/golang/glog"
	common "github.com/iMax-pp/go-flexget/app/common"
	services "github.com/iMax-pp/go-flexget/app/services"
	"net/http"
)

const (
	confServerPort = "server.port"
)

// Server main function
func Server() {
	// Serve static content
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Serve dynamic content
	http.Handle("/api/status", http.HandlerFunc(services.StatusHandler))
	http.Handle("/api/logs", http.HandlerFunc(services.LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(services.ConfigHandler))
	http.Handle("/api/flexget/start", http.HandlerFunc(services.StartFlexGetHandler))
	http.Handle("/api/flexget/stop", http.HandlerFunc(services.StopFlexGetHandler))
	http.Handle("/api/flexget/reload", http.HandlerFunc(services.ReloadFlexGetHandler))

	// Up and listening
	glog.Info("Will start listening on port ", common.Props()[confServerPort])
	if err := http.ListenAndServe(":"+common.Props()[confServerPort], nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
