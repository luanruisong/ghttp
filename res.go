package ghttp

import jsoniter "github.com/json-iterator/go"

func (res *Resp) UnmarshalJson(p interface{}) error {
	return jsoniter.Unmarshal(res.RawBody, p)
}
