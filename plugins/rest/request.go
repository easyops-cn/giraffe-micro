package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/pkg/urlpb"

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

func getPattern(rule giraffe.HttpRule) (verb string, path string, err error) {
	switch {
	case rule.GetGet() != "":
		return http.MethodGet, rule.GetGet(), nil
	case rule.GetPost() != "":
		return http.MethodPost, rule.GetPost(), nil
	case rule.GetPut() != "":
		return http.MethodPut, rule.GetPut(), nil
	case rule.GetDelete() != "":
		return http.MethodDelete, rule.GetDelete(), nil
	case rule.GetPatch() != "":
		return http.MethodPatch, rule.GetPatch(), nil
	default:
		if r, ok := rule.(interface {
			GetCustom() *giraffeproto.CustomHttpPattern
		}); ok && r.GetCustom().GetPath() != "" {
			return r.GetCustom().GetKind(), r.GetCustom().GetPath(), nil
		}
	}
	return "", "", errors.New("http method was not defined")
}

func newRequest(md *giraffe.MethodDesc, in interface{}) (*http.Request, error) {
	pb, _ := isProtoMessage(in)

	rule := md.HttpRule
	if rule == nil {
		return nil, errors.New("http rule was nil")
	}

	verb, path, err := getPattern(rule)
	if err != nil {
		return nil, err
	}

	url, err := urlpb.ParseURL(path, pb, verb == http.MethodGet || verb == http.MethodDelete)
	if err != nil {
		return nil, err
	}

	var reader io.Reader
	switch {
	case rule.GetBody() != "":
		out := new(bytes.Buffer)
		if err := marshalDataField(out, rule.GetBody(), pb); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(out.Bytes())
	case verb != http.MethodGet && verb != http.MethodDelete:
		out := new(bytes.Buffer)
		if err := jsonpbMarshaler.Marshal(out, pb); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(out.Bytes())
	default:
		reader = nil
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
