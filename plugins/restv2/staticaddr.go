package restv2

import "github.com/easyops-cn/giraffe-micro"

type StaticAddress string

func (s StaticAddress) GetAddress(contract giraffe.Contract) (string, error) {
	return string(s), nil
}

func (s StaticAddress) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{string(s)}, nil
}
