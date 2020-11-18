package restv2

import (
	"context"
	"reflect"
	"testing"

	"github.com/easyops-cn/giraffe-micro"
)

func TestStaticAddress_GetAddress(t *testing.T) {
	type args struct {
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		s       StaticAddress
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test_HappyPath",
			s:       "192.168.100.162:8080",
			args:    args{contract: nil},
			want:    "192.168.100.162:8080",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAddress(context.Background(), tt.args.contract)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStaticAddress_GetAllAddresses(t *testing.T) {
	type args struct {
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		s       StaticAddress
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Test_HappyPath",
			s:       "192.168.100.162:8080",
			args:    args{contract: nil},
			want:    []string{"192.168.100.162:8080"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllAddresses(context.Background(), tt.args.contract)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}
