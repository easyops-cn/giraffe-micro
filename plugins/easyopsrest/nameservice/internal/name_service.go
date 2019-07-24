// +build linux
// +build !mock
// +build !travis

package internal

import (
	"errors"
	"fmt"
	"unsafe"
)

// #cgo amd64 darwin LDFLAGS: -L. -lens_sdk
// #include <stdlib.h>
// #include "ens_sdk.h"
import "C"

const (
	IP_LEN              = 16
	MAX_IP_NUM          = 8
	MAX_LICENSE_KEY_NUM = 16
	LICENSE_KEY_LEN     = 32
	LICENSE_VAL_LEN     = 512
)

var appName string

func SetAppName(name string) {
	appName = name
}

func cArrayToGoArray(cArray unsafe.Pointer, size int) (goArray []byte) {
	p := uintptr(cArray)
	for i := 0; i < size; i++ {
		j := *(*byte)(unsafe.Pointer(p))
		if j != 0 {
			goArray = append(goArray, j)
		}
		p += unsafe.Sizeof(j)
	}
	return
}

func GetServiceByName(dstName string) (ip string, port int, err error) {
	srcNameC := C.CString(appName)
	defer C.free(unsafe.Pointer(srcNameC))
	dstNameC := C.CString(dstName)
	defer C.free(unsafe.Pointer(dstNameC))

	var c_port C.int
	var c_ip [IP_LEN]C.char
	ret := C.get_service_by_name(srcNameC, dstNameC, (*C.char)(&c_ip[0]), IP_LEN, &c_port)
	if ret < 0 {
		return "", 0, errors.New("get name service error,name:" + dstName)
	}

	ip = string(cArrayToGoArray(unsafe.Pointer(&c_ip[0]), IP_LEN))
	port = int(c_port)
	return ip, port, nil
}

func GetAllServiceByName(dstName string) ([]string, error) {
	srcNameC := C.CString(appName)
	defer C.free(unsafe.Pointer(srcNameC))
	dstNameC := C.CString(dstName)
	defer C.free(unsafe.Pointer(dstNameC))

	var portArray [MAX_IP_NUM]C.int
	var ipArray [MAX_IP_NUM][IP_LEN]C.char
	var arrSize C.int = MAX_IP_NUM
	ret := C.get_multi_service_by_name(srcNameC, dstNameC, (*[IP_LEN]C.char)(&ipArray[0]), (*C.int)(&portArray[0]), &arrSize)

	if ret < 0 {
		return []string{}, errors.New("get all name service error,name:" + dstName)
	}

	if arrSize <= 0 || arrSize > MAX_IP_NUM {
		return []string{}, errors.New(fmt.Sprintf("invalid arrSize = %d", arrSize))
	}

	var ipList []string
	for i := 0; i < int(arrSize); i++ {
		// ip
		ipValue := ipArray[i]
		ip := string(cArrayToGoArray(unsafe.Pointer(&ipValue[0]), IP_LEN))

		// port
		portValue := portArray[i]
		port := int(portValue)

		ipList = append(ipList, fmt.Sprintf("%s:%d", ip, port))
	}
	return ipList, nil
}

func GetEasyopsLicenseInfo() (map[string]string, error) {
	var keyArray [MAX_LICENSE_KEY_NUM][LICENSE_KEY_LEN]C.char
	var valueArray [MAX_LICENSE_KEY_NUM][LICENSE_VAL_LEN]C.char
	var arrSize C.int = MAX_LICENSE_KEY_NUM
	num := C.get_easyops_license_info((*[LICENSE_KEY_LEN]C.char)(&keyArray[0]), (*[LICENSE_VAL_LEN]C.char)(&valueArray[0]), &arrSize)
	if num < 0 {
		return nil, errors.New(fmt.Sprintf("get license info error:%d", int(num)))
	}
	keyNum := int(arrSize)
	result := make(map[string]string, keyNum)
	for i := 0; i < keyNum; i++ {
		keyValue := keyArray[i]
		key := string(cArrayToGoArray(unsafe.Pointer(&keyValue[0]), LICENSE_KEY_LEN))

		valueValue := valueArray[i]
		value := string(cArrayToGoArray(unsafe.Pointer(&valueValue[0]), LICENSE_VAL_LEN))

		result[key] = value
	}
	return result, nil
}
