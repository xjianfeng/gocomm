package timer

import (
	log "github.com/xjianfeng/gocomm/logger"
	"runtime"
	"sync"
	"time"
)

var timeManager = new(sync.Map)

//保证key的唯一性
func CallOut(key interface{}, tick time.Duration, fuc func()) {
	Stop(key)
	f := func() {
		callFunc(key, fuc)
	}
	t := time.AfterFunc(tick, f)
	timeManager.Store(key, t)
}

//回收timeManager
func callFunc(key interface{}, fun func()) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 2048)
			l := runtime.Stack(buf, false)
			log.LogError("%v:%s", r, buf[:l])
			erroring.SendErrorResponse()
		}
	}()

	v := timeManager.Load(key)
	if v != nil {
		timeManager.Delete(key)
	}
	fun()
}

func Stop(key interface{}) {
	t := timeManager.Load(key)
	if t == nil {
		return
	}
	v := t.(*time.Timer)
	if v != nil {
		v.Stop()
		timeManager.Delete(key)
	}
}
