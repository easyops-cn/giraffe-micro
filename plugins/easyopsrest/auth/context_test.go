package auth

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

var (
	ctxWithValue = context.WithValue(context.Background(), ctxKey, UserInfo{User:"foo",Org:8888})
	userInfo = UserInfo{User:"foo", Org:8888}
)

func TestWithUserInfo(t *testing.T) {
	type args struct {
		ctx  context.Context
		info UserInfo
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "Test_HappyPath",
			args: args{
				ctx: context.Background(),
				info: userInfo,
			},
			want: ctxWithValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithUserInfo(tt.args.ctx, tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name  string
		args  args
		want  UserInfo
		want1 bool
	}{
		{
			name: "Test_HappyPath",
			args: args{
				ctxWithValue,
			},
			want: userInfo,
			want1: true,
		},
		{
			name: "Test_WithoutUserInfo",
			args: args{
				context.Background(),
			},
			want: UserInfo{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FromContext(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FromContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFromRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "Test_HappyPath",
			args: args{
				req: func() *http.Request {
					r, _ := http.NewRequest("GET", "/", nil)
					r.Header.Add("user", "foo")
					r.Header.Add("org", "8888")
					return r
				}(),
			},
			want: ctxWithValue,
		},
		{
			name: "Test_WithoutUserInfo",
			args: args{
				req: func() *http.Request {
					r, _ := http.NewRequest("GET", "/", nil)
					return r
				}(),
			},
			want: context.Background(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromRequest(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
