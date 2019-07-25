package hack

import (
	"context"
	"errors"
	"reflect"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"

	"github.com/easyops-cn/giraffe-micro"
)

var mockContract = &giraffeproto.Contract{
	Name:    "mock",
	Version: "",
}

type mockClient struct{}

func (c *mockClient) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}) error {
	if md.Contract.GetName() != mockContract.GetName() {
		return errors.New("")
	}
	return nil
}

func (c *mockClient) NewStream(ctx context.Context, sd *giraffe.StreamDesc) (giraffe.ClientStream, error) {
	if sd.Contract.GetName() != mockContract.GetName() {
		return nil, errors.New("")
	}
	return nil, nil
}

func Test_client_Invoke(t *testing.T) {
	type fields struct {
		c        giraffe.Client
		contract giraffe.Contract
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
				c:        &mockClient{},
				contract: mockContract,
			},
			args: args{
				ctx: context.Background(),
				md:  &giraffe.MethodDesc{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				c:        tt.fields.c,
				contract: tt.fields.contract,
			}
			if err := c.Invoke(tt.args.ctx, tt.args.md, tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("client.Invoke() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_NewStream(t *testing.T) {
	type fields struct {
		c        giraffe.Client
		contract giraffe.Contract
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
			name: "Test_HappyPath",
			fields: fields{
				c:        &mockClient{},
				contract: mockContract,
			},
			args: args{
				ctx: context.Background(),
				sd:  &giraffe.StreamDesc{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				c:        tt.fields.c,
				contract: tt.fields.contract,
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

func TestClientWithServiceName(t *testing.T) {
	type args struct {
		serviceName string
		c           giraffe.Client
	}
	tests := []struct {
		name string
		args args
		want giraffe.Client
	}{
		{
			name: "Test_HappyPath",
			args: args{
				serviceName: "logic.cmdb",
				c:           &mockClient{},
			},
			want: &client{
				c: &mockClient{},
				contract: &giraffeproto.Contract{
					Name:    "logic.cmdb",
					Version: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClientWithServiceName(tt.args.serviceName, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
