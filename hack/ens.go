package hack

import (
	"fmt"

	"github.com/easyops-cn/giraffe-micro"
)

type ens struct {
	addr string
}

//GetAddress 根据契约查询地址
//Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (e *ens) GetAddress(contract giraffe.Contract) (string, error) {
	return e.addr, nil
}

//GetAllAddresses 根据契约查询所有地址
//Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func (e *ens) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{e.addr}, nil
}

//StaticAddress 指定地址的名字服务, 使用该名字服务的契约均按照指定地址路由, 一般用于测试
//Deprecated: Replace by go.easyops.local/giraffe-micro/pkg/hack
func StaticAddress(host string, port int) giraffe.NameService {
	return &ens{addr: fmt.Sprintf("%s:%d", host, port)}
}
