package hack

import (
	"github.com/easyops-cn/giraffe-micro"
)

type staticNameService struct {
	Host string
	Port int
}

func (n *staticNameService) GetAddress(method giraffe.Method) (*giraffe.Address, error) {
	return &giraffe.Address{
		Host: n.Host,
		Port: n.Port,
		Name: method.ServiceName(),
	}, nil
}

func (n *staticNameService) GetAllAddresses(method giraffe.Method) ([]giraffe.Address, error) {
	return []giraffe.Address{
		{
			Host: n.Host,
			Port: n.Port,
			Name: method.ServiceName(),
		},
	}, nil
}

func StaticAddress(host string, port int) giraffe.NameService {
	return &staticNameService{
		Host: host,
		Port: port,
	}
}
