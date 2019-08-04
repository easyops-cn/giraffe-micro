package restv2

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
)

func Test_parseResponse(t *testing.T) {
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
			if tt.args.resp == nil {
				tt.args.resp = tt.args.respRec.Result()
			}
			if err := parseResponse(tt.args.rule, tt.args.resp, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("parseResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(tt.args.out, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
