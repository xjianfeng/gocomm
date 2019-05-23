package lhttp

import (
	"fmt"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	var reqParam UrlParam
	reqParam.Data = make(map[string]string)
	reqParam.Data["uid"] = "1234456"
	reqParam.Data["token"] = "sdfsafasdfe"
	urlStr := reqParam.UrlEncode()
	fmt.Println("urlStr. ", urlStr)
}

func TestHttpPost(t *testing.T) {
	url := "http://www.sdfsdf.cm"
	body := "a=1"
	HttpPost(url, body, nil)
}
