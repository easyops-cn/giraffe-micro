package giraffe

import "strings"

//methodType Method类型的数据类型
//私有类型, 确保类型合法
type methodType string

//Method类型常量表
const (
	MethodtypeUnary               methodType = "Unary"
	MethodtypeBidirectionalStream methodType = "BidirectionalStream"
	MethodtypeServerStream        methodType = "ServerStream"
	MethodtypeClientStream        methodType = "ClientStream"
)

//MethodDesc 接口定义
type MethodDesc struct {
	ServiceName  string                 // 服务名, 对应服务定义的名字, 主要是 gRPC 需要用到该字段作为 URL
	MethodName   string                 // 方法名, 对应接口方法的名字, 主要是 gRPC 需要用到该字段作为 URL
	Type         methodType             // 方法类型, 分为“单次调用”、“客户端流式调用”、“服务端流式调用”、“双向流式调用”
	RequestType  interface{}            // 请求信息的nil指针
	ResponseType interface{}            // 响应信息的nil指针
	MetaData     map[string]interface{} // 其他元数据
}

func (md *MethodDesc) Name() string {
	if v, ok := md.MetaData["contract_name"]; ok {
		return v.(string)
	}
	dotIndex := strings.LastIndex(md.ServiceName, ".")
	return md.ServiceName[:dotIndex]
}

func (md *MethodDesc) Version() string {
	if v, ok := md.MetaData["contract_version"]; ok {
		return v.(string)
	}
	return ""
}

func (md *MethodDesc) Desc() *MethodDesc {
	return md
}
