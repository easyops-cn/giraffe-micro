package giraffe

import (
	"context"

	"github.com/easyops-cn/giraffe-micro/codes"
)

type UnaryEndpoint func(ctx context.Context, req interface{}) (interface{}, error)

type StreamEndpoint func(ctx context.Context, stream ServiceStream) error

type Server interface {
	RegisterUnaryEndpoint(md *MethodDesc, handle UnaryEndpoint)
	RegisterStreamEndpoint(sd *StreamDesc, handle StreamEndpoint)
}

type Client interface {
	Invoke(ctx context.Context, md *MethodDesc, in interface{}, out interface{}) error
	NewStream(ctx context.Context, sd *StreamDesc) (ClientStream, error)
}

type ClientStream interface {
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
	CloseSend() error
}

type ServiceStream interface {
	// TODO add support SetHeader() SendHeader() SetTrailer()
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type ContractService interface {
	GetAddress(contract Contract) (string, error)
	GetAllAddresses(contract Contract) ([]string, error)
}

type Contract interface {
	GetName() string
	GetVersion() string
}

type HttpRule interface {
	GetGet() string
	GetPut() string
	GetPost() string
	GetDelete() string
	GetPatch() string
	GetBody() string
	GetResponseBody() string
}

type StatusCode interface {
	GiraffeStatusCode() codes.Code
}

type MethodDesc struct {
	Contract     Contract
	ServiceName  string
	MethodName   string
	RequestType  interface{}
	ResponseType interface{}
	HttpRule     HttpRule
}

type StreamDesc struct {
	Contract      Contract
	ServiceName   string
	StreamName    string
	ClientStreams bool
	ServerStreams bool
}
