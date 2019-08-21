package util

import (
	"io/ioutil"
	"net/url"
	"path"

	"github.com/axgle/mahonia"
)

func SaveToDisk(body []byte, urlLink string, filePath string) error {

	src := string(body)
	enc := mahonia.NewEncoder("UTF-8")
	src = enc.ConvertString(src)

	filepath := path.Join(filePath, url.QueryEscape(urlLink))

	err := ioutil.WriteFile(filepath, []byte(src), 0644)
	if err != nil {
		// log.Logger.Error("WriteFile error: %s", err)
		return err
	}
	return nil
}

// func LogUsefulMessage(config *conf.Config, confFile string, logPath string) {
// 	log.Logger.Info("mini_spider starts!")
// 	log.Logger.Info("mini_spider version: 1.10.0")
// 	log.Logger.Info("mini_spider confFile: %s", confFile)
// 	log.Logger.Info("mini_spider logPath: %s", logPath)
// 	log.Logger.Info("mini_spider UrlListFile: %s", config.UrlListFile)
// 	log.Logger.Info("mini_spider OutputDirectory: %s", config.OutputDirectory)
// 	log.Logger.Info("mini_spider MaxDepth: %d", config.MaxDepth)
// 	log.Logger.Info("mini_spider CrawlTimeout: %d", config.CrawlTimeout)
// 	log.Logger.Info("mini_spider TargetUrl: %s", config.TargetUrl)
// 	log.Logger.Info("mini_spider SthreadCount: %d", config.SthreadCount)
// }
