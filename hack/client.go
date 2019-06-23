package hack

import (
	"context"
	"github.com/easyops-cn/giraffe-micro"
)

type clientWithServiceName struct {
	c           giraffe.Client
	serviceName string
}

type contractWithServiceName struct {
	serviceName string
	giraffe.Contract
}

func (c *contractWithServiceName) Name() string {
	return c.serviceName
}

func (c *contractWithServiceName) Version() string {
	return ""
}

func (c *clientWithServiceName) Invoke(ctx context.Context, contract giraffe.Contract, in interface{}, out interface{}) error {
	return c.c.Invoke(ctx, &contractWithServiceName{c.serviceName, contract}, in, out)
}

func (c *clientWithServiceName) NewStream(ctx context.Context, contract giraffe.Contract) (giraffe.ClientStream, error) {
	return c.c.NewStream(ctx, &contractWithServiceName{c.serviceName, contract})
}

func ClientWithServiceName(serviceName string, c giraffe.Client) *clientWithServiceName {
	return &clientWithServiceName{
		c: c,
		serviceName: serviceName,
	}
}
