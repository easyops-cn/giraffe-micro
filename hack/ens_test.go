package hack

import (
	"reflect"
	"testing"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"

	"github.com/easyops-cn/giraffe-micro"
)

func Test_ens_GetAddress(t *testing.T) {
	type fields struct {
		addr string
	}
	type args struct {
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				addr: "192.168.100.162:80",
			},
			args: args{
				contract: &giraffeproto.Contract{
					Name:    "xxx",
					Version: "",
				},
			},
			want:    "192.168.100.162:80",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ens{
				addr: tt.fields.addr,
			}
			got, err := e.GetAddress(tt.args.contract)
			if (err != nil) != tt.wantErr {
				t.Errorf("ens.GetAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ens.GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ens_GetAllAddresses(t *testing.T) {
	type fields struct {
		addr string
	}
	type args struct {
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				addr: "192.168.100.162:80",
			},
			args: args{
				contract: &giraffeproto.Contract{
					Name:    "xxx",
					Version: "",
				},
			},
			want:    []string{"192.168.100.162:80"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ens{
				addr: tt.fields.addr,
			}
			got, err := e.GetAllAddresses(tt.args.contract)
			if (err != nil) != tt.wantErr {
				t.Errorf("ens.GetAllAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ens.GetAllAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStaticAddress(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want giraffe.NameService
	}{
		{
			name: "Test_HappyPath",
			args: args{
				host: "192.168.100.162",
				port: 8080,
			},
			want: &ens{
				addr: "192.168.100.162:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StaticAddress(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaticAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
