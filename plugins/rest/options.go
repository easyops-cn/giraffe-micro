package rest

import (
	"net/http"
	"time"

	"github.com/openzipkin/zipkin-go"

	"github.com/easyops-cn/giraffe-micro"
)

type ClientOptions struct {
	NameService giraffe.ContractService
	Timeout     time.Duration
	Tracer      *zipkin.Tracer
	Transport   http.RoundTripper
}

type ClientOption func(o *ClientOptions)

func newClientOptions(opts ...ClientOption) ClientOptions {
	opt := ClientOptions{
		Timeout:   time.Second * 60,
		Tracer:    nil,
		Transport: http.DefaultTransport,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.Timeout = timeout
	}
}

func WithNameService(ns giraffe.ContractService) ClientOption {
	return func(o *ClientOptions) {
		o.NameService = ns
	}
}

func WithTracer(tracer *zipkin.Tracer) ClientOption {
	return func(o *ClientOptions) {
		o.Tracer = tracer
	}
}

func WithRoundTripper(rt http.RoundTripper) ClientOption {
	return func(o *ClientOptions) {
		o.Transport = rt
	}
}
