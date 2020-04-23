package ghttp

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defPoolSize = 100
)

var (
	defClient   *Client
	buildReqErr = errors.New("build req error")
	jsonHeader  = http.Header{
		"Content-Type": []string{"application/json;charset=utf-8"},
	}

	formHeader = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded;charset=utf-8"},
	}
)

func jsonReader(i interface{}) *strings.Reader {
	str, err := jsoniter.MarshalToString(i)
	if err != nil {
		return nil
	}
	return strings.NewReader(str)
}

func formReader(i interface{}) *strings.Reader {
	str, err := jsoniter.MarshalToString(i)
	if err != nil {
		return nil
	}
	return strings.NewReader(str)
}

func defReader(i interface{}) *strings.Reader {
	var str string
	switch i.(type) {
	case string:
		str = i.(string)
	default:
		str = fmt.Sprintf("%v", i)
	}
	return strings.NewReader(str)
}

func getDefaultClient() *Client {
	if defClient == nil {
		defClient = NewClient(defPoolSize)
	}
	return defClient
}

func NewClient(poolSize int) *Client {
	return &Client{
		c: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment, //代理使用
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second, //连接超时时间
					KeepAlive: 30 * time.Second, //连接保持超时时间
				}).DialContext,
				MaxIdleConns:          poolSize,         //client对与所有host最大空闲连接数总和
				IdleConnTimeout:       90 * time.Second, //空闲连接在连接池中的超时时间
				TLSHandshakeTimeout:   10 * time.Second, //TLS安全连接握手超时时间
				ExpectContinueTimeout: 1 * time.Second,  //发送完请求到接收到响应头的超时时间
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		pool: make(chan struct{}, poolSize),
	}
}

func New(url string) *Request {
	return getDefaultClient().NewDefRequest(url)
}

func NewJson(url string) *Request {
	return getDefaultClient().NewJsonRequest(url)
}

func NewForm(url string) *Request {
	return getDefaultClient().NewFormRequest(url)
}
