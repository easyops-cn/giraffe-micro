package restv2

import (
	"bytes"
	"net/http"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
)

func Test_newRequest(t *testing.T) {
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
			got, err := newRequest(tt.args.rule, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("newRequest() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_isProtoMessage(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  proto.Message
		want1 bool
	}{
		{
			name: "Test_HappyPath",
			args: args{
				v: createInstanceRequest{},
			},
			want:  &createInstanceRequest{},
			want1: true,
		},
		{
			name: "Test_HappyPath2",
			args: args{
				v: &createInstanceRequest{},
			},
			want:  &createInstanceRequest{},
			want1: true,
		},
		{
			name: "Test_NotProtoMessage",
			args: args{
				v: new(int),
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isProtoMessage(tt.args.v)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("isProtoMessage() got = %v, want %v", got, tt.want)
			//}
			//if got1 != tt.want1 {
			//	t.Errorf("isProtoMessage() got1 = %v, want %v", got1, tt.want1)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
			if diff := deep.Equal(got1, tt.want1); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_marshalDataField(t *testing.T) {
	type args struct {
		name string
		pb   proto.Message
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "Test_CallWithNil",
			args: args{
				name: "data",
				pb:   nil,
			},
			wantOut: "",
			wantErr: true,
		},
		{
			name: "Test_WithProtoWrapper",
			args: args{
				name: "wrapper",
				pb: &getDetailRequestWrapper{
					Wrapper: &getDetailRequest{},
				},
			},
			wantOut: "{}",
			wantErr: false,
		},
		{
			name: "Test_WithWrongDataField",
			args: args{
				name: "xxx",
				pb: &getDetailRequestWrapper{
					Wrapper: &getDetailRequest{},
				},
			},
			wantOut: "",
			wantErr: true,
		},
		{
			name: "Test_DataFieldWasNotProtoMessage",
			args: args{
				name: "objectId",
				pb: &getDetailRequestWrapper{
					Wrapper: &getDetailRequest{},
				},
			},
			wantOut: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := marshalDataField(out, tt.args.name, tt.args.pb); (err != nil) != tt.wantErr {
				t.Errorf("marshalDataField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if gotOut := out.String(); gotOut != tt.wantOut {
			//	t.Errorf("marshalDataField() = %v, want %v", gotOut, tt.wantOut)
			//}
			gotOut := out.String()
			if diff := deep.Equal(gotOut, tt.wantOut); diff != nil {
				t.Error(diff)
			}
		})
	}
}
