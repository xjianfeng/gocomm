package lhttp

import (
	"bytes"
	"crypto/tls"
	"github.com/xjianfeng/gocomm/logger"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

var (
	tlsCaKey      = "cakey/apiclient_key.pem"
	tlsCacert     = "cakey/apiclient_cert.pem"
	defaultHeader = "application/x-www-form-urlencoded"
)

var log = logger.New("http_client.log")

type HttpCustom interface {
	UrlEncode() string
}

type UrlParam struct {
	Data map[string]string
}

func (u *UrlParam) UrlEncode() string {
	data := u.Data
	if data == nil {
		return ""
	}

	buff := bytes.Buffer{}
	for k, v := range data {
		buff.WriteString(k)
		buff.WriteString("=")
		buff.WriteString(v)
		buff.WriteString("&")
	}
	buf := buff.Bytes()
	end := len(buf) - 1
	return string(buf[0:end])
}

func (u *UrlParam) SortUrlEncode() string {
	data := u.Data
	if data == nil {
		return ""
	}
	keysList := []string{}
	for k, _ := range data {
		keysList = append(keysList, k)
	}
	sort.Strings(keysList)

	buff := bytes.Buffer{}
	for _, k := range keysList {
		v := data[k]
		buff.WriteString(k)
		buff.WriteString("=")
		buff.WriteString(v)
		buff.WriteString("&")
	}
	buf := buff.Bytes()
	end := len(buf) - 1
	return string(buf[0:end])
}

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.LogError("HttpGet url %s error %s", url, err.Error())
		return []byte{}, err
	}
	resp.Close = true

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return result, nil
}

func HttpPost(url string, postBody []byte, header map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(postBody))
	if err != nil {
		log.LogError("HttpPost Error %s", err.Error())
		return []byte{}, err
	}
	req.Close = true

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", defaultHeader)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.LogError("HttpPost Client Do Error %s", err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.LogError("HttpPost Error %s", err.Error())
		return []byte{}, err
	}
	if len(body) > 1024 {
		log.LogInfo("HttpPost Success Url %s header:%v Data %s, ret %s", url, header, postBody, body[:1024])
	} else {
		log.LogInfo("HttpPost Success Url %s header:%v Data %s, ret %s", url, header, postBody, body)
	}
	return body, nil
}

func LoadCaConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(tlsCacert, tlsCaKey)
	if err != nil {
		log.LogError("LoadCaKey", err)
		return nil, err
	}

	config := &tls.Config{
		//这里先不验证服务端证书，是自己签发的呀
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
	return config, nil
}

func TlsPost(url string, postBody string, header map[string]string) ([]byte, error) {
	tlsConf, err := LoadCaConfig()
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, strings.NewReader(postBody))
	req.Close = true
	if err != nil {
		log.LogError("HttpPost Error %s", err.Error())
		return []byte{}, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", defaultHeader)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.LogError("HttpPost Error %s", err.Error())
		return []byte{}, err
	}
	log.LogInfo("HttpPost Success Data %s, ret %s", postBody, body)
	return body, nil
}

//设置双向认证请求证书
func SetTLSKey(caKey, caCert string) {
	tlsCaKey = caKey
	tlsCacert = caCert
}
