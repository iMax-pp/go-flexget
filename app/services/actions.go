// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package services

import (
	"fmt"
	"github.com/golang/glog"
	common "github.com/iMax-pp/go-flexget/app/common"
	"net/http"
	"strings"
)

var (
	startFlexGetCmd  = "/opt/local/bin/flexget daemon start"
	stopFlexGetCmd   = "/opt/local/bin/flexget daemon stop"
	reloadFlexGetCmd = "/opt/local/bin/flexget daemon reload"
)

// StartFlexGetHandler '/api/flexget/start' request handler.
func StartFlexGetHandler(w http.ResponseWriter, req *http.Request) {
	glog.Info("Start FlexGet")
	if result, err := execFlexGetAction(startFlexGetCmd); err != nil {
		glog.Error("Error starting FlexGet: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(result))
	}
}

// StopFlexGetHandler '/api/flexget/stop' request handler.
func StopFlexGetHandler(w http.ResponseWriter, req *http.Request) {
	glog.Info("Stop FlexGet")
	if result, err := execFlexGetAction(stopFlexGetCmd); err != nil {
		glog.Error("Error stopping FlexGet: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(result))
	}
}

// ReloadFlexGetHandler '/api/flexget/reload' request handler.
func ReloadFlexGetHandler(w http.ResponseWriter, req *http.Request) {
	glog.Info("Reload FlexGet")
	if result, err := execFlexGetAction(reloadFlexGetCmd); err != nil {
		glog.Error("Error reloading FlexGet: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(result))
	}
}

func execFlexGetAction(cmd string) (string, error) {
	body, err := common.ExecSSHCmd(cmd)
	if err != nil && !strings.Contains(err.Error(), "Process exited with: 1") {
		return "", err
	}
	return body, nil
}
