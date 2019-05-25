package lru

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type NodeLruData struct {
	Head      string
	Tail      string
	Timestamp int64
	Data      interface{}
}

type LruQueue struct {
	sync.Mutex
	CapSize int32
	Size    int32
	Head    string
	Tail    string
	List    map[string]*NodeLruData
}

func (l *LruQueue) AddNode(key string, data interface{}) {
	if atomic.LoadInt32(&l.Size) >= l.CapSize {
		l.delTailNode(l.Tail)
	}
	l.Lock()
	defer l.Unlock()
	if _, ok := l.List[key]; ok {
		return
	}
	oldHead, ok := l.List[l.Head]
	if ok {
		oldHead.Head = key
	}
	l.List[key] = &NodeLruData{
		Tail:      l.Head,
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
	l.Head = key
	if l.Tail == "" {
		l.Tail = key
	}
	atomic.AddInt32(&l.Size, 1)
}

func (l *LruQueue) GetNode(key string) interface{} {
	return l.refreshNode(key)
}

func (l *LruQueue) delTailNode(key string) {
	l.Lock()
	defer l.Unlock()
	data, ok := l.List[key]
	if !ok {
		return
	}
	l.Tail = data.Head
	delete(l.List, key)
	atomic.AddInt32(&l.Size, -1)
}

func (l *LruQueue) printList() {
	l.Lock()
	defer l.Unlock()
	head, ok := l.List[l.Head]
	if !ok {
		return
	}
	fmt.Printf("%v, ", head.Data)
	tail := head.Tail
	for {
		next, ok := l.List[tail]
		if !ok {
			break
		}
		fmt.Printf("%v, ", next.Data)
		tail = next.Tail
	}
	fmt.Printf("\n")
}

func (l *LruQueue) refreshNode(key string) interface{} {
	l.Lock()
	defer l.Unlock()

	node, ok := l.List[key]
	if !ok {
		return nil
	}
	head, ok := l.List[l.Head]
	if !ok {
		return nil
	}
	//访问的元素在头部不用刷新
	if key == l.Head {
		return node.Data
	}
	if node.Tail == "" || node.Tail == l.Tail {
		prevNode, ok := l.List[node.Head]
		if ok {
			prevNode.Tail = node.Tail
		}
		l.Tail = node.Head
	} else {
		nextNode, ok := l.List[node.Tail]
		if ok {
			nextNode.Head = node.Head
		}
		prevNode, ok := l.List[node.Head]
		if ok {
			prevNode.Tail = node.Tail
		}
	}
	node.Head = ""
	node.Tail = l.Head

	head.Head = key
	l.Head = key
	return node.Data
}

func InitLruCap(capSize int32) (*LruQueue, error) {
	if capSize <= 1 {
		return nil, errors.New("capSize error")
	}
	lru := new(LruQueue)
	lru.CapSize = capSize
	lru.List = make(map[string]*NodeLruData)
	return lru, nil
}
