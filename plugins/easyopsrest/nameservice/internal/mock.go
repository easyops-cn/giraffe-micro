// +build mock windows darwin

// mock 数据请自行从 go.easyops.local/nameservice/internal/mock.json 获取，
// 将其中的 data 数据保存文件，并设置环境变量 `EASYOPS_NAME_SERVICE_MOCK` 为该文件路径。
// mock license文件信息 请设置环境变量 `EASYOPS_LICENSE_MOCK` 为license文件路径。

package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	envMockJson    = "EASYOPS_NAME_SERVICE_MOCK"
	envMockLicense = "EASYOPS_LICENSE_MOCK"
)

func SetAppName(name string) {
}

func GetServiceByName(name string) (ip string, port int, err error) {
	items := getItems()
	for _, it := range items {
		if it.ServiceName == name {
			return it.Hosts[0].IP, it.Hosts[0].Port, nil
		}
	}
	return "", 0, errors.New("name not found: " + name)
}

func GetAllServiceByName(name string) ([]string, error) {
	items := getItems()

	toAddr := func(hosts []host) []string {
		vsm := make([]string, len(hosts))
		for i, h := range hosts {
			vsm[i] = fmt.Sprintf("%v:%v", h.IP, h.Port)
		}
		return vsm
	}

	for _, it := range items {
		if it.ServiceName == name {
			return toAddr(it.Hosts), nil
		}
	}

	return nil, errors.New("name not found: " + name)
}

func GetEasyopsLicenseInfo() (map[string]string, error) {
	fname := os.Getenv(envMockLicense)
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Panic(err)
	}
	lines := strings.Split(string(data), "\n")
	m := make(map[string]string, len(lines))
	for _, line := range lines {
		var key, value string
		_, err := fmt.Sscanf(line, "%s = %s", &key, &value)
		if err != nil {
			continue
		}
		m[key] = value
	}
	return m, nil
}

func getItems() []item {
	fname := os.Getenv(envMockJson)
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Panic(err)
	}

	var c []item
	json.Unmarshal(raw, &c)
	return c
}

type item struct {
	Ctime       int    `json:"ctime"`
	ServiceName string `json:"service_name"`
	Hosts       []host `json:"hosts"`
	Mtime       int    `json:"mtime"`
	ID          string `json:"id"`
}

type host struct {
	IP     string   `json:"ip"`
	Tag    []string `json:"tag"`
	Port   int      `json:"port"`
	Weight int      `json:"weight"`
}
