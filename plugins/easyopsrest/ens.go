package easyopsrest

import (
	"strconv"
	"strings"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice"
)

var DefaultNS = &ens{}

type ens struct{}

func serviceName(contract giraffe.Contract) string {
	serviceName := contract.ContractName()
	if contract.ContractVersion() != "" {
		serviceName += contract.ContractVersion()
	}
	return serviceName
}

func (e *ens) GetAddress(contract giraffe.Contract) (*giraffe.Address, error) {
	name := serviceName(contract)
	ip, port, err := nameservice.GetServiceByName(name)
	if err != nil {
		return nil, err
	}
	return &giraffe.Address{
		Host: ip,
		Port: port,
		Name: name,
	}, nil
}

func (e *ens) GetAllAddresses(contract giraffe.Contract) ([]giraffe.Address, error) {
	name := serviceName(contract)
	strs, err := nameservice.GetAllServiceByName(name)
	if err != nil {
		return nil, err
	}
	addrs := make([]giraffe.Address, len(strs))
	for i, s := range strs {
		colonIndex := strings.LastIndex(s, ":")
		addrs[i] = giraffe.Address{
			Host: s[:colonIndex],
			Port: func() int {
				if v, err := strconv.Atoi(s[colonIndex+1:]); err == nil {
					return v
				}
				return 80
			}(),
			Name: name,
		}
	}

	return addrs, nil
}
