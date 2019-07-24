package urlpb

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/go-test/deep"
	"github.com/gogo/protobuf/proto"

	"github.com/easyops-cn/giraffe-micro/pkg/urlpb/mock"
)

var queryData = &mock.Query{
	StringValue: "abc",
	BoolValue:   true,
	Int32Value:  -123,
	Int64Value:  -456,
	Uint32Value: 789,
	Uint64Value: 10,
	FloatValue:  1.0,
	DoubleValue: 1.01,
	StructValue: &mock.S{Name: "index"},
	SliceValue:  []string{"a", "b", "c"},
}

func TestMarshal(t *testing.T) {
	type args struct {
		pb proto.Message
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr bool
	}{
		{
			name: "TestHappyPath",
			args: args{
				pb: queryData,
			},
			want: url.Values{
				"string_value": []string{"abc"},
				"bool_value":   []string{"true"},
				"int32_value":  []string{"-123"},
				"int64_value":  []string{"-456"},
				"uint32_value": []string{"789"},
				"uint64_value": []string{"10"},
				"float_value":  []string{"1"},
				"double_value": []string{"1.01"},
				"slice_value":  []string{"a", "b", "c"},
			},
			wantErr: false,
		},
		{
			name: "TestNilProtoMessage",
			args: args{
				pb: (*mock.Query)(nil),
			},
			wantErr: true,
		},
		{
			name: "TestWithoutProtobufTag",
			args: args{
				pb: &mock.SS{Name: "aa"},
			},
			want:    url.Values{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Marshal() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestParseURL(t *testing.T) {
	type args struct {
		rawurl     string
		pb         proto.Message
		parseQuery bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestHappyPath",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/:string_value/bool/:bool_value?a=123",
				pb:         queryData,
				parseQuery: true,
			},
			want:    "http://127.0.0.1:8080/string/abc/bool/true?a=123&double_value=1.01&float_value=1&int32_value=-123&int64_value=-456&slice_value=a&slice_value=b&slice_value=c&uint32_value=789&uint64_value=10",
			wantErr: false,
		},
		{
			name: "TestWithNoQueryParam",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/:string_value/bool/:bool_value",
				pb:         queryData,
				parseQuery: true,
			},
			want:    "http://127.0.0.1:8080/string/abc/bool/true?double_value=1.01&float_value=1&int32_value=-123&int64_value=-456&slice_value=a&slice_value=b&slice_value=c&uint32_value=789&uint64_value=10",
			wantErr: false,
		},
		{
			name: "TestWithoutParseQuery",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/:string_value/bool/:bool_value?a=123",
				pb:         queryData,
				parseQuery: false,
			},
			want:    "http://127.0.0.1:8080/string/abc/bool/true?a=123",
			wantErr: false,
		},
		{
			name: "TestNormalURLWithoutParseQuery",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/abc/bool/true?a=123",
				pb:         queryData,
				parseQuery: false,
			},
			want:    "http://127.0.0.1:8080/string/abc/bool/true?a=123",
			wantErr: false,
		},
		{
			name: "TestWithWrongURLParams",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/:wrong_param/bool/:bool_value?a=123",
				pb:         queryData,
				parseQuery: true,
			},
			want:    "http://127.0.0.1:8080/string//bool/true?a=123&double_value=1.01&float_value=1&int32_value=-123&int64_value=-456&slice_value=a&slice_value=b&slice_value=c&string_value=abc&uint32_value=789&uint64_value=10",
			wantErr: false,
		},
		{
			name: "TestWithWrongURL",
			args: args{
				rawurl:     ":wrong url",
				pb:         queryData,
				parseQuery: true,
			},
			want:    "<nil>",
			wantErr: true,
		},
		{
			name: "TestWithNilProtoMessage",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/abc/bool/true?a=123",
				pb:         (*mock.Query)(nil),
				parseQuery: true,
			},
			want:    "<nil>",
			wantErr: true,
		},
		{
			name: "TestWithoutURLParams",
			args: args{
				rawurl:     "http://127.0.0.1:8080/string/abc/bool/true?a=123",
				pb:         queryData,
				parseQuery: true,
			},
			want:    "http://127.0.0.1:8080/string/abc/bool/true?a=123&bool_value=true&double_value=1.01&float_value=1&int32_value=-123&int64_value=-456&slice_value=a&slice_value=b&slice_value=c&string_value=abc&uint32_value=789&uint64_value=10",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.args.rawurl, tt.args.pb, tt.args.parseQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotString := fmt.Sprint(got)
			//if !reflect.DeepEqual(gotString, tt.want) {
			//	t.Errorf("ParseURL() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(gotString, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
