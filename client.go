package greq

import (
	"net/http"
)

func (c *Client) Send(rawReq *http.Request) (*http.Response, error) {
	//并发控制
	c.pool <- struct{}{}
	defer func() { <-c.pool }()
	if rawReq == nil {
		return nil, buildReqErr
	}
	resp, err := c.c.Do(rawReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) NewRequest(url string, header http.Header, handler ReaderHandler) *Request {
	if handler == nil {
		handler = defReader
	}
	return &Request{
		client: c,
		url:    url,
		header: header,
		reader: handler,
	}
}

func (c *Client) NewDefRequest(url string) *Request {
	return c.NewRequest(url, formHeader, nil)
}

func (c *Client) NewJsonRequest(url string) *Request {
	return c.NewRequest(url, jsonHeader, jsonReader)
}

func (c *Client) NewFormRequest(url string) *Request {
	return c.NewRequest(url, formHeader, formReader)
}
