package easyopsrest

import (
	"testing"

	"github.com/easyops-cn/giraffe-micro"
)

type mockMethod struct {
	giraffe.Method
	serviceName string
}
func (m *mockMethod) ServiceName() string { return m.serviceName }

type mockContract struct {
	giraffe.Method
	contractName string
	contractVersion string
}
func (m *mockContract) ContractName() string { return m.contractName }
func (m *mockContract) ContractVersion() string { return m.contractVersion }

func Test_serviceName(t *testing.T) {
	type args struct {
		method giraffe.Method
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_HappyPath",
			args: args{
				method: &mockMethod{
					serviceName: "logic.cmdb_service",
				},
			},
			want: "logic.cmdb_service",
		},
		{
			name: "Test_WithContract",
			args: args{
				method: &mockContract{
					contractName: "logic.cmdb_service",
					contractVersion: "V1.0",
				},
			},
			want: "logic.cmdb_service@V1.0",
		},
		{
			name: "Test_WithEmptyVersionContract",
			args: args{
				method: &mockContract{
					contractName: "logic.cmdb_service",
					contractVersion: "",
				},
			},
			want: "logic.cmdb_service",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serviceName(tt.args.method); got != tt.want {
				t.Errorf("serviceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
