package restv2

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
)

type getDetailRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *getDetailRequest) Reset()         { *m = getDetailRequest{} }
func (m *getDetailRequest) String() string { return proto.CompactTextString(m) }
func (*getDetailRequest) ProtoMessage()    {}

type deleteInstanceRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *deleteInstanceRequest) Reset()         { *m = deleteInstanceRequest{} }
func (m *deleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*deleteInstanceRequest) ProtoMessage()    {}

type createInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	Instance             *types.Struct `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *createInstanceRequest) Reset()         { *m = createInstanceRequest{} }
func (m *createInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*createInstanceRequest) ProtoMessage()    {}

type updateInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string        `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Instance             *types.Struct `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *updateInstanceRequest) Reset()         { *m = updateInstanceRequest{} }
func (m *updateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*updateInstanceRequest) ProtoMessage()    {}

type getDetailRequestWrapper struct {
	Data                 []byte
	ObjectId             string            `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string            `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Wrapper              *getDetailRequest `protobuf:"bytes,2,opt,name=wrapper,proto3" json:"wrapper"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *getDetailRequestWrapper) Reset()         { *m = getDetailRequestWrapper{} }
func (m *getDetailRequestWrapper) String() string { return proto.CompactTextString(m) }
func (m *getDetailRequestWrapper) ProtoMessage()  {}

type errReadCloser struct{}

func (*errReadCloser) Read(p []byte) (n int, err error) { return 0, errors.New("always error") }
func (*errReadCloser) Close() error                     { return nil }

func Test_middleware_NewRequest(t *testing.T) {
	type args struct {
		rule giraffe.HttpRule
		in   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "Test_HappyPath_GET",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Get{
						Get: "/object/:objectId/instance/:instanceId",
					},
					Body: "",
				},
				in: &getDetailRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("GET", "/object/APP/instance/abc", nil)
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_HappyPath_DELETE",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Delete{
						Delete: "/object/:objectId/instance/:instanceId",
					},
					Body: "",
				},
				in: &deleteInstanceRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("DELETE", "/object/APP/instance/abc", nil)
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_HappyPath_POST_With_Body",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/object/:objectId/instance",
					},
					Body: "instance",
				},
				in: &createInstanceRequest{
					ObjectId: "APP",
					Instance: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {Kind: &types.Value_StringValue{StringValue: "abc"}},
						},
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("POST", "/v2/object/APP/instance", bytes.NewReader([]byte("{\"name\":\"abc\"}")))
				r.Header.Add("Content-Type", "application/json")
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_HappyPath_POST",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/object/APP/instance",
					},
				},
				in: &types.Struct{
					Fields: map[string]*types.Value{
						"name": {Kind: &types.Value_StringValue{StringValue: "abc"}},
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("POST", "/v2/object/APP/instance", bytes.NewReader([]byte("{\"name\":\"abc\"}")))
				r.Header.Add("Content-Type", "application/json")
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_HappyPath_PUT",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Put{
						Put: "/object/:objectId/instance/:instanceId",
					},
					Body: "instance",
				},
				in: &updateInstanceRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
					Instance: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {Kind: &types.Value_StringValue{StringValue: "abc"}},
						},
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("PUT", "/object/APP/instance/abc", bytes.NewReader([]byte("{\"name\":\"abc\"}")))
				r.Header.Add("Content-Type", "application/json")
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_HappyPath_PATCH",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Patch{
						Patch: "/object/:objectId/instance/:instanceId",
					},
					Body: "instance",
				},
				in: &updateInstanceRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
					Instance: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {Kind: &types.Value_StringValue{StringValue: "abc"}},
						},
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("PATCH", "/object/APP/instance/abc", bytes.NewReader([]byte("{\"name\":\"abc\"}")))
				r.Header.Add("Content-Type", "application/json")
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_MethodNotDefined",
			args: args{
				rule: &giraffeproto.HttpRule{},
				in:   &updateInstanceRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_HttpRuleNotDefined",
			args: args{
				rule: nil,
				in:   &updateInstanceRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_WithNilProtoMessage",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Get{
						Get: "/object/:objectId/instance/:instanceId",
					},
					Body: "",
				},
				in: (*getDetailRequest)(nil),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_WithNilMessage",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/instance",
					},
				},
				in: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_WithNilBody",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/object/:objectId/instance",
					},
					Body: "instance",
				},
				in: &createInstanceRequest{
					ObjectId: "APP",
					Instance: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_WithWrongMethod",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Custom{
						Custom: &giraffeproto.CustomHttpPattern{
							Kind: " ",
							Path: "/xxx",
						},
					},
					Body: "",
				},
				in: &getDetailRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &middleware{}
			got, err := m.NewRequest(tt.args.rule, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_middleware_ParseResponse(t *testing.T) {
	type args struct {
		rule    giraffe.HttpRule
		resp    *http.Response
		respRec *httptest.ResponseRecorder
		out     interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			args: args{
				rule: &giraffeproto.HttpRule{
					ResponseBody: "data",
				},
				respRec: &httptest.ResponseRecorder{
					Code: http.StatusCreated,
					Body: bytes.NewBuffer([]byte("{" +
						"\"code\": 0," +
						"\"codeExplain\": \"succeed\"," +
						"\"message\": \"成功\"," +
						"\"data\": {}" +
						"}")),
				},
				out: new(types.Struct),
			},
			want: &types.Struct{
				Fields: map[string]*types.Value{},
			},
			wantErr: false,
		},
		{
			name: "Test_UnexpectedMessage",
			args: args{
				rule: &giraffeproto.HttpRule{},
				respRec: &httptest.ResponseRecorder{
					Code: http.StatusCreated,
					Body: bytes.NewBuffer([]byte("[" + "]")),
				},
				out: new(types.Struct),
			},
			want:    &types.Struct{},
			wantErr: true,
		},
		{
			name: "Test_404",
			args: args{
				rule: &giraffeproto.HttpRule{
					ResponseBody: "data",
				},
				respRec: &httptest.ResponseRecorder{
					Code: http.StatusNotFound,
				},
				out: new(types.Struct),
			},
			want:    &types.Struct{},
			wantErr: true,
		},
		{
			name: "Test_ReadResponseBodyFailed",
			args: args{
				rule: &giraffeproto.HttpRule{
					ResponseBody: "data",
				},
				resp: &http.Response{
					Body: &errReadCloser{},
				},
				out: new(types.Struct),
			},
			want:    &types.Struct{},
			wantErr: true,
		},
		{
			name: "Test_ReadDataFailed",
			args: args{
				rule: &giraffeproto.HttpRule{
					ResponseBody: "data",
				},
				respRec: &httptest.ResponseRecorder{
					Code: http.StatusNoContent,
					Body: bytes.NewBuffer([]byte("")),
				},
				out: new(types.Struct),
			},
			want:    &types.Struct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &middleware{}
			if tt.args.resp == nil {
				tt.args.resp = tt.args.respRec.Result()
			}
			if err := m.ParseResponse(tt.args.rule, tt.args.resp, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("ParseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := deep.Equal(tt.args.out, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
