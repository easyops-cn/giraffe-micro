package hack

import (
	"context"

	"github.com/easyops-cn/giraffe-micro"
)

type clientWithServiceName struct {
	c           giraffe.Client
	serviceName string
}

type methodWithServiceName struct {
	serviceName string
	giraffe.Method
}

func (c *methodWithServiceName) ServiceName() string {
	return c.serviceName
}

type streamMethodWithServiceName struct {
	serviceName string
	giraffe.StreamMethod
}

func (c *streamMethodWithServiceName) ServiceName() string {
	return c.serviceName
}

func (c *clientWithServiceName) Invoke(ctx context.Context, method giraffe.Method, in interface{}, out interface{}) error {
	return c.c.Invoke(ctx, &methodWithServiceName{c.serviceName, method}, in, out)
}

func (c *clientWithServiceName) NewStream(ctx context.Context, method giraffe.StreamMethod) (giraffe.ClientStream, error) {
	return c.c.NewStream(ctx, &streamMethodWithServiceName{c.serviceName, method})
}

func ClientWithServiceName(serviceName string, c giraffe.Client) *clientWithServiceName {
	return &clientWithServiceName{
		c:           c,
		serviceName: serviceName,
	}
}
