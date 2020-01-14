package restv2

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/test/mock/mock_restv2"
)

type errReadCloser struct{}

func (*errReadCloser) Read(p []byte) (n int, err error) { return 0, errors.New("always error") }
func (*errReadCloser) Close() error                     { return nil }

func TestBaseMiddleware_NewRequest(t *testing.T) {
	type fields struct {
		Marshaler   Marshaler
		Unmarshaler Unmarshaler
	}
	type args struct {
		rule giraffe.HttpRule
		in   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
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
				in: &mock_restv2.GetDetailRequest{
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
				in: &mock_restv2.DeleteInstanceRequest{
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
				in: &mock_restv2.CreateInstanceRequest{
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
				in: &mock_restv2.UpdateInstanceRequest{
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
				in: &mock_restv2.UpdateInstanceRequest{
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
				in:   &mock_restv2.UpdateInstanceRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_HttpRuleNotDefined",
			args: args{
				rule: nil,
				in:   &mock_restv2.UpdateInstanceRequest{},
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
				in: (*mock_restv2.GetDetailRequest)(nil),
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
				in: &mock_restv2.CreateInstanceRequest{
					ObjectId: "APP",
					Instance: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_NoBodyField",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/object/:objectId/instance",
					},
					Body: "instance2",
				},
				in: &mock_restv2.CreateInstanceRequest{
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
				in: &mock_restv2.GetDetailRequest{
					ObjectId:   "APP",
					InstanceId: "abc",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_PostWithNilProtoMessage",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/",
					},
					Body: "",
				},
				in: (*mock_restv2.CreateInstanceRequest)(nil),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_PostWithNilProtoMessage2",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/",
					},
					Body: "instance",
				},
				in: (*mock_restv2.CreateInstanceRequest)(nil),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_GetDetailRequestWrapper",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/",
					},
					Body: "wrapper",
				},
				in: &mock_restv2.GetDetailRequestWrapper{
					Wrapper: mock_restv2.GetDetailRequest{
						ObjectId:   "APP",
						InstanceId: "abc",
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("POST", "/", bytes.NewReader([]byte("{\"objectId\":\"APP\",\"instanceId\":\"abc\"}")))
				r.Header.Add("content-type", "application/json")
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_GetDetailRequestWrapper2",
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/",
					},
					Body: "Data",
				},
				in: &mock_restv2.GetDetailRequestWrapper{
					Data: []byte{},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				rule: &giraffeproto.HttpRule{
					Pattern: &giraffeproto.HttpRule_Post{
						Post: "/v2/import/instance",
					},
					Body: "instances",
				},
				in: &mock_restv2.MultiCreateInstanceRequest{
					Instances: []*mock_restv2.CreateInstanceRequest{
						{
							ObjectId: "APP",
							Instance: &types.Struct{
								Fields: map[string]*types.Value{
									"name": {Kind: &types.Value_StringValue{StringValue: "abc"}},
								},
							},
						},
						{
							ObjectId: "APP",
							Instance: &types.Struct{
								Fields: map[string]*types.Value{
									"name": {Kind: &types.Value_StringValue{StringValue: "def"}},
								},
							},
						},
					},
				},
			},
			want: func() *http.Request {
				r, _ := http.NewRequest("POST", "/v2/import/instance", bytes.NewReader([]byte("[{\"objectId\":\"APP\",\"instance\":{\"name\":\"abc\"}},{\"objectId\":\"APP\",\"instance\":{\"name\":\"def\"}}]")))
				r.Header.Add("Content-Type", "application/json")
				return r
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BaseMiddleware{
				Marshaler:   tt.fields.Marshaler,
				Unmarshaler: tt.fields.Unmarshaler,
			}
			got, err := m.NewRequest(tt.args.rule, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.GetBody = nil
			}
			if tt.want != nil {
				tt.want.GetBody = nil
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMiddleware_ParseResponse(t *testing.T) {
	type fields struct {
		Marshaler   Marshaler
		Unmarshaler Unmarshaler
	}
	type args struct {
		rule giraffe.HttpRule
		resp *http.Response
		out  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
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
				resp: (&httptest.ResponseRecorder{
					Code: http.StatusCreated,
					Body: bytes.NewBuffer([]byte("{" +
						"\"code\": 0," +
						"\"codeExplain\": \"succeed\"," +
						"\"message\": \"成功\"," +
						"\"data\": {}" +
						"}")),
				}).Result(),
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
				resp: (&httptest.ResponseRecorder{
					Code: http.StatusCreated,
					Body: bytes.NewBuffer([]byte("[" + "]")),
				}).Result(),
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
				resp: (&httptest.ResponseRecorder{
					Code: http.StatusNotFound,
					Body: nil,
				}).Result(),
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
				resp: (&httptest.ResponseRecorder{
					Code: http.StatusNoContent,
					Body: bytes.NewBuffer([]byte("")),
				}).Result(),
				out: new(types.Struct),
			},
			want:    &types.Struct{},
			wantErr: true,
		},
		{
			name: "Test_WithNilOutput",
			args: args{
				rule: &giraffeproto.HttpRule{
					ResponseBody: "data",
				},
				resp: (&httptest.ResponseRecorder{
					Code: http.StatusNoContent,
					Body: bytes.NewBuffer([]byte("")),
				}).Result(),
				out: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BaseMiddleware{
				Marshaler:   tt.fields.Marshaler,
				Unmarshaler: tt.fields.Unmarshaler,
			}
			if err := m.ParseResponse(tt.args.rule, tt.args.resp, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("ParseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseMiddleware_unmarshaler(t *testing.T) {
	type mockUnmarshaler struct {
		Unmarshaler
	}
	type fields struct {
		Marshaler   Marshaler
		Unmarshaler Unmarshaler
	}
	tests := []struct {
		name   string
		fields fields
		want   Unmarshaler
	}{
		{
			name: "Test_DefaultUnmarshaler",
			fields: fields{
				Marshaler:   nil,
				Unmarshaler: nil,
			},
			want: defaultUnmarshaler,
		},
		{
			name: "Test_MockUnmarshaler",
			fields: fields{
				Marshaler:   nil,
				Unmarshaler: &mockUnmarshaler{},
			},
			want: &mockUnmarshaler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BaseMiddleware{
				Marshaler:   tt.fields.Marshaler,
				Unmarshaler: tt.fields.Unmarshaler,
			}
			if got := m.unmarshaler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshaler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseMiddleware_marshaler(t *testing.T) {
	type mockMarshaler struct {
		Marshaler
	}
	type fields struct {
		Marshaler   Marshaler
		Unmarshaler Unmarshaler
	}
	tests := []struct {
		name   string
		fields fields
		want   Marshaler
	}{
		{
			name: "Test_DefaultMarshaler",
			fields: fields{
				Marshaler:   nil,
				Unmarshaler: nil,
			},
			want: defaultMarshaler,
		},
		{
			name: "Test_MockMarshaler",
			fields: fields{
				Marshaler:   &mockMarshaler{},
				Unmarshaler: nil,
			},
			want: &mockMarshaler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BaseMiddleware{
				Marshaler:   tt.fields.Marshaler,
				Unmarshaler: tt.fields.Unmarshaler,
			}
			if got := m.marshaler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshaler() = %v, want %v", got, tt.want)
			}
		})
	}
}
