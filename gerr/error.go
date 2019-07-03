package gerr

import (
	"fmt"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/status"
)

type _error struct {
	s *status.Status
}

func (e *_error) Proto() *status.Status {
	return e.s
}

func (e *_error) Code() status.Code {
	return e.s.GetCode()
}

func (e *_error) Message() string {
	return e.s.GetMessage()
}

func (e *_error) WithMessage(message string) giraffe.Error {
	e.s.Message = message
	return e
}

func (e *_error) Error() string {
	return e.s.GetError()
}

func FromProto(s *status.Status) giraffe.Error {
	if s.Code == status.Code_OK {
		return nil
	}
	return &_error{s: s}
}

func FromError(err error) giraffe.Error {
	if e, ok := err.(giraffe.Error); ok {
		return e
	}
	return UnknownErrorf("%s", err.Error())
}

func newErrorf(code status.Code) func(format string, a ...interface{}) giraffe.Error {
	return func(format string, a ...interface{}) giraffe.Error {
		return &_error{
			s: &status.Status{
				Code:  status.Code(code),
				Error: fmt.Sprintf("%s %s", status.Code_name[int32(code)], fmt.Sprintf(format, a...)),
			},
		}
	}
}

var (
	CancelledErrorf          = newErrorf(status.Code_CANCELLED)
	UnknownErrorf            = newErrorf(status.Code_UNKNOWN)
	InvalidArgumentErrorf    = newErrorf(status.Code_INVALID_ARGUMENT)
	FailedPreconditionErrorf = newErrorf(status.Code_FAILED_PRECONDITION)
	OutOfRangeErrorf         = newErrorf(status.Code_OUT_OF_RANGE)
	UnauthenticatedErrorf    = newErrorf(status.Code_UNAUTHENTICATED)
	PermissionDeniederrorf   = newErrorf(status.Code_PERMISSION_DENIED)
	NotFoundErrorf           = newErrorf(status.Code_NOT_FOUND)
	AbortedErrorf            = newErrorf(status.Code_ABORTED)
	AlreadyExistsErrorf      = newErrorf(status.Code_ALREADY_EXISTS)
	ResourceExhaustedErrorf  = newErrorf(status.Code_RESOURCE_EXHAUSTED)
	DataLossErrorf           = newErrorf(status.Code_DATA_LOSS)
	InternalErrorf           = newErrorf(status.Code_INTERNAL)
	UnimplementedErrorf      = newErrorf(status.Code_UNIMPLEMENTED)
	UnavailableErrorf        = newErrorf(status.Code_UNAVAILABLE)
	DeadlineExceededErrorf   = newErrorf(status.Code_DEADLINE_EXCEEDED)
)
