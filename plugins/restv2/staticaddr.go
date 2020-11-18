package restv2

import (
	"context"

	"github.com/easyops-cn/giraffe-micro"
)

type StaticAddress string

func (s StaticAddress) GetAddress(ctx context.Context, contract giraffe.Contract) (string, error) {
	return string(s), nil
}

func (s StaticAddress) GetAllAddresses(ctx context.Context, contract giraffe.Contract) ([]string, error) {
	return []string{string(s)}, nil
}
