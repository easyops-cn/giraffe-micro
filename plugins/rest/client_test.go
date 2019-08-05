package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/types"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/easyops-cn/giraffe-micro"
)

type fakeNameService struct{}

func (*fakeNameService) GetAddress(contract giraffe.Contract) (string, error) { return "", nil }
func (*fakeNameService) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{}, nil
}

type errNameService struct{}

func (*errNameService) GetAddress(contract giraffe.Contract) (string, error) {
	return "", errors.New("")
}
func (*errNameService) GetAllAddresses(contract giraffe.Contract) ([]string, error) {
	return []string{}, errors.New("")
}

var ns = &fakeNameService{}
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
					NameService: ns,
					Tracer:      tracer,
					Timeout:     120 * time.Second,
					Transport:   transport,
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
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewClient() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
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
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewClient() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
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
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("client.NewStream() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

type statusOKTransport struct{}

func (t *statusOKTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	rec.Code = http.StatusOK
	rec.Header().Add("Content-Type", "application/json")
	rec.Body = bytes.NewBuffer([]byte("{}"))
	return rec.Result(), nil
}

type statusNotFoundTransport struct{}

func (t *statusNotFoundTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	rec.Code = http.StatusNotFound
	rec.Header().Add("Content-Type", "application/json")
	rec.Body = bytes.NewBuffer([]byte("{}"))
	return rec.Result(), nil
}

type failedTransport struct{}

func (t *failedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("always error")
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
		{
			name: "Test_HappyPath",
			fields: fields{
				c: &http.Client{
					Transport: &statusOKTransport{},
				},
				options: ClientOptions{
					NameService: ns,
				},
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					ServiceName: "instance.rpc",
					MethodName:  "GetDetail",
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/object/:objectId/instance/:instanceId",
						},
						Body: "",
					},
				},
				in:  &getDetailRequest{},
				out: &types.Struct{},
			},
			wantErr: false,
		},
		{
			name: "Test_WithoutHttpRule",
			fields: fields{
				c: &http.Client{
					Transport: &statusOKTransport{},
				},
				options: ClientOptions{
					NameService: ns,
				},
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					ServiceName: "instance.rpc",
					MethodName:  "GetDetail",
				},
				in:  &getDetailRequest{},
				out: &types.Struct{},
			},
			wantErr: true,
		},
		{
			name: "Test_WhenNameServiceError",
			fields: fields{
				c: &http.Client{
					Transport: &statusOKTransport{},
				},
				options: ClientOptions{
					NameService: &errNameService{},
				},
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					ServiceName: "instance.rpc",
					MethodName:  "GetDetail",
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/object/:objectId/instance/:instanceId",
						},
						Body: "",
					},
				},
				in:  &getDetailRequest{},
				out: &types.Struct{},
			},
			wantErr: true,
		},
		{
			name: "Test_ErrorResponse",
			fields: fields{
				c: &http.Client{
					Transport: &statusNotFoundTransport{},
				},
				options: ClientOptions{
					NameService: ns,
				},
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					ServiceName: "instance.rpc",
					MethodName:  "GetDetail",
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/object/:objectId/instance/:instanceId",
						},
						Body: "",
					},
				},
				in:  &getDetailRequest{},
				out: &types.Struct{},
			},
			wantErr: true,
		},
		{
			name: "Test_RequestFailed",
			fields: fields{
				c: &http.Client{
					Transport: &failedTransport{},
				},
				options: ClientOptions{
					NameService: ns,
				},
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					ServiceName: "instance.rpc",
					MethodName:  "GetDetail",
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/object/:objectId/instance/:instanceId",
						},
						Body: "",
					},
				},
				in:  &getDetailRequest{},
				out: &types.Struct{},
			},
			wantErr: true,
		},
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
