package giraffe

import "fmt"

type Address struct {
	Host string
	Port int
	Name string
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type NameService interface {
	GetAddress(method Method) (*Address, error)
	GetAllAddresses(method Method) ([]Address, error)
}
