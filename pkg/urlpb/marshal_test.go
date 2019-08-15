package urlpb

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

type structType struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *structType) Reset()         { *m = structType{} }
func (m *structType) String() string { return proto.CompactTextString(m) }
func (*structType) ProtoMessage()    {}

type queryMessage struct {
	StringValue          string      `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BoolValue            bool        `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Int32Value           int32       `protobuf:"varint,3,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value           int64       `protobuf:"varint,4,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint32Value          uint32      `protobuf:"varint,5,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value          uint64      `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	FloatValue           float32     `protobuf:"fixed32,7,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue          float64     `protobuf:"fixed64,8,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	StructValue          *structType `protobuf:"bytes,9,opt,name=struct_value,json=structValue,proto3" json:"struct_value,omitempty"`
	SliceValue           []string    `protobuf:"bytes,10,rep,name=slice_value,json=sliceValue,proto3" json:"slice_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *queryMessage) Reset()         { *m = queryMessage{} }
func (m *queryMessage) String() string { return proto.CompactTextString(m) }
func (*queryMessage) ProtoMessage()    {}

type sampleMessage struct {
	Name                 string   `json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *sampleMessage) Reset()         { *m = sampleMessage{} }
func (m *sampleMessage) String() string { return proto.CompactTextString(m) }
func (*sampleMessage) ProtoMessage()    {}

var queryData = &queryMessage{
	StringValue: "abc",
	BoolValue:   true,
	Int32Value:  -123,
	Int64Value:  -456,
	Uint32Value: 789,
	Uint64Value: 10,
	FloatValue:  1.0,
	DoubleValue: 1.01,
	StructValue: &structType{Name: "index"},
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
				pb: (*queryMessage)(nil),
			},
			wantErr: true,
		},
		{
			name: "TestWithoutProtobufTag",
			args: args{
				pb: &sampleMessage{Name: "aa"},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
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
				pb:         (*queryMessage)(nil),
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
			if !reflect.DeepEqual(gotString, tt.want) {
				t.Errorf("ParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
