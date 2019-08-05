package rest

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
)

type errReadCloser struct{}

func (*errReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("always error")
}

func (*errReadCloser) Close() error {
	return nil
}

func Test_parseResponse(t *testing.T) {
	type args struct {
		md      *giraffe.MethodDesc
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
				md: &giraffe.MethodDesc{
					HttpRule: &giraffeproto.HttpRule{
						ResponseBody: "data",
					},
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
				md: &giraffe.MethodDesc{
					HttpRule: &giraffeproto.HttpRule{},
				},
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
				md: &giraffe.MethodDesc{
					HttpRule: &giraffeproto.HttpRule{
						ResponseBody: "data",
					},
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
				md: &giraffe.MethodDesc{
					HttpRule: &giraffeproto.HttpRule{
						ResponseBody: "data",
					},
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
				md: &giraffe.MethodDesc{
					HttpRule: &giraffeproto.HttpRule{
						ResponseBody: "data",
					},
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
			if tt.args.resp == nil {
				tt.args.resp = tt.args.respRec.Result()
			}
			if err := parseResponse(tt.args.md, tt.args.resp, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("parseResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(tt.args.out, tt.want) {
			//	t.Errorf("parseResponse() = %v, want %v", tt.args.out, tt.want)
			//}
			if diff := deep.Equal(tt.args.out, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
