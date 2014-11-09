// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
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
