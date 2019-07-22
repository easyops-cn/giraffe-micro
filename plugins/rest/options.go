package rest

import (
	"net/http"
	"time"

	"github.com/openzipkin/zipkin-go"

	"github.com/easyops-cn/giraffe-micro"
)

type ClientOptions struct {
	nameService giraffe.NameService
	timeout     time.Duration
	tracer      *zipkin.Tracer
	rt          http.RoundTripper
}

type ClientOption func(o *ClientOptions)

func newClientOptions(opts ...ClientOption) ClientOptions {
	opt := ClientOptions{
		timeout: time.Second * 60,
		tracer:  nil,
		rt:      http.DefaultTransport,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.timeout = timeout
	}
}

func WithNameService(ns giraffe.NameService) ClientOption {
	return func(o *ClientOptions) {
		o.nameService = ns
	}
}

func WithTracer(tracer *zipkin.Tracer) ClientOption {
	return func(o *ClientOptions) {
		o.tracer = tracer
	}
}

func WithRoundTripper(rt http.RoundTripper) ClientOption {
	return func(o *ClientOptions) {
		o.rt = rt
	}
}
