package auth

import (
	"net/http"
	"strconv"
)

type transport struct {
	rt http.RoundTripper
}

type TransportOption func(*transport)

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	ctx := r.Context()
	if user, ok := FromContext(ctx); ok {
		r.Header.Add("user", user.User)
		r.Header.Add("org", strconv.Itoa(user.Org))
	}
	return t.rt.RoundTrip(r)
}

func NewTransport(options ...TransportOption) http.RoundTripper {
	t := &transport{
		rt: http.DefaultTransport,
	}
	for _, o := range options {
		o(t)
	}
	return t
}

func RoundTripper(rt http.RoundTripper) TransportOption {
	return func(t *transport) {
		t.rt = rt
	}
}
