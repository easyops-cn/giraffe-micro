package restv2

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func Test_copyBody(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantBs  []byte
		wantErr bool
	}{
		{
			name: "error path",
			args: args{
				req: &http.Request{
					Body: &errReadCloser{},
				},
			},
			wantErr: true,
			wantBs:  []byte{},
		},
		{
			name: "happy path",
			args: args{
				req: &http.Request{
					Body: io.NopCloser(bytes.NewReader([]byte("Hello"))),
				},
			},
			wantErr: false,
			wantBs:  []byte("Hello"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBs, err := copyBody(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("copyBody() gotBs = %v, want %v", gotBs, tt.wantBs)
			}

			if !tt.wantErr {
				tt.args.req.GetBody()
			}
		})
	}
}
