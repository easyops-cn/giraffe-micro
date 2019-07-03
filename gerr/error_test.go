package gerr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/status"
)

func Test_newErrorf(t *testing.T) {
	type args struct {
		code status.Code
	}
	tests := []struct {
		name string
		args args
		want giraffe.Error
	}{
		{
			name: "Test_happy_path",
			args: args{
				code: 0,
			},
			want: &_error{
				s: &status.Status{
					Code:  status.Code(0),
					Error: "OK ",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newErrorf(tt.args.code)(""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newErrorf()(\"\") = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want giraffe.Error
	}{
		{
			name: "TestWithError",
			args: args{
				err: InvalidArgumentErrorf(""),
			},
			want: &_error{
				s: &status.Status{
					Code:  100000,
					Error: "INVALID_ARGUMENT ",
				},
			},
		},
		{
			name: "TestUnknownError",
			args: args{
				err: fmt.Errorf("unknown"),
			},
			want: &_error{
				s: &status.Status{
					Code:  2,
					Error: "UNKNOWN unknown",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromProto(t *testing.T) {
	type args struct {
		s *status.Status
	}
	tests := []struct {
		name string
		args args
		want giraffe.Error
	}{
		{
			name: "TestOK",
			args: args{
				s: &status.Status{
					Code: 0,
				},
			},
			want: nil,
		},
		{
			name: "TestCodeNotZero",
			args: args{
				s: &status.Status{
					Code: 10,
				},
			},
			want: &_error{
				s: &status.Status{
					Code: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromProto(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_error_Proto(t *testing.T) {
	pb := &status.Status{}

	type fields struct {
		s *status.Status
	}
	tests := []struct {
		name   string
		fields fields
		want   *status.Status
	}{
		{
			name: "TestHappyPath",
			fields: fields{
				s: pb,
			},
			want: pb,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &_error{
				s: tt.fields.s,
			}
			if got := e.Proto(); !reflect.DeepEqual(&got, &tt.want) {
				t.Errorf("_error.Proto() = %v, want %v", &got, &tt.want)
			}
		})
	}
}

func Test_error(t *testing.T) {
	type fields struct {
		s *status.Status
	}
	tests := []struct {
		name        string
		fields      fields
		wantCode    status.Code
		wantError   string
		wantMessage string
	}{
		{
			name: "TestHappyPath",
			fields: fields{
				s: &status.Status{
					Code:    100000,
					Error:   "INVALID_ARGUMENT id could not be null",
					Message: "id 字段不能为空",
				},
			},
			wantCode:    status.Code_INVALID_ARGUMENT,
			wantError:   "INVALID_ARGUMENT id could not be null",
			wantMessage: "id 字段不能为空",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &_error{
				s: tt.fields.s,
			}
			if got := e.Code(); !reflect.DeepEqual(got, tt.wantCode) {
				t.Errorf("_error.Code() = %v, want %v", got, tt.wantCode)
			}
			if got := e.Error(); !reflect.DeepEqual(got, tt.wantError) {
				t.Errorf("_error.Error() = %v, want %v", got, tt.wantError)
			}
			if got := e.Message(); !reflect.DeepEqual(got, tt.wantMessage) {
				t.Errorf("_error.Message() = %v, want %v", got, tt.wantMessage)
			}
		})
	}
}

func Test_error_WithMessage(t *testing.T) {
	type fields struct {
		s *status.Status
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   giraffe.Error
	}{
		{
			name: "TestHappyPath",
			fields: fields{
				s: &status.Status{},
			},
			args: args{
				message: "hello world",
			},
			want: &_error{
				s: &status.Status{
					Message: "hello world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &_error{
				s: tt.fields.s,
			}
			if got := e.WithMessage(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_error.WithMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
