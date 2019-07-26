package auth

import (
	"context"
	"net/http"
	"strconv"
)

type keyType struct{}

var ctxKey = keyType{}

//WithUserInfo
func WithUserInfo(ctx context.Context, info UserInfo) context.Context {
	return context.WithValue(ctx, ctxKey, info)
}

//FromContext
func FromContext(ctx context.Context) (UserInfo, bool) {
	v, ok := ctx.Value(ctxKey).(UserInfo)
	return v, ok
}

//FromRequest
func FromRequest(req *http.Request) context.Context {
	user := req.Header.Get("user")
	org, _ := strconv.Atoi(req.Header.Get("org"))
	if user == "" || org == 0 {
		return req.Context()
	}
	return context.WithValue(req.Context(), ctxKey, UserInfo{user, org})
}
