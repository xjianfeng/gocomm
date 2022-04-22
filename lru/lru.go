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
	Expired   int64
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

// 超时节点
func (l *LruQueue) AddExpiredNode(key string, data interface{}, sec int64) {
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
	now := time.Now().Unix()
	l.List[key] = &NodeLruData{
		Tail:      l.Head,
		Timestamp: now,
		Expired:   now + sec,
		Data:      data,
	}
	l.Head = key
	if l.Tail == "" {
		l.Tail = key
	}
	atomic.AddInt32(&l.Size, 1)
}

func (l *LruQueue) GetNode(key string) *NodeLruData {
	return l.refreshNode(key)
}

func (l *LruQueue) DelNode(key string) {
	l.Lock()
	defer l.Unlock()
	data, ok := l.List[key]
	if !ok {
		return
	}
	if key == l.Head {
		l.Head = data.Tail
	}
	if key == l.Tail || data.Tail == "" {
		l.Tail = data.Head
	}
	prevNode, ok := l.List[data.Head]
	if ok {
		prevNode.Tail = data.Tail
	}
	nextNode, ok := l.List[data.Tail]
	if ok {
		nextNode.Head = data.Head
	}
	delete(l.List, key)
	atomic.AddInt32(&l.Size, -1)
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

func (l *LruQueue) refreshNode(key string) *NodeLruData {
	l.Lock()
	var expired = false
	// 超时删除键值
	// 在defer删除防止死锁
	defer func() {
		l.Unlock()
		if !expired {
			return
		}
		l.DelNode(key)
	}()

	node, ok := l.List[key]
	if !ok {
		return nil
	}
	now := time.Now().Unix()
	// 超时删除
	if node.Expired > 0 && node.Expired < now {
		expired = true
		return nil
	}
	head, ok := l.List[l.Head]
	if !ok {
		return nil
	}
	//访问的元素在头部不用刷新
	if key == l.Head {
		return node
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
	return node
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
