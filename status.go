// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"strings"
)

type Status struct {
	Status  bool
	Version string
}

// Command to retrieve FlexGet status
var getStatusCmd = "ps | grep flexget | grep -v grep"

// Command to retrieve FlexGet version
var getVersionCmd = "cat /opt/local/bin/flexget | grep __requires__ | sed 's/__requires__ = .FlexGet==\\(.*\\)./\\1/'"

// '/api/status' request handler.
func StatusHandler(w http.ResponseWriter, req *http.Request) {
	glog.Info("Retrieve FlexGet status from server")
	status, err := getStatus()
	if err != nil {
		glog.Error("Error retrieving FlexGet status: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	version, err := getVersion()
	if err != nil {
		glog.Error("Error retrieving FlexGet version: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if body, err := json.Marshal(Status{status, version}); err != nil {
		glog.Error("Error preparing server response: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(body))
	}
}

func getStatus() (bool, error) {
	body, err := ExecSSHCmd(getStatusCmd)
	if err != nil && !strings.Contains(err.Error(), "Process exited with: 1") {
		return false, err
	}
	return body != "", nil
}

func getVersion() (string, error) {
	body, err := ExecSSHCmd(getVersionCmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(body), nil
}
