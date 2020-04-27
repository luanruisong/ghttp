package greq

import (
	"io"
	"io/ioutil"
	"net/http"
)

func (r *Request) buildResp(res *http.Response) *Resp {
	defer res.Body.Close()
	myResp := &Resp{
		Ok:     res.StatusCode == http.StatusOK,
		Status: res.StatusCode,
		Header: res.Header,
	}
	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		myResp.Ok = false
		myResp.Err = err
	} else {
		myResp.RawBody = rawBody
		body := string(rawBody)
		myResp.Body = body
	}
	return myResp
}

func (r *Request) buildReq(method string, param io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, r.url, param)
	if err != nil {
		return nil, err
	}
	req.Header = r.header
	return req, nil
}

func (r *Request) do(method string, param interface{}) *Resp {
	reader := r.reader(param)
	rawReq, err := r.buildReq(method, reader)
	if err != nil {
		return &Resp{
			Ok:  false,
			Err: err,
		}
	}
	resp, err := r.client.Send(rawReq)
	if err != nil {
		return &Resp{
			Ok:  false,
			Err: err,
		}
	}
	return r.buildResp(resp)
}

func (r *Request) Get(p interface{}) *Resp {
	return r.do(http.MethodGet, p)
}

func (r *Request) Post(p interface{}) *Resp {
	return r.do(http.MethodPost, p)
}

func (r *Request) Put(p interface{}) *Resp {
	return r.do(http.MethodPut, p)
}

func (r *Request) Delete(p interface{}) *Resp {
	return r.do(http.MethodDelete, p)
}

func (r *Request) Options(p interface{}) *Resp {
	return r.do(http.MethodOptions, p)
}
