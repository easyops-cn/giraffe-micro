package hack

import (
	"context"

	"github.com/easyops-cn/giraffe-micro"
)

type clientWithServiceName struct {
	c           giraffe.Client
	contract    giraffe.Contract
}

type contractWithServiceName struct {
	serviceName string
}

func (c *contractWithServiceName) ContractName() string {
	return c.serviceName
}

func (c *contractWithServiceName) ContractVersion() string {
	return ""
}

type unaryMethod struct {
	giraffe.Contract
	giraffe.Method
}

type httpMethod struct {
	giraffe.Contract
	giraffe.Method
	giraffe.HttpRule
}

type streamMethod struct {
	giraffe.Contract
	giraffe.StreamMethod
}

func (c *clientWithServiceName) Invoke(ctx context.Context, method giraffe.Method, in interface{}, out interface{}) error {
	if httpRule, ok := method.(giraffe.HttpRule); ok {
		return c.c.Invoke(ctx, &httpMethod{c.contract, method, httpRule}, in, out)
	}
	return c.c.Invoke(ctx, &unaryMethod{c.contract, method}, in, out)
}

func (c *clientWithServiceName) NewStream(ctx context.Context, method giraffe.StreamMethod) (giraffe.ClientStream, error) {
	return c.c.NewStream(ctx, &streamMethod{c.contract, method})
}

func ClientWithServiceName(serviceName string, c giraffe.Client) giraffe.Client {
	return &clientWithServiceName{
		c:           c,
		contract: &contractWithServiceName{serviceName:serviceName},
	}
}
