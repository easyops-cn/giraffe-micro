package giraffe

import (
	"context"

	"github.com/easyops-cn/giraffe-micro/codes"
	"github.com/easyops-cn/giraffe-micro/metadata"
)

//UnaryEndpoint 单次调用函数
type UnaryEndpoint func(ctx context.Context, req interface{}) (interface{}, error)

//StreamEndpoint 流式调用函数
type StreamEndpoint func(ctx context.Context, stream ServiceStream) error

//CallOption 请求配置函数
type CallOption func(o *CallOptions)

//Server 服务端接口
type Server interface {
	RegisterUnaryEndpoint(md *MethodDesc, handle UnaryEndpoint)
	RegisterStreamEndpoint(sd *StreamDesc, handle StreamEndpoint)
}

//Client 客户端接口
type Client interface {
	Invoke(ctx context.Context, md *MethodDesc, in interface{}, out interface{}, opts ...CallOption) error
	NewStream(ctx context.Context, sd *StreamDesc, opts ...CallOption) (ClientStream, error)
}

//ClientStream 流式客户端接口
type ClientStream interface {
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
	CloseSend() error
}

//ServiceStream 流式服务接口
type ServiceStream interface {
	// TODO add support SetHeader() SendHeader() SetTrailer()
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

//NameService 名字服务接口
type NameService interface {
	GetAddress(ctx context.Context, contract Contract) (string, error)
	GetAllAddresses(ctx context.Context, contract Contract) ([]string, error)
}

//Contract 契约定义接口
type Contract interface {
	GetName() string
	GetVersion() string
}

//HttpRule HTTP规则定义接口
type HttpRule interface {
	GetGet() string
	GetPut() string
	GetPost() string
	GetDelete() string
	GetPatch() string
	GetBody() string
	GetResponseBody() string
}

//StatusCode 统一状态码接口
type StatusCode interface {
	StatusCode() codes.Code
}

//MethodDesc 单次调用方法定义
type MethodDesc struct {
	Contract     Contract
	ServiceName  string
	MethodName   string
	RequestType  interface{}
	ResponseType interface{}
	HttpRule     HttpRule
}

//StreamDesc 流式调用方法定义
type StreamDesc struct {
	Contract      Contract
	ServiceName   string
	StreamName    string
	ClientStreams bool
	ServerStreams bool
}

//CallOptions 请求配置
type CallOptions struct {
	Metadata metadata.MD
}
