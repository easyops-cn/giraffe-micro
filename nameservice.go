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

//NameService
//名字服务接口, 按照接口定义路由
type NameService interface {
	GetAddress(contract Contract) (*Address, error)
	GetAllAddresses(contract Contract) ([]Address, error)
}
