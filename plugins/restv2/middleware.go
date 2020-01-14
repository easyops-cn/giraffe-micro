package restv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/pkg/urlpb"
)

var defaultMarshaler = &jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
}

var defaultUnmarshaler = &jsonpb.Unmarshaler{
	AllowUnknownFields: true,
}

//DefaultMiddleware 默认中间件
var DefaultMiddleware = &BaseMiddleware{}

//Marshaler 序列化接口
type Marshaler interface {
	Marshal(out io.Writer, pb proto.Message) error
}

//Unmarshaler 反序列化接口
type Unmarshaler interface {
	Unmarshal(r io.Reader, pb proto.Message) error
}

//BaseMiddleware 基础中间件
type BaseMiddleware struct {
	Marshaler   Marshaler
	Unmarshaler Unmarshaler
}

//NewRequest 创建 http.Request 方法
func (m *BaseMiddleware) NewRequest(rule giraffe.HttpRule, in interface{}) (*http.Request, error) {
	pb, ispb := in.(proto.Message)
	if !ispb {
		return nil, errors.New("in interface{} was not proto.Message")
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
		if err := m.marshalDataField(out, rule.GetBody(), pb); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(out.Bytes())
	case verb != http.MethodGet && verb != http.MethodDelete:
		out := new(bytes.Buffer)
		if err := m.marshaler().Marshal(out, pb); err != nil {
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

//ParseResponse 解析 http.Response 方法
func (m *BaseMiddleware) ParseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error {
	pb, ispb := out.(proto.Message)
	if !ispb {
		return errors.New("in interface{} was not proto.Message")
	}

	// 读取 response body
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unexpected error when reading response body: %s", err.Error())
	}

	// 如果指定了 response message 字段, 则提取字段
	if body := rule.GetResponseBody(); body != "" {
		m := map[string]json.RawMessage{}
		if err := json.Unmarshal(bodyData, &m); err != nil {
			return fmt.Errorf("unmarshal response body failed: %s", err)
		}
		bodyData = m[body]
	}

	// 错误码正常, 解析 message
	if err := m.unmarshaler().Unmarshal(bytes.NewReader(bodyData), pb); err != nil {
		return fmt.Errorf("unexpected message when unmarshal body data: %s", err.Error())
	}
	return nil
}

func (m *BaseMiddleware) marshalDataField(out io.Writer, name string, pb proto.Message) error {
	v := reflect.ValueOf(pb)
	if pb == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return errors.New("func marshalDataField called with nil")
	}
	var field reflect.Value
	for i, prop := range proto.GetProperties(v.Elem().Type()).Prop {
		if prop.OrigName == name {
			field = v.Elem().Field(i)
			break
		}
	}
	if !field.IsValid() {
		return fmt.Errorf("data field %s was not found", strconv.Quote(name))
	}
	pbType := reflect.TypeOf((*proto.Message)(nil)).Elem()
	if field.Kind() == reflect.Ptr && field.Type().Implements(pbType) {
		return m.marshaler().Marshal(out, field.Interface().(proto.Message))
	} else if field.Kind() == reflect.Struct && reflect.PtrTo(field.Type()).Implements(pbType) {
		return m.marshaler().Marshal(out, field.Addr().Interface().(proto.Message))
	} else if field.Kind() == reflect.Slice && field.Type().Elem().Implements(pbType) {
		s := make([][]byte, 0, field.Len())
		var err error
		for i := 0; i < field.Len(); i++ {
			if err == nil {
				buff := bytes.NewBuffer([]byte{})
				err = m.marshaler().Marshal(buff, field.Index(i).Interface().(proto.Message))
				s = append(s, buff.Bytes())
			}
		}
		if err == nil {
			_, err = out.Write([]byte{'['})
		}
		if err == nil {
			_, err = out.Write(bytes.Join(s, []byte{','}))
		}
		if err == nil {
			_, err = out.Write([]byte{']'})
		}
		return err
	}
	return fmt.Errorf("data field %s was not proto.Message", strconv.Quote(name))
}

func (m *BaseMiddleware) unmarshaler() Unmarshaler {
	if m.Unmarshaler != nil {
		return m.Unmarshaler
	}
	return defaultUnmarshaler
}

func (m *BaseMiddleware) marshaler() Marshaler {
	if m.Marshaler != nil {
		return m.Marshaler
	}
	return defaultMarshaler
}

func getPattern(rule giraffe.HttpRule) (verb string, path string, err error) {
	if rule == nil {
		return "", "", errors.New("http rule was nil")
	}
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
