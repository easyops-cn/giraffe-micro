package easyopsctx

import "context"

type key struct {}

var ctxKey = key{}

type UserInfo struct {
	User string
	Org int
}

func WithUserInfo(ctx context.Context, u UserInfo) context.Context {
	return context.WithValue(ctx, ctxKey, u)
}

func FromContext(ctx context.Context) UserInfo {
	if r, ok := ctx.Value(ctxKey).(UserInfo); ok {
		return r
	} else {
		return r
	}
}
