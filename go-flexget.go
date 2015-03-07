package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/iMax-pp/go-flexget/app"
	"os"
)

func init() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go-flexget -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	glog.Info("Starting go-flexget...")
	app.Server()
}
