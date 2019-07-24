package rest

import (
	"context"
	"errors"
	"net/http"

	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/easyops-cn/giraffe-micro"
)

// for unit testing convenient
var zipkinTransportFactory = zipkinhttp.NewTransport

type Client interface {
	giraffe.Client
	NewRequest(md *giraffe.MethodDesc, in interface{}) (*http.Request, error)
	Call(ctx context.Context, md *giraffe.MethodDesc, req *http.Request, out interface{}) error
}

type client struct {
	c       *http.Client
	options ClientOptions
}

func (c *client) init(opt ClientOptions) error {
	c.c.Timeout = opt.Timeout
	c.options = opt

	if opt.NameService == nil {
		return errors.New("nameservice should not be nil")
	}

	rt := opt.Transport
	if opt.Tracer != nil {
		t, err := zipkinTransportFactory(opt.Tracer, zipkinhttp.RoundTripper(opt.Transport))
		if err != nil {
			return err
		}
		rt = t
	}
	c.c.Transport = rt

	return nil
}

func (c *client) NewRequest(md *giraffe.MethodDesc, in interface{}) (*http.Request, error) {
	request, err := newRequest(md, in)
	if err != nil {
		return nil, err
	}
	addr, err := c.options.NameService.GetAddress(md.Contract)
	if err != nil {
		return nil, err
	}
	request.URL.Host = addr
	request.URL.Scheme = "http"

	return request, nil
}

func (c *client) Call(ctx context.Context, md *giraffe.MethodDesc, req *http.Request, out interface{}) error {
	request := req.WithContext(ctx)
	response, err := c.c.Do(request)
	if err != nil {
		return err
	}
	if err := parseResponse(md, response, out); err != nil {
		return err
	}
	return nil
}

func (c *client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}) error {
	req, err := c.NewRequest(md, in)
	if err != nil {
		return err
	}
	return c.Call(ctx, md, req, out)
}

func (c *client) NewStream(ctx context.Context, sd *giraffe.StreamDesc) (giraffe.ClientStream, error) {
	return nil, errors.New("not supported")
}

func NewClient(opts ...ClientOption) (Client, error) {
	opt := newClientOptions(opts...)

	c := &client{
		c: &http.Client{},
	}

	if err := c.init(opt); err != nil {
		return nil, err
	}

	return c, nil
}
