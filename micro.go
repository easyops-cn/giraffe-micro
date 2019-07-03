package giraffe

import (
	"context"

	"github.com/easyops-cn/giraffe-micro/status"
)

type UnaryEndpoint func(ctx context.Context, req interface{}) (interface{}, error)

type StreamEndpoint func(ctx context.Context, stream ServiceStream) error

type Server interface {
	RegisterUnaryEndpoint(method Method, handle UnaryEndpoint)
	RegisterStreamEndpoint(method StreamMethod, handle StreamEndpoint)
}

type Client interface {
	Invoke(ctx context.Context, method Method, in interface{}, out interface{}) error
	NewStream(ctx context.Context, method StreamMethod) (ClientStream, error)
}

type Method interface {
	ServiceName() string
	MethodName() string
	RequestMessage() interface{}  // for gRPC handler transform
	ResponseMessage() interface{} // for gRPC handler transform
}

type StreamMethod interface {
	Method
	ClientStreams() bool
	ServerStreams() bool
}

type Contract interface {
	ContractName() string
	ContractVersion() string
}

type HttpRule interface {
	Pattern() (string, string)
	Body() string
}

type Error interface {
	Code() status.Code
	Error() string
	Message() string
	WithMessage(message string) Error
	Proto() *status.Status
}
