package restv2

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/easyops-cn/giraffe-micro"
)

//Middleware 中间件定义
type Middleware interface {
	NewRequest(rule giraffe.HttpRule, in interface{}) (*http.Request, error)
	ParseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error
}

//Client REST Client对象
type Client struct {
	*http.Client
	Middleware  Middleware
	NameService giraffe.NameService
}

//Invoke 单次请求方法
func (c *Client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}, opts ...giraffe.CallOption) error {
	req, err := c.middleware().NewRequest(md.HttpRule, in)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	resp, err := c.Call(md.Contract, req, opts...)
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

//NewStream 流式请求方法(未实现)
func (c *Client) NewStream(ctx context.Context, sd *giraffe.StreamDesc, opts ...giraffe.CallOption) (giraffe.ClientStream, error) {
	return nil, errors.New("not supported")
}

func (c *Client) middleware() Middleware {
	if c.Middleware != nil {
		return c.Middleware
	}
	return DefaultMiddleware
}

func (c *Client) httpClient() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

//Call 请求函数
func (c *Client) Call(contract giraffe.Contract, req *http.Request, opts ...giraffe.CallOption) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("request was nil")
	}

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
		req.URL.Scheme = "http"
	}

	return c.httpClient().Do(req)
}

//NewClient Client实例化函数
func NewClient(addr string) giraffe.Client {
	return &Client{
		Client:      &http.Client{},
		Middleware:  DefaultMiddleware,
		NameService: StaticAddress(addr),
	}
}
