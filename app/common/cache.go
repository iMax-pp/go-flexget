// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.
package common

import (
	cache "github.com/robfig/go-cache"
	"strconv"
	"time"
)

const (
	CONF_CACHE_EXPIR = "cache.expiration"
	CONF_CACHE_CLEAN = "cache.cleanup"
)

var fgCache *cache.Cache

// Init FlexGet Cache
func init() {
	expir, _ := strconv.Atoi(Props()[CONF_CACHE_EXPIR])
	clean, _ := strconv.Atoi(Props()[CONF_CACHE_CLEAN])
	fgCache = cache.New(time.Duration(expir)*time.Second, time.Duration(clean)*time.Second)
}

func Cache() *cache.Cache {
	return fgCache
}
