package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-test/deep"
)

var resp200 = (&httptest.ResponseRecorder{Code: http.StatusOK}).Result()
var resp404 = (&httptest.ResponseRecorder{Code: http.StatusNotFound}).Result()

func Test_restError_Error(t *testing.T) {
	tests := []struct {
		name string
		r    *restError
		want string
	}{
		{
			name: "Test_HappyPath",
			r:    (*restError)(resp404),
			want: resp404.Status,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := tt.r.Error(); got != tt.want {
			//	t.Errorf("restError.Error() = %v, want %v", got, tt.want)
			//}
			got := tt.r.Error()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_restError_HttpResponse(t *testing.T) {
	tests := []struct {
		name string
		r    *restError
		want *http.Response
	}{
		{
			name: "Test_HappyPath",
			r:    (*restError)(resp404),
			want: resp404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := tt.r.HttpResponse(); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("restError.HttpResponse() = %v, want %v", got, tt.want)
			//}
			got := tt.r.HttpResponse()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_isErrorResponse(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want *restError
	}{
		{
			name: "Test_HappyPath_StatusOK",
			args: args{
				resp: resp200,
			},
			want: nil,
		},
		{
			name: "Test_HappyPath_StatusNotFound",
			args: args{
				resp: resp404,
			},
			want: (*restError)(resp404),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := isErrorResponse(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("isErrorResponse() = %v, want %v", got, tt.want)
			//}
			got := isErrorResponse(tt.args.resp)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
