package lru

import (
	"fmt"
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
			l.GetNode(fmt.Sprint("%d", i))
		}

		for i := 1; i < 30000; i++ {
			if i%3 == 0 {
				l.GetNode(fmt.Sprint("%d", i))
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
	l.GetNode("1")
	println("GetNode 1")
	l.printList()
}

func TestOverFlow(t *testing.T) {
	l, _ := InitLruCap(5)
	for i := 1; i < 10; i++ {
		l.AddNode(fmt.Sprint(i), i)
	}
	l.printList()
}
