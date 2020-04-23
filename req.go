package greq

import (
	"io/ioutil"
	"net/http"
	"strings"
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
		myResp.Body = string(myResp.RawBody)
	}
	return myResp
}

func (r *Request) buildReq(method string, param *strings.Reader) (*http.Request,error) {
	req, err := http.NewRequest(r.url, method, param)
	if err != nil {
		return nil,err
	}
	req.Header = r.header
	return req,nil
}

func (r *Request) do(method string, param *strings.Reader) *Resp {
	rawReq,err := r.buildReq(method, param)
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
	return r.do(http.MethodGet, r.reader(p))
}

func (r *Request) Post(p interface{}) *Resp {
	return r.do(http.MethodPost, r.reader(p))
}

func (r *Request) Put(p interface{}) *Resp {
	return r.do(http.MethodPut, r.reader(p))
}

func (r *Request) Delete(p interface{}) *Resp {
	return r.do(http.MethodDelete, r.reader(p))
}

func (r *Request) Options(p interface{}) *Resp {
	return r.do(http.MethodOptions, r.reader(p))
}
