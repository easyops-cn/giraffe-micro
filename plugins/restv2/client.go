package restv2

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/easyops-cn/giraffe-micro"
)

type Middleware interface {
	NewRequest(rule giraffe.HttpRule, in interface{}) (*http.Request, error)
	ParseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error
}

type Client struct {
	*http.Client
	Middleware  Middleware
	NameService giraffe.NameService
}

func (c *Client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}, opts ...giraffe.CallOption) error {
	req, err := c.middleware().NewRequest(md.HttpRule, in)
	if err != nil {
		return err
	}

	resp, err := c.Call(md.Contract, req.WithContext(ctx), opts...)
	if resp != nil {
		defer func() {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}()
	}

	if err != nil {
		return err
	}
	return c.middleware().ParseResponse(md.HttpRule, resp, out)
}

func (c *Client) NewStream(ctx context.Context, sd *giraffe.StreamDesc, opts ...giraffe.CallOption) (giraffe.ClientStream, error) {
	panic("implement me")
}

func (c *Client) middleware() Middleware {
	if c.Middleware != nil {
		return c.Middleware
	}
	return DefaultMiddleware
}

func (c *Client) Call(contract giraffe.Contract, req *http.Request, opts ...giraffe.CallOption) (*http.Response, error) {
	options := &giraffe.CallOptions{
		Metadata: map[string][]string{},
	}
	for _, o := range opts {
		o(options)
	}

	for k, v := range options.Metadata {
		if v != nil && len(v) > 0 {
			req.Header.Set(k, v[0])
		}
	}
	if hostname := req.Header.Get("host"); hostname != "" {
		req.Host = hostname
	}

	if c.NameService != nil {
		addr, err := c.NameService.GetAddress(contract)
		if err != nil {
			return nil, err
		}
		req.URL.Host = addr
	}
	if c.Client == nil {
		return http.DefaultClient.Do(req)
	}
	return c.Do(req)
}

func NewClient(addr string) giraffe.Client {
	return &Client{
		Client:      &http.Client{},
		Middleware:  DefaultMiddleware,
		NameService: StaticAddress(addr),
	}
}
