package greq

import (
	"bytes"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func getDefaultClient() *Client {
	if defClient == nil {
		defClient = NewClient(defPoolSize)
	}
	return defClient
}

func jsonReader(i interface{}) io.Reader {
	str, err := jsoniter.Marshal(i)
	if err != nil {
		return nil
	}
	return bytes.NewReader(str)
}

func formReader(i interface{}) io.Reader {
	if i == nil {
		return nil
	}
	values := url.Values{}
	v := reflect.ValueOf(i)
	t := v.Type()

	switch t.Kind(){
	case reflect.Ptr:
		v = v.Elem()
		fallthrough
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			k := fmt.Sprintf("%v",iter.Key().Interface())
			iv := iter.Value()
			kind := iv.Kind()
			if kind == reflect.Slice || kind == reflect.Array  {
				for idx := 0 ;idx < iv.Len() ; idx ++ {
					currValue := iv.Index(idx).Interface()
					values.Add(k,fmt.Sprintf("%v", currValue))
				}
			} else {
				values.Add(
					k,
					fmt.Sprintf("%v", iter.Value().Interface()),
				)
			}
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			vf := v.Field(i)
			name := f.Tag.Get("form")
			if len(name) == 0 {
				name = f.Name
			}
			values.Add(name, fmt.Sprintf("%v", vf.Interface()))
		}
	}
	return strings.NewReader(values.Encode())
}

func defReader(i interface{}) io.Reader {
	var str string
	switch i.(type) {
	case string:
		str = i.(string)
	default:
		str = fmt.Sprintf("%v", i)
	}
	return strings.NewReader(str)
}

