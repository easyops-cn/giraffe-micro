```bash
go get go.easyops.local/nameservice
dep ensure -add go.easyops.local/nameservice
```
# mock用法
名字服务可以手动指向某个IP或者端口， 配置文件在${PROJECT_PATH}/vendor/go.easyops.local/nameservice/internal/mock.json
启动程序前export一下`EASYOPS_NAME_SERVICE_MOCK`这个环境变量即可
```bash
export EASYOPS_NAME_SERVICE_MOCK=${PROJECT_PATH}/vendor/go.easyops.local/nameservice/internal/mock.json
```
license信息mock
启动程序前export一下`EASYOPS_LICENSE_MOCK`这个环境变量即可
```bash
export EASYOPS_LICENSE_MOCK=${PROJECT_PATH}/vendor/go.easyops.local/nameservice/internal/mock_license.lic
```

mac和linux用户还需要指定动态库路径
```bash
export DYLD_LIBRARY_PATH=${PROJECT_PATH}/vendor/go.easyops.local/nameservice/internal/
```