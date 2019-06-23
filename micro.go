package giraffe

import (
	"context"
	"github.com/easyops-cn/giraffe-micro/status"
)

//UnaryEndpoint 单次调用的接口函数定义
type UnaryEndpoint func(ctx context.Context, req interface{}) (interface{}, error)

//StreamEndpoint 流式调用的接口函数定义
type StreamEndpoint func(ctx context.Context, stream ServiceStream) error

//Server 服务器接口, 提供method与stream接口的注册
type Server interface {
	RegisterUnaryEndpoint(contract Contract, endpoint UnaryEndpoint)
	RegisterStreamEndpoint(contract Contract, endpoint StreamEndpoint)
}

//Client 客户端接口
//提供Invoke单次调用方法和NewStream流式调用方法
type Client interface {
	Invoke(ctx context.Context, contract Contract, in interface{}, out interface{}) error
	NewStream(ctx context.Context, contract Contract) (ClientStream, error)
}

type Contract interface {
	Name() string
	Version() string
	Desc() *MethodDesc
}

type Error interface {
	Code() status.Code
	Error() string
	Message() string
	WithMessage(message string) Error
	Proto() *status.Status
}
