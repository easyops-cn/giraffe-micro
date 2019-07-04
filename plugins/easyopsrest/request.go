package easyopsrest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/urlpb"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

var jsonpbMarshaler = jsonpb.Marshaler{OrigName: true}

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

func NewRequest(method giraffe.Method, in interface{}) (*http.Request, error) {
	if pb, yes := isProtoMessage(in); yes {
		return newRequest(method, pb)
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

func newRequest(method giraffe.Method, pb proto.Message) (*http.Request, error) {
	rule, ok := method.(giraffe.HttpRule)
	if !ok {
		return nil, errors.New("method was not implement giraffe.HttpRule")
	}

	verb, path := rule.Pattern()
	url, err := urlpb.ParseURL(path, pb, verb == "GET" || verb == "DELETE")
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	switch {
	case rule.Body() != "":
		out := new(bytes.Buffer)
		if err := marshalDataField(out, rule.Body(), pb); err != nil {
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

	request, err := http.NewRequest(verb, url.String(), reader)
	if err != nil {
		return nil, err
	}

	if reader != nil {
		request.Header.Add("Content-Type", "application/json")
	}

	return request, nil
}
