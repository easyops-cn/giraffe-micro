package rest

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/easyops-cn/giraffe-micro"
)

type FakeNameService struct{}

func (*FakeNameService) GetAddress(contract giraffe.Contract) (string, error) { return "", nil }
func (*FakeNameService) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{}, nil
}

var ns = &FakeNameService{}
var tracer = &zipkin.Tracer{}
var transport = &http.Transport{MaxIdleConns: 100}

func TestNewClient(t *testing.T) {
	type args struct {
		opts []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    giraffe.Client
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			args: args{
				opts: []ClientOption{
					WithTimeout(120 * time.Second),
					WithNameService(ns),
					WithTracer(tracer),
					WithRoundTripper(transport),
				},
			},
			want: &client{
				c: &http.Client{
					Timeout: 120 * time.Second,
					Transport: func() http.RoundTripper {
						rt, _ := zipkinhttp.NewTransport(tracer, zipkinhttp.RoundTripper(transport))
						return rt
					}(),
				},
				options: ClientOptions{
					nameService: ns,
					tracer:      tracer,
					timeout:     120 * time.Second,
					rt:          transport,
				},
			},
			wantErr: false,
		},
		{
			name: "Test_WithoutNameService",
			args: args{
				opts: []ClientOption{
					WithTimeout(120 * time.Second),
					WithTracer(tracer),
					WithRoundTripper(transport),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient2(t *testing.T) {
	type args struct {
		opts []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    giraffe.Client
		wantErr bool
	}{
		{
			name: "Test_IfZipkinNewTransportFailed",
			args: args{
				opts: []ClientOption{
					WithTimeout(120 * time.Second),
					WithNameService(ns),
					WithTracer(tracer),
					WithRoundTripper(transport),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	zipkinTransportFactory = func(tracer *zipkin.Tracer, options ...zipkinhttp.TransportOption) (tripper http.RoundTripper, e error) {
		return nil, errors.New("always failed")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_NewStream(t *testing.T) {
	type fields struct {
		c       *http.Client
		options ClientOptions
	}
	type args struct {
		ctx context.Context
		sd  *giraffe.StreamDesc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    giraffe.ClientStream
		wantErr bool
	}{
		{
			name:    "Test_NotSupported",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				c:       tt.fields.c,
				options: tt.fields.options,
			}
			got, err := c.NewStream(tt.args.ctx, tt.args.sd)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.NewStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.NewStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Invoke(t *testing.T) {
	type fields struct {
		c       *http.Client
		options ClientOptions
	}
	type args struct {
		ctx context.Context
		md  *giraffe.MethodDesc
		in  interface{}
		out interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		//{
		//	name: "Test_HappyPath",
		//	fields: fields{
		//		c: &http.Client{},
		//		options: nil,
		//	},
		//	args: args{
		//		ctx: context.Background(),
		//		md: &giraffe.MethodDesc{},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				c:       tt.fields.c,
				options: tt.fields.options,
			}
			if err := c.Invoke(tt.args.ctx, tt.args.md, tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("client.Invoke() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
