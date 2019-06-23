package easyopsctx

import (
	"context"
	"reflect"
	"testing"
)

func TestWithUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		u   UserInfo
	}
	tests := []struct {
		name string
		args args
		want UserInfo
	}{
		{
			name: "TestHappyPath",
			args: args{
				ctx: context.Background(),
				u: UserInfo{
					User: "indexzhuo",
					Org:  8888,
				},
			},
			want: UserInfo{
				User: "indexzhuo",
				Org:  8888,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithUserInfo(tt.args.ctx, tt.args.u)
			got := FromContext(ctx)
			if !reflect.DeepEqual(got, tt.want) {
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
		name string
		args args
		want UserInfo
	}{
		{
			name: "TestBackgroundContext",
			args: args{
				ctx: context.Background(),
			},
			want: UserInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
