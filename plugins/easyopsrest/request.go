package easyopsrest

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/urlpb"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var jsonpbMarshaler = jsonpb.Marshaler{OrigName: true}

type httpRule struct {
	method string
	path string
	body string
}

func getHTTPRule(contract giraffe.Contract) (*httpRule, error) {
	rule := &httpRule{
		method: "GET",
		path:   fmt.Sprintf("/%s/%s", contract.Desc().ServiceName, contract.Desc().MethodName),
	}
	if pattern, ok := contract.Desc().MetaData["url_pattern"]; ok {
		match := regexp.MustCompile(`^(\w+)\s+((((/:)|(/))[\w]+)+|/)$`).FindStringSubmatch(pattern.(string))
		if match == nil {
			return nil, fmt.Errorf("not valid url format: %s", pattern)
		}
		rule.method = strings.ToUpper(match[1])
		rule.path = match[2]
	}
	if body, ok := contract.Desc().MetaData["data_field"]; ok {
		rule.body = body.(string)
	}
	return rule, nil
}

func isProtoMessage(v interface{}) (proto.Message, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Struct {
		nv := reflect.New(rv.Type())
		nv.Elem().Set(rv)
		v = nv.Interface()
	}
	if pb, ok := v.(proto.Message); ok {
		return pb, ok
	}
	return nil, false
}

func NewRequest(contract giraffe.Contract, in interface{}) (*http.Request, error) {
	if pb, yes := isProtoMessage(in); yes {
		return newRequest(contract, pb)
	}
	return nil, errors.New("func NewRequest called with not proto.Message")
}

func marshalDataField(out io.Writer, name string, pb proto.Message) error {
	v := reflect.ValueOf(pb)
	if pb == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return errors.New("func marshalDataField called with nil")
	}
	s := v.Elem()
	for i := 0; i < s.NumField(); i++ {
		valueField := s.Type().Field(i)
		if strings.HasPrefix(valueField.Name, "XXX_") {
			continue
		}
		//this is not a protobuf field
		if valueField.Tag.Get("protobuf") == "" && valueField.Tag.Get("protobuf_oneof") == "" {
			continue
		}
		//not data field
		prop := fieldProperties(valueField)
		if prop.OrigName != name {
			continue
		}

		value := s.Field(i)
		if pb, yes := isProtoMessage(value.Interface()); yes {
			return jsonpbMarshaler.Marshal(out, pb)
		}
		return fmt.Errorf("data field %s was not proto.Message", strconv.Quote(name))
	}
	return fmt.Errorf("data field %s was not found", strconv.Quote(name))
}

func fieldProperties(f reflect.StructField) *proto.Properties {
	var prop proto.Properties
	prop.Init(f.Type, f.Name, f.Tag.Get("protobuf"), &f)
	return &prop
}

func newRequest(contract giraffe.Contract, pb proto.Message) (*http.Request, error) {
	rule, err := getHTTPRule(contract)
	if err != nil {
		return nil, err
	}

	url, err := urlpb.ParseURL(rule.path, pb, rule.method == "GET" || rule.method == "DELETE")
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	switch {
	case rule.body != "":
		out := new(bytes.Buffer)
		if err := marshalDataField(out, rule.body, pb); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(out.Bytes())
	default:
		out := new(bytes.Buffer)
		if err := jsonpbMarshaler.Marshal(out, pb); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(out.Bytes())
	}


	request, err := http.NewRequest(rule.method, url.String(), reader)
	if err != nil {
		return nil, err
	}

	if reader != nil {
		request.Header.Add("Content-Type", "application/json")
	}

	return request, nil
}
