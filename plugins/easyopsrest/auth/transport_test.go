package auth

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

type mockTransport struct {}

func (*mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := FromRequest(req)
	_, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("no user info")
	}

	return &http.Response{}, nil
}

func Test_transport_RoundTrip(t *testing.T) {
	type fields struct {
		rt http.RoundTripper
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				rt: &mockTransport{},
			},
			args: args{
				r: func() *http.Request {
					r, _ := http.NewRequest("GET", "/", nil)
					return r.WithContext(WithUserInfo(context.Background(), UserInfo{User:"foo", Org:8888}))
				}(),
			},
			want: &http.Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &transport{
				rt: tt.fields.rt,
			}
			got, err := transport.RoundTrip(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("transport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("transport.RoundTrip() = %v, want %v", got, tt.want)
			//}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestNewTransport(t *testing.T) {
	type args struct {
		rt http.RoundTripper
	}
	tests := []struct {
		name string
		args args
		want http.RoundTripper
	}{
		{
			name: "Test_HappyPath",
			args: args{
				rt: &mockTransport{},
			},
			want: &transport{
				rt: &mockTransport{},
			},
		},
		{
			name: "Test_WithNilTransport",
			args: args{
				rt: nil,
			},
			want: &transport{
				rt: http.DefaultTransport,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := NewTransport(tt.args.rt); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewTransport() = %v, want %v", got, tt.want)
			//}
			got := NewTransport(tt.args.rt)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
