package urlpb

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/gogo/protobuf/proto"
)

type query struct {
	url.Values
}

//Add 添加 query 参数
func (q *query) Add(key string, rv reflect.Value) {
	value := ""
	if v := fmt.Sprintf("%v", rv.Interface()); v != "<nil>" {
		value = v
	}
	q.Values.Add(key, value)
}

func marshal(pb proto.Message) (url.Values, error) {

	v := reflect.ValueOf(pb)
	if pb == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil, errors.New("func Marshal called with nil")
	}

	out := query{url.Values{}}
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

		value := s.Field(i)
		valueKind := value.Kind()
		prop := fieldProperties(valueField)

		switch {
		case isSimpleValueKind(valueKind):
			out.Add(prop.OrigName, value)
		case reflect.Array == valueKind, reflect.Slice == valueKind:
			if isSimpleValueKind(value.Type().Elem().Kind()) {
				for i := 0; i < value.Len(); i++ {
					out.Add(prop.OrigName, value.Index(i))
				}
			}
		default:
			// Struct Not Supported
		}
	}
	return out.Values, nil
}

func isSimpleValueKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}

func fieldProperties(f reflect.StructField) *proto.Properties {
	var prop proto.Properties
	prop.Init(f.Type, f.Name, f.Tag.Get("protobuf"), &f)
	return &prop
}

//Marshal 根据 message 解析成 url.Value, message 中只有第一层字段会被处理
func Marshal(pb proto.Message) (url.Values, error) {
	return marshal(pb)
}

//ParseURL 解释rawurl
func ParseURL(rawurl string, pb proto.Message, parseQuery bool) (*url.URL, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`/:\w+`)
	pathParams := reg.FindAllString(u.Path, -1)
	if len(pathParams) == 0 && !parseQuery {
		return u, nil
	}

	urlValue, err := marshal(pb)
	if err != nil {
		return nil, err
	}

	a := make([]interface{}, len(pathParams))
	for i, m := range pathParams {
		v, exist := urlValue[m[2:]]
		if !exist {
			a[i] = ""
			continue
		}
		a[i] = strings.Join(v, ",")
		urlValue.Del(m[2:])
	}
	pathFormat := reg.ReplaceAllString(u.Path, "/%s")
	u.Path = fmt.Sprintf(pathFormat, a...)

	if !parseQuery {
		return u, nil
	}

	addQuery(u, urlValue)

	return u, nil
}

func addQuery(u *url.URL, urlValue url.Values) {
	if len(u.RawQuery) == 0 {
		u.RawQuery = urlValue.Encode()
		return
	}

	query := u.Query()
	for k, a := range urlValue {
		for _, v := range a {
			query.Add(k, v)
		}
	}
	u.RawQuery = query.Encode()
}
