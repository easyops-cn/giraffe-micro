package giraffe

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
