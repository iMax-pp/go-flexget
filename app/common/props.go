// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package common

import (
	utils "github.com/iMax-pp/go-utils"
)

const (
	APP_PROPS_FILE = "application.properties"
)

var props = initProps()

// Init Application properties
func initProps() map[string]string {
	properties, _ := utils.LoadConfig(APP_PROPS_FILE)
	return properties
}

func Props() map[string]string {
	return props
}
