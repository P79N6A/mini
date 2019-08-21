package core

import "sync"

type Item interface{}

type SyncQueue struct {
	mutex  sync.Mutex
	item   []Item
	enalbe bool
}

func NewQueue() *SyncQueue {
	return &SyncQueue{
		item:   []Item{},
		enalbe: true,
	}
}

func (sq *SyncQueue) Enqueue(v Item) {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()
	if sq.enalbe {
		sq.item = append(sq.item, v)
	}

}

func (sq *SyncQueue) Dequeue() Item {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()
	if sq.enalbe {
		result := sq.item[0]
		sq.item = sq.item[1:]
		return result
	}
	return nil

}

func (sq *SyncQueue) Peek() Item {
	if sq.enalbe {
		return sq.item[len(sq.item)-1]
	}
	return nil
}

func (sq *SyncQueue) Size() int {
	if sq.enalbe {
		return len(sq.item)
	}
	return -1
}

func (sq *SyncQueue) Close() {
	sq.enalbe = false
}
