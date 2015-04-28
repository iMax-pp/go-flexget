// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package common

import (
	cache "github.com/robfig/go-cache"
	"strconv"
	"time"
)

const (
	confCacheExpir = "cache.expiration"
	confCacheClean = "cache.cleanup"
)

var fgCache *cache.Cache

// Init FlexGet Cache
func init() {
	expir, _ := strconv.Atoi(Props()[confCacheExpir])
	clean, _ := strconv.Atoi(Props()[confCacheClean])
	fgCache = cache.New(time.Duration(expir)*time.Second, time.Duration(clean)*time.Second)
}

// Cache returns the cache object
func Cache() *cache.Cache {
	return fgCache
}
