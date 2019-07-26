package rest

import (
	"net/http"
	"time"

	"github.com/openzipkin/zipkin-go"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/ens"
	"github.com/easyops-cn/giraffe-micro/plugins/rest"
)

//ClientOption 客户端配置数据结构
//Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
type ClientOption func(o *rest.ClientOptions)

func newClientOptions(opts ...ClientOption) rest.ClientOptions {
	opt := rest.ClientOptions{
		Timeout:     time.Second * 60,
		Tracer:      nil,
		NameService: ens.NewNameService(),
		Transport:   http.DefaultTransport,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

//WithTimeout 超时配置项
//Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *rest.ClientOptions) {
		o.Timeout = timeout
	}
}

//WithNameService 名字服务配置项
//Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
func WithNameService(ns giraffe.NameService) ClientOption {
	return func(o *rest.ClientOptions) {
		o.NameService = ns
	}
}

//WithTracer Tracer配置项
//Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
func WithTracer(tracer *zipkin.Tracer) ClientOption {
	return func(o *rest.ClientOptions) {
		o.Tracer = tracer
	}
}

//WithRoundTripper HTTP配置项
//Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
func WithRoundTripper(rt http.RoundTripper) ClientOption {
	return func(o *rest.ClientOptions) {
		o.Transport = rt
	}
}
