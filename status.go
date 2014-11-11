// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
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
	logger.TraceBegin("StatusHandler")

	status, err := getStatus()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	version, err := getVersion()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	body, err := json.Marshal(Status{status, version})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprint(w, string(body))
	logger.TraceEnd("StatusHandler")
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
