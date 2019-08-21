package main

import (
	"flag"
	"fmt"
	"goodcoder/pkg/conf"
	"goodcoder/pkg/core"
	"goodcoder/pkg/url"
	"os"
	"time"
	// "icode.baidu.com/baidu/go-lib/log"
	// "icode.baidu.com/baidu/go-lib/log/log4go"
)

const logFileName = "mini_spider.log"

var (
	help     bool
	version  bool
	confFile string
	logPath  string
)

func init() {
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&version, "version", false, "show version")
	flag.StringVar(&logPath, "log", "./log", "logpath")
	flag.StringVar(&confFile, "confFile", "./conf/spider.conf", "configfile")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}
	if version {
		fmt.Println(`mini_spider/1.10.0`)
	}

	// get config
	config, err := conf.GetConfigFromFile(confFile)
	if err != nil {
		return
	}
	// set log
	// log4go.SetLogFormat(log4go.FORMAT_DEFAULT_WITH_PID)
	// log.Init("mini_spider", "INFO", logPath, true, "midnight", 5)

	// log
	// util.LogUsefulMessage(config, confFile, logPath)

	// get url
	urls, _ := url.GetUrlFromFile(config.UrlListFile)
	if err != nil {
		// log.Logger.Error("GetUrlFromFile urls: %v", err)
	}

	// new mini_spider
	ms := core.NewMiniSpider(config, urls)
	if ms == nil {
		// log.Logger.Error("NewMiniSpider err get nil")
	}
	ms.Start()
	time.Sleep(2)

}

func usage() {
	fmt.Fprintf(os.Stderr, `mini_spider version: mini_spider/1.10.0
Usage: mini_spider [-help help] [-confFile confFile] [-logPath logPath]

Options:
`)
	flag.PrintDefaults()
}
