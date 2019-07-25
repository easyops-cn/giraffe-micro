package hack

import (
	"fmt"

	"github.com/easyops-cn/giraffe-micro"
)

type ens struct {
	addr string
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (e *ens) GetAddress(contract giraffe.Contract) (string, error) {
	return e.addr, nil
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (e *ens) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{e.addr}, nil
}

// Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func StaticAddress(host string, port int) giraffe.NameService {
	return &ens{addr: fmt.Sprintf("%s:%d", host, port)}
}
