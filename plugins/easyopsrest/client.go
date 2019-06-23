package easyopsrest

import (
	"context"
	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/gerr"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsctx"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"net/http"
	"strconv"
	"time"
)

type ClientOption func(c *client)

type client struct {
	c *http.Client
	nameService giraffe.NameService
}

func (c *client) Invoke(ctx context.Context, contract giraffe.Contract, in interface{}, out interface{}) error {
	request, err := NewRequest(contract, in)
	if err != nil {
		return err
	}

	addr, err := c.nameService.GetAddress(contract)
	if err != nil {
		return err
	}
	request.URL.Host = addr.String()
	request.URL.Scheme = "http"
	userInfo := easyopsctx.FromContext(ctx)
	if userInfo.User != "" {
		request.Header.Add("user", userInfo.User)
	}
	if userInfo.Org != 0 {
		request.Header.Add("org", strconv.Itoa(userInfo.Org))
	}

	response, err := c.c.Do(request)
	if err != nil {
		return gerr.UnknownErrorf("unexpected error %#v when doing request", err)
	}
	if err := parseResponse(response, out); err != nil {
		return err
	}
	return nil
}

func (c *client) NewStream(ctx context.Context, contract giraffe.Contract) (giraffe.ClientStream, error) {
	panic("Not Supported.")
}

func NewClient(opt ...ClientOption) *client {

	c := &client{
		c: &http.Client{
			Timeout: time.Second * 60,
		},
		nameService: &ens{},
	}

	for _, o := range opt {
		o(c)
	}

	return c
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *client) {
		c.c.Timeout = timeout
	}
}

func WithNameService(ns giraffe.NameService) ClientOption {
	return func(c *client) {
		c.nameService = ns
	}
}

func WithTracer(tracer *zipkin.Tracer) ClientOption {
	return func(c *client) {
		tr, err := zipkinhttp.NewTransport(tracer)
		if err != nil { return }
		c.c.Transport = tr
	}
}

func WithTransport(tr http.RoundTripper) ClientOption {
	return func(c *client) {
		c.c.Transport = tr
	}
}
