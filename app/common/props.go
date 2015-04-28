// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package common

import (
	utils "github.com/iMax-pp/go-utils"
)

const (
	appPropsFile = "application.properties"
)

var props = initProps()

// Init Application properties
func initProps() map[string]string {
	properties, _ := utils.LoadConfig(appPropsFile)
	return properties
}

// Props returns the application properties
func Props() map[string]string {
	return props
}
