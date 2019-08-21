package conf

import "testing"

func TestGetConfigFromFile(t *testing.T) {
	path := "../../conf/spider.conf"
	config, err := GetConfigFromFile(path)
	if err != nil {
		t.Errorf("getConfig err, err: %s", err)
	}
	if config.UrlListFile != "./data/url.data" || config.OutputDirectory != "./output" ||
		config.MaxDepth != 1 || config.CrawlInterval != 1 ||
		config.CrawlTimeout != 1 || config.TargetUrl != ".*.(htm|html)$" ||
		config.SthreadCount != 8 {
		t.Errorf("getConfig err, get config: %v", spider)
	}
}
