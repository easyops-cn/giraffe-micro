package easyopsrest

import (
	"strconv"
	"strings"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice"
)

var DefaultNS = &ens{}

type ens struct{}

func serviceName(method giraffe.Method) string {
	if contract, ok := method.(giraffe.Contract); ok && contract.ContractName() != "" {
		a := []string{contract.ContractName()}
		if contract.ContractVersion() != "" {
			a = append(a, contract.ContractVersion())
		}
		return strings.Join(a, "@")
	}
	return method.ServiceName()
}

func (e *ens) GetAddress(method giraffe.Method) (*giraffe.Address, error) {
	name := serviceName(method)
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

func (e *ens) GetAllAddresses(method giraffe.Method) ([]giraffe.Address, error) {
	name := serviceName(method)
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
