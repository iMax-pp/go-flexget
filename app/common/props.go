// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package common

import (
	utils "github.com/iMax-pp/go-utils"
)

const (
	APP_PROPS_FILE = "application.properties"
)

var props map[string]string

// Init Application properties
func init() {
	props, _ = utils.LoadConfig(APP_PROPS_FILE)
}

func Props() map[string]string {
	return props
}
