package easyopsrest

import (
	"context"
	"errors"
	"net/http"

	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/gerr"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/auth"
)

type client struct {
	c       *http.Client
	options ClientOptions
}

func (c *client) Invoke(ctx context.Context, method giraffe.Method, in interface{}, out interface{}) error {
	request, err := NewRequest(method, in)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	addr, err := c.options.nameService.GetAddress(method)
	if err != nil {
		return err
	}
	request.URL.Host = addr.String()
	request.URL.Scheme = "http"

	response, err := c.c.Do(request)
	if err != nil {
		return gerr.UnknownErrorf("unexpected error %s when doing request", err.Error())
	}
	if err := parseResponse(response, out); err != nil {
		return err
	}
	return nil
}

func (c *client) NewStream(ctx context.Context, method giraffe.StreamMethod) (giraffe.ClientStream, error) {
	return nil, errors.New("not supported")
}

func NewClient(opts ...ClientOption) (giraffe.Client, error) {
	opt := newClientOptions(opts...)

	c := &client{
		c: &http.Client{
			Timeout: opt.timeout,
		},
		options: opt,
	}

	rt := opt.rt
	if opt.tracer != nil {
		t, err := zipkinhttp.NewTransport(opt.tracer, zipkinhttp.RoundTripper(opt.rt))
		if err != nil {
			return nil, err
		}
		rt = t
	}
	c.c.Transport = auth.NewTransport(auth.RoundTripper(rt))

	return c, nil
}
