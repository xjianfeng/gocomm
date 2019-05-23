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

	ret, err := lhttp.HttpGet("http://www.abc.com")
	if err != nil {
		logger.LogError(err.Error())
	}
	logger.LogInfo("GetData %s,", ret)

	data := sort.SortMap{
		"abc":  1,
		"sedw": 3,
		"kkkk": 5,
		"vvvv": 6,
	}
	v := data.SortMapStrKey(false)
	logger.LogInfo("sortMap %v,", v)

	//注册三个定时器
	timer.CallOut("unitKey1", time.Second, func() { println("1111111111111111111") })
	timer.CallOut("unitKey2", 3*time.Second, func() { println("222222222222222222") })
	timer.CallOut("unitKey3", 5*time.Second, func() { println("33333333333333333") })
	//停掉一个
	timer.Stop("unitKey3")

	time.Sleep(10 * time.Second)
}
