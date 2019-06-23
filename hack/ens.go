package hack

import (
	"github.com/easyops-cn/giraffe-micro"
)

type staticNameService struct {
	Host string
	Port int
}

func (n *staticNameService) GetAddress(contract giraffe.Contract) (*giraffe.Address, error) {
	name := contract.Name()
	if contract.Version() != "" {
		name  = name + "@" + contract.Version()
	}
	return &giraffe.Address{
		Host: n.Host,
		Port: n.Port,
		Name: name,
	}, nil
}

func (n *staticNameService) GetAllAddresses(contract giraffe.Contract) ([]giraffe.Address, error) {
	name := contract.Name()
	if contract.Version() != "" {
		name  = name + "@" + contract.Version()
	}
	return []giraffe.Address{
		{
			Host: n.Host,
			Port: n.Port,
			Name: name,
		},
	}, nil
}

func StaticAddress(host string, port int) giraffe.NameService {
	return &staticNameService{
		Host: host,
		Port: port,
	}
}
