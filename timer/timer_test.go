package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	CallOut("abc", 5*time.Second, func() {
		println("111111111111111111111")
	})
	time.Sleep(10 * time.Second)
}
