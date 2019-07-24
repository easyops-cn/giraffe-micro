package ens

import (
	"errors"
	"fmt"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice"
)

type contractService struct {
	ns nameservice.NameService
}

func serviceName(contract giraffe.Contract) (string, error) {
	switch {
	case contract.GetName() == "":
		return "", errors.New("contract name was empty")
	case contract.GetVersion() == "":
		return contract.GetName(), nil
	default:
		return fmt.Sprintf("%s@%s", contract.GetName(), contract.GetVersion()), nil
	}
}

func (c *contractService) GetAddress(contract giraffe.Contract) (string, error) {
	name, err := serviceName(contract)
	if err != nil {
		return "", err
	}
	ip, port, err := c.ns.GetServiceByName(name)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", ip, port), nil
}

func (c *contractService) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	name, err := serviceName(contract)
	if err != nil {
		return nil, err
	}
	return c.ns.GetAllServiceByName(name)
}

func NewNameService() giraffe.NameService {
	return &contractService{
		nameservice.New(),
	}
}
