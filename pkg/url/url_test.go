package url

import (
	"net/http"
	"regexp"
	"testing"
	"time"

	"icode.baidu.com/baidu/go-lib/log"
)

func TestGetUrlFromFile(t *testing.T) {
	path := "../../data/url.data"
	result, err := GetUrlFromFile(path)
	if err != nil {
		t.Errorf("getUrl err, err: %s", err)
	}
	if result[0] != "http://www.baidu.com" ||
		result[1] != "http://www.sina.com.cn" {
		t.Errorf("getUrl err, get url: %v", result)
	}
}

func TestGetResolveReference(t *testing.T) {
	relUrl := "./sdasd"
	baseUrl := "http://www.sina.com.cn"
	result := GetResolveReference(relUrl, baseUrl)
	if result != "http://www.sina.com.cn/sdasd" {
		t.Errorf("getResolveReference err, get result: %s want: ", result)
	}
}
func TestGetUrlsFromRes(t *testing.T) {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Get("http://www.baidu.com")
	if err != nil {
		log.Logger.Error("http get error: %s", err)
		return
	}

	regexp, _ := regexp.Compile(".*.(htm|html)$")

	url := "http://www.baidu.com"

	result, err := GetUrlsFromRes(resp, regexp, url)
	if err != nil {
		t.Errorf("GetUrlsFromRes err, err: %s", err)
	}
	if result[0] != "http://www.baidu.com/gaoji/preferences.html" ||
		result[1] != "http://www.baidu.com/cache/sethelp/help.html" {
		t.Errorf("GetUrlsFromRes err, get: %v", result)
	}
	defer resp.Body.Close()
}
