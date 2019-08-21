package core

import (
	"bytes"
	"fmt"
	"goodcoder/pkg/conf"
	"goodcoder/pkg/url"
	"goodcoder/pkg/util"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type Worker struct {
	syncQueue *SyncQueue
	//urlsMap   *url.UrlsMap
	config *conf.Config
	rege   *regexp.Regexp
}

func NewWorker(syncQueue *SyncQueue, config *conf.Config, regexp *regexp.Regexp) *Worker {
	return &Worker{
		syncQueue: syncQueue,
		config:    config,
		rege:      regexp,
	}
}

func (w *Worker) Start() {
	if w.syncQueue.Size() == 0 {
		// log.Logger.Info("all jobs are done!")
		time.Sleep(time.Second * 1)
		fmt.Printf("stop")
		stopCh <- stop{}
		return
	}
	urlTask := w.syncQueue.Dequeue()

	timeout := time.Duration(w.config.CrawlTimeout)

	c := &http.Client{
		Timeout: 2 * timeout * time.Second,
	}

	resp, err := c.Get(urlTask.(UrlTask).url)
	if err != nil {
		// log.Logger.Error("http get error: %s, url: %s", err, urlTask.(UrlTask).url)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close() //  must close
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	defer resp.Body.Close()

	if urlTask.(UrlTask).maxDepth < w.config.MaxDepth {
		urls, err := url.GetUrlsFromRes(resp.Body, w.rege, urlTask.(UrlTask).url)
		// log.Logger.Info("get new urltaskettse: %d w.rege: %v", len(urls), w.rege)
		if err != nil {
			// log.Logger.Error("GetUrlsFromRes error, %s", err)
		}
		for _, url := range urls {
			urlTask := UrlTask{
				url:      url,
				maxDepth: urlTask.(UrlTask).maxDepth + 1,
			}
			w.syncQueue.Enqueue(urlTask)
		}
	}

	// save data
	if w.rege.MatchString(urlTask.(UrlTask).url) {
		// 转码
		err = util.SaveToDisk(bodyBytes, urlTask.(UrlTask).url, w.config.OutputDirectory)
		if err != nil {
			// log.Logger.Error("SaveToDisk error, %s", err)
			return
		}
	}

}
