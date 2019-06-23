package nameservice

import "sync"

var (
	globalMu sync.RWMutex
	globalNS = New()
)

func ReplaceGlobalNS(service NameService) func() {
	globalMu.Lock()
	prev := globalNS
	globalNS = service
	globalMu.Unlock()
	return func() { ReplaceGlobalNS(prev) }
}

func NS() NameService {
	globalMu.RLock()
	ns := globalNS
	globalMu.RUnlock()
	return ns
}

func SetAppName(name string) {
	NS().SetAppName(name)
}

func GetAllServiceByName(name string) ([]string, error) {
	return NS().GetAllServiceByName(name)
}

func GetServiceByName(name string) (ip string, port int, err error) {
	return NS().GetServiceByName(name)
}

func GetEasyopsLicenseInfo() (map[string]string, error) {
	return NS().GetEasyopsLicenseInfo()
}
