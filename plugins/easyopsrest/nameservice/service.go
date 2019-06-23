package nameservice

import "github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice/internal"

type NameService interface {
	SetAppName(name string)
	GetAllServiceByName(name string) ([]string, error)
	GetServiceByName(name string) (ip string, port int, err error)
	GetEasyopsLicenseInfo() (map[string]string, error)
}

// Deprecated: Use dependency injection
func New() NameService {
	return &service{}
}

type service struct {
}

func (s *service) SetAppName(name string) {
	internal.SetAppName(name)
}

func (s *service) GetAllServiceByName(name string) ([]string, error) {
	return internal.GetAllServiceByName(name)
}

func (s *service) GetServiceByName(name string) (ip string, port int, err error) {
	return internal.GetServiceByName(name)
}

func (s *service) GetEasyopsLicenseInfo() (map[string]string, error) {
	return internal.GetEasyopsLicenseInfo()
}
