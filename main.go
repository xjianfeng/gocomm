package main

import (
	"github.com/xjianfeng/gocomm/decry"
	"github.com/xjianfeng/gocomm/lhttp"
	"github.com/xjianfeng/gocomm/logger"
	"github.com/xjianfeng/gocomm/sort"
	"github.com/xjianfeng/gocomm/timer"
	"time"
)

func main() {
	logger.LogInfo("Md5Sum: %s", decry.Md5Sum([]byte("1234567890")))
	ret, err := lhttp.HttpGet("http://www.baidu.com")
	if err != nil {
		logger.LogError(err.Error())
	}
	Logger.LogInfo("GetData %s,", ret)

	data := sort.SortMap{
		"abc":  1,
		"sedw": 3,
		"kkkk": 5,
		"vvvv": 6,
	}
	v := data.SortMapStrKey(false)
	Logger.LogInfo("sortMap %v,", v)

	timer.CallOut("unitKey1", time.Second, func() { println("1111111111111111111") })
	timer.CallOut("unitKey2", 3*time.Second, func() { println("1111111111111111111") })
	timer.CallOut("unitKey3", 5*time.Second, func() { println("1111111111111111111") })

	time.Sleep(3 * time.Second)
	timer.Stop("unitKey3")
	time.Sleep(10 * time.Second)
}
