// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package main

import (
	"github.com/martini-contrib/render"
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
func StatusHandler(r render.Render) {
	status, err := getStatus()
	if err != nil {
		r.Error(http.StatusInternalServerError)
		return
	}
	version, err := getVersion()
	if err != nil {
		r.Error(http.StatusInternalServerError)
		return
	}
	r.JSON(http.StatusOK, Status{status, version})
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
