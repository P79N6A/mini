package core

import (
	"fmt"
	"goodcoder/pkg/conf"
	"regexp"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

var stopCh = make(chan struct{})

type stop struct {
}

type UrlTask struct {
	url      string
	maxDepth int
}

type MiniSpider struct {
	syncQueue *SyncQueue
	workers   []*Worker
	//urlsMap   *url.UrlsMap
	config *conf.Config
}

func NewMiniSpider(conf *conf.Config, urls []string) *MiniSpider {
	// log.Logger.Info("NewMiniSpider starts!")
	ms := &MiniSpider{
		syncQueue: NewQueue(),
		workers:   []*Worker{},
		//urlsMap:   url.NewUrlsMap(),
		config: conf,
	}

	// url in queue
	for _, url := range urls {
		urlTask := UrlTask{
			url:      url,
			maxDepth: 0,
		}
		ms.syncQueue.Enqueue(urlTask)
	}

	urlre, err := regexp.Compile(conf.TargetUrl)
	if err != nil {
		// log.Logger.Error(fmt.Sprintf("failed to compile regexp: %s", err))
		return nil
	}

	for i := 0; i < conf.SthreadCount; i++ {
		worker := NewWorker(ms.syncQueue, ms.config, urlre)
		ms.workers = append(ms.workers, worker)
	}
	// log.Logger.Info("NewMiniSpider ends! get syncQueue: %v", ms.syncQueue)

	return ms
}

func (ms *MiniSpider) Start() {
	for i := 0; i < ms.config.SthreadCount; i++ {
		fmt.Printf("----%d----\n", i)
		go wait.Until(ms.workers[i].Start, 1*time.Second, stopCh)
	}
	<-stopCh
}
