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

type RESTClient interface {
	giraffe.Client
	Init(opt ClientOptions) error
}

type client struct {
	c       *http.Client
	options ClientOptions
}

func (c *client) Init(opt ClientOptions) error {
	c.c.Timeout = opt.Timeout
	c.options = opt

	if opt.ContractService == nil {
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

func (c *client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}) error {
	request, err := newRequest(md, in)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	addr, err := c.options.ContractService.GetAddress(md.Contract)
	if err != nil {
		return err
	}
	request.URL.Host = addr
	request.URL.Scheme = "http"

	response, err := c.c.Do(request)
	if err != nil {
		return err
	}
	if err := parseResponse(md, response, out); err != nil {
		return err
	}
	return nil
}

func (c *client) NewStream(ctx context.Context, sd *giraffe.StreamDesc) (giraffe.ClientStream, error) {
	return nil, errors.New("not supported")
}

func NewClient(opts ...ClientOption) (RESTClient, error) {
	opt := newClientOptions(opts...)

	c := &client{
		c: &http.Client{},
	}

	if err := c.Init(opt); err != nil {
		return nil, err
	}

	return c, nil
}
