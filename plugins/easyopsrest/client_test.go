package rest

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/openzipkin/zipkin-go"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/auth"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/ens"
	"github.com/easyops-cn/giraffe-micro/plugins/rest"
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
					WithTracer(tracer),
					WithRoundTripper(transport),
				},
			},
			want: func() giraffe.Client {
				r, _ := rest.NewClient(
					rest.WithTimeout(120*time.Second),
					rest.WithTracer(tracer),
					rest.WithRoundTripper(auth.NewTransport(transport)),
					rest.WithNameService(ens.NewNameService()),
				)
				return r
			}(),
			wantErr: false,
		},
		{
			name: "Test_WithNameService",
			args: args{
				opts: []ClientOption{
					WithTimeout(120 * time.Second),
					WithTracer(tracer),
					WithRoundTripper(transport),
					WithNameService(ns),
				},
			},
			want: func() giraffe.Client {
				r, _ := rest.NewClient(
					rest.WithTimeout(120*time.Second),
					rest.WithTracer(tracer),
					rest.WithRoundTripper(auth.NewTransport(transport)),
					rest.WithNameService(ns),
				)
				return r
			}(),
			wantErr: false,
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
	restClientFactory = func(opts ...rest.ClientOption) (client rest.Client, e error) {
		return nil, errors.New("always failed")
	}
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
			name: "Test_WithNameService",
			args: args{
				opts: []ClientOption{
					WithTimeout(120 * time.Second),
					WithTracer(tracer),
					WithRoundTripper(transport),
					WithNameService(ns),
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
