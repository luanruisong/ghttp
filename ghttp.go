package greq

import (
	"net"
	"net/http"
	"time"
)

const (
	defPoolSize = 100
)

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
