package lru

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestLru(t *testing.T) {
	return
	l, _ := InitLruCap(50000)
	g := sync.WaitGroup{}
	g.Add(2)
	go func() {
		for i := 1; i < 100000; i++ {
			l.AddNode(fmt.Sprint(i), i)
		}
		g.Done()
	}()
	go func() {
		for i := 1; i < 30000; i++ {
			l.GetNode(strconv.Itoa(i))
		}

		for i := 1; i < 30000; i++ {
			if i%3 == 0 {
				l.GetNode(strconv.Itoa(i))
			} else {
				l.GetNode(l.Head)
			}
		}
		t.Logf("Tail %s", l.Tail)
		t.Logf("Head %s", l.Head)
		g.Done()
	}()
	g.Wait()
	t.Logf("lSize %d", l.Size)
}

func TestPrint(t *testing.T) {
	l, _ := InitLruCap(10)
	for i := 1; i < 10; i++ {
		l.AddNode(fmt.Sprint(i), i)
	}
	l.printList()
	l.GetNode("4")
	println("GetNode 4")
	l.printList()

	l.DelNode("3")
	println("DelNone 3")
	l.printList()
	l.GetNode("2")
	println("GetNode 2")
	l.printList()
}

func TestOverFlow(t *testing.T) {
	l, _ := InitLruCap(5)
	for i := 1; i < 10; i++ {
		l.AddNode(fmt.Sprint(i), i)
	}
	l.printList()
}

func TestSetNode(t *testing.T) {
	l, _ := InitLruCap(5)
	l.AddNode("1", 333)
	l.printList()
	node := l.GetNode("1")
	node.Data = 555
	l.printList()
}
