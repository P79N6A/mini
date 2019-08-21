package url

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
	// "icode.baidu.com/baidu/go-lib/log"
)

type UrlsMap struct {
	urlMap map[string]bool
	mutex  sync.Mutex
}

func NewUrlsMap() *UrlsMap {
	return nil
}

func (*UrlsMap) addUrl() {

}

func (*UrlsMap) deleteUrl() {

}

func GetUrlFromFile(path string) ([]string, error) {
	result := []string{}
	out, err := ioutil.ReadFile(path)
	if err != nil {
		// log.Logger.Error("ReadFile error, err: %s", err)
		return nil, err
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		// log.Logger.Error("Unmarshal error, err: %s", err)
		return nil, err
	}
	return result, nil
}

func GetUrlsFromRes(body io.Reader, regexp *regexp.Regexp, url string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		// log.Logger.Error("goquery NewDocumentFromReader error")
		return nil, err
	}
	var urlLinks = make([]string, 0)
	doc.Find("a").Each(func(_ int, link *goquery.Selection) {
		href, ok := link.Attr("href")
		if ok {
			if regexp.MatchString(href) {
				href = GetResolveReference(href, url)
				urlLinks = append(urlLinks, href)
			}
		}
	})
	return urlLinks, nil
}

func GetResolveReference(relUrl string, baseUrl string) string {
	rel, err := url.Parse(relUrl)
	if err != nil {
		// log.Logger.Error(fmt.Sprintf("resolveReference url[%s]: %s", relUrl, err))
		return ""
	}
	base, err := url.Parse(baseUrl)
	if err != nil {
		// log.Logger.Error(fmt.Sprintf("resolveReference url[%s]: %s", baseUrl, err))
		return ""
	}
	return base.ResolveReference(rel).String()
}
