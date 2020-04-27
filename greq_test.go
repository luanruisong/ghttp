package greq

import (
	"io/ioutil"
	"testing"
)

func getParam() interface{} {
	return struct {
		Name string `form:"n" json:"name"`
		Key  string `form:"k" json:"key"`
	}{
		Name: "name",
		Key:  "key",
	}
}

func TestJsonReader(t *testing.T){
	b , err := ioutil.ReadAll(jsonReader(getParam()))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("read",len(b),"content",string(b))
	}
}

func TestFormReader(t *testing.T){
	b , err := ioutil.ReadAll(formReader(getParam()))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("read",len(b),"content",string(b))
	}
}


func TestDefReader(t *testing.T){
	b , err := ioutil.ReadAll(defReader(getParam()))
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("read",len(b),"content",string(b))
	}
}

func TestSendJsonReq(t *testing.T) {
	p := getParam()
	resp := NewJson("https://baidu.com").Get(p)
	if resp.Ok {
		t.Log("send to baidu success")
	} else {
		t.Error(resp.Err.Error())
	}
}

func TestFormJsonReq(t *testing.T) {
	p := getParam()
	resp := NewForm("https://baidu.com").Get(p)
	if resp.Ok {
		t.Log("send to baidu success")
	} else {
		t.Error(resp.Err.Error())
	}
}

func TestErrRes(t *testing.T){
	p := getParam()
	resp := NewForm("https://.com").Get(p)
	if resp.Ok {
		t.Log("send to baidu success")
	} else {
		t.Error(resp.Err.Error())
	}
}

