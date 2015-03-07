// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package app

import (
	"github.com/golang/glog"
	. "github.com/iMax-pp/go-flexget/app/common"
	. "github.com/iMax-pp/go-flexget/app/services"
	"net/http"
)

const (
	CONF_SERVER_PORT = "server.port"
)

func Server() {
	// Serve static content
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Serve dynamic content
	http.Handle("/api/status", http.HandlerFunc(StatusHandler))
	http.Handle("/api/logs", http.HandlerFunc(LogsHandler))
	http.Handle("/api/config", http.HandlerFunc(ConfigHandler))

	// Up and listening
	glog.Info("Will start listening on port ", Props()[CONF_SERVER_PORT])
	if err := http.ListenAndServe(":"+Props()[CONF_SERVER_PORT], nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
