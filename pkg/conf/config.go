package conf

import (
	"fmt"

	gcfg "gopkg.in/gcfg.v1"
)

type Spider struct {
	Config Config `gcfg:"spider"`
}

type Config struct {
	UrlListFile     string `gcfg:"urlListFile"`
	OutputDirectory string `gcfg:"outputDirectory"`
	MaxDepth        int    `gcfg:"maxDepth"`
	CrawlInterval   int    `gcfg:"crawlInterval"`
	CrawlTimeout    int    `gcfg:"crawlTimeout"`
	TargetUrl       string `gcfg:"targetUrl"`
	SthreadCount    int    `gcfg:"sthreadCount"`
}

var spider Spider

func GetConfigFromFile(path string) (*Config, error) {
	err := gcfg.ReadFileInto(&spider, path)
	if err != nil {
	}
	err = CheckConfig(spider.Config)
	if err != nil {
		return nil, err
	}
	return &spider.Config, nil
}

func CheckConfig(config Config) error {
	if config.UrlListFile == "" {
		return fmt.Errorf("urlListFile is empty")
	}

	if config.OutputDirectory == "" {
		return fmt.Errorf("outputDirectory is empty")
	}

	if config.MaxDepth <= 0 {
		return fmt.Errorf("maxDepth's value is wrong")
	}

	if config.CrawlInterval <= 0 {
		return fmt.Errorf("crawlInterval's value is wrong")
	}

	if config.CrawlTimeout <= 0 {
		return fmt.Errorf("crawlTimeout's value is wrong")
	}

	if config.TargetUrl == "" {
		return fmt.Errorf("targetUrl's value is wrong")
	}

	if config.SthreadCount <= 0 {
		return fmt.Errorf("sthreadCount's value is wrong")
	}

	return nil
}
