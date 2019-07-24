package rest

import (
	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/auth"
	"github.com/easyops-cn/giraffe-micro/plugins/rest"
)

var restClientFactory = rest.NewClient

// Deprecated: Replace by go.easyops.local/giraffe-micro/v2/rest
func NewClient(opts ...ClientOption) (giraffe.Client, error) {
	opt := newClientOptions(opts...)

	client, err := restClientFactory(
		rest.WithRoundTripper(auth.NewTransport(opt.Transport)),
		rest.WithTracer(opt.Tracer),
		rest.WithNameService(opt.NameService),
		rest.WithTimeout(opt.Timeout),
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}
