package hack

import (
	"context"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"

	"github.com/easyops-cn/giraffe-micro"
)

type client struct {
	c        giraffe.Client
	contract giraffe.Contract
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (c *client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}) error {
	t := *md // copy MethodDesc
	t.Contract = c.contract
	return c.c.Invoke(ctx, &t, in, out)
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (c *client) NewStream(ctx context.Context, sd *giraffe.StreamDesc) (giraffe.ClientStream, error) {
	t := *sd // copy StreamDesc
	t.Contract = c.contract
	return c.c.NewStream(ctx, &t)
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func ClientWithServiceName(serviceName string, c giraffe.Client) giraffe.Client {
	return &client{
		contract: &giraffeproto.Contract{
			Name:    serviceName,
			Version: "",
		},
		c: c,
	}
}
