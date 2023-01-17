package restv2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	giraffeproto "github.com/easyops-cn/go-proto-giraffe"
	"github.com/go-test/deep"
	"github.com/gogo/protobuf/types"

	"github.com/easyops-cn/giraffe-micro"
	"github.com/easyops-cn/giraffe-micro/test/mock/mock_restv2"
)

type errNameService struct {
	addr  string
	addrs []string
	err   error
}

func (e *errNameService) GetAddress(ctx context.Context, contract giraffe.Contract) (string, error) {
	return e.addr, e.err
}

func (e *errNameService) GetAllAddresses(ctx context.Context, contract giraffe.Contract) ([]string, error) {
	return e.addrs, e.err
}

type mockTransport struct {
	err    error
	listen string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.err != nil {
		return nil, t.err
	}
	rec := httptest.NewRecorder()
	if t.listen != req.URL.Host {
		rec.Code = http.StatusBadGateway
		return rec.Result(), nil
	}
	rec.Code = http.StatusOK
	rec.Header().Add("Content-Type", "application/json")
	rec.Body = bytes.NewBuffer([]byte("{\"code\":\"100014\"}"))
	resp := rec.Result()
	resp.Request = req
	return resp, nil
}

func TestClient_Call(t *testing.T) {
	var emptyRequest = func() *http.Request {
		req, _ := http.NewRequest("GET", "/", nil)
		return req
	}
	var callOption = func(key string, val string) giraffe.CallOption {
		return func(o *giraffe.CallOptions) {
			o.Metadata[key] = []string{val}
		}
	}
	type fields struct {
		Client      *http.Client
		Middleware  Middleware
		NameService giraffe.NameService
	}
	type args struct {
		contract giraffe.Contract
		req      *http.Request
		opts     []giraffe.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "Test_WithCallOptions",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						listen: "192.168.100.162:8080",
					},
				},
				Middleware:  DefaultMiddleware,
				NameService: StaticAddress("192.168.100.162:8080"),
			},
			args: args{
				contract: &giraffeproto.Contract{
					Name:    "easyops.api.cmdb.instance.GetDetail",
					Version: "1.0",
				},
				req: emptyRequest(),
				opts: []giraffe.CallOption{callOption("host", "cmdb.easyops-only.com"),
					callOption("user", "index"),
					callOption("org", "8888"),
				},
			},
			want: func() *http.Response {
				rec := httptest.NewRecorder()
				rec.Code = http.StatusOK
				rec.Header().Add("Content-Type", "application/json")
				rec.Body = bytes.NewBuffer([]byte("{\"code\":\"100014\"}"))
				resp := rec.Result()
				resp.Request = emptyRequest()
				resp.Request.URL.Scheme = "http"
				resp.Request.URL.Host = "192.168.100.162:8080"
				resp.Request.Header.Add("host", "cmdb.easyops-only.com")
				resp.Request.Header.Add("user", "index")
				resp.Request.Header.Add("org", "8888")
				resp.Request.Host = "cmdb.easyops-only.com"
				return resp
			}(),
			wantErr: false,
		},
		{
			name: "Test_WhenRequestWasNil",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						listen: "192.168.100.162:8080",
					},
				},
				Middleware: DefaultMiddleware,
				NameService: &errNameService{
					err: errors.New("mock error"),
				},
			},
			args: args{
				contract: &giraffeproto.Contract{
					Name:    "easyops.api.cmdb.instance.GetDetail",
					Version: "1.0",
				},
				req:  nil,
				opts: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_WhenNameServiceFailed",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						listen: "192.168.100.162:8080",
					},
				},
				Middleware: DefaultMiddleware,
				NameService: &errNameService{
					err: errors.New("mock error"),
				},
			},
			args: args{
				contract: &giraffeproto.Contract{
					Name:    "easyops.api.cmdb.instance.GetDetail",
					Version: "1.0",
				},
				req:  emptyRequest(),
				opts: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:      tt.fields.Client,
				Middleware:  tt.fields.Middleware,
				NameService: tt.fields.NameService,
			}
			got, err := c.Call(tt.args.contract, tt.args.req, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestClient_Invoke(t *testing.T) {
	type fields struct {
		Client      *http.Client
		Middleware  Middleware
		NameService giraffe.NameService
	}
	type args struct {
		ctx  context.Context
		md   *giraffe.MethodDesc
		in   interface{}
		out  interface{}
		opts []giraffe.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Struct
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						listen: "192.168.100.162:8080",
					},
				},
				Middleware:  DefaultMiddleware,
				NameService: StaticAddress("192.168.100.162:8080"),
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/",
						},
					},
				},
				in: &mock_restv2.GetDetailRequest{
					ObjectId:   "HOST",
					InstanceId: "xxx",
				},
				out:  &types.Struct{},
				opts: nil,
			},
			want: &types.Struct{
				Fields: map[string]*types.Value{
					"code": {
						Kind: &types.Value_StringValue{StringValue: "100014"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test_WithoutHttpRule",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						listen: "192.168.100.162:8080",
					},
				},
				Middleware:  DefaultMiddleware,
				NameService: StaticAddress("192.168.100.162:8080"),
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
				},
				in:   &mock_restv2.GetDetailRequest{},
				out:  &types.Struct{},
				opts: nil,
			},
			want:    &types.Struct{},
			wantErr: true,
		},
		{
			name: "",
			fields: fields{
				Client: &http.Client{
					Transport: &mockTransport{
						err:    errors.New("always failed"),
						listen: "192.168.100.162:8080",
					},
				},
				Middleware:  DefaultMiddleware,
				NameService: StaticAddress("192.168.100.162:8080"),
			},
			args: args{
				ctx: context.Background(),
				md: &giraffe.MethodDesc{
					Contract: &giraffeproto.Contract{
						Name:    "easyops.api.cmdb.instance.GetDetail",
						Version: "1.0",
					},
					HttpRule: &giraffeproto.HttpRule{
						Pattern: &giraffeproto.HttpRule_Get{
							Get: "/",
						},
					},
				},
				in: &mock_restv2.GetDetailRequest{
					ObjectId:   "HOST",
					InstanceId: "xxx",
				},
				out:  &types.Struct{},
				opts: nil,
			},
			want:    &types.Struct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:      tt.fields.Client,
				Middleware:  tt.fields.Middleware,
				NameService: tt.fields.NameService,
			}
			err := c.Invoke(tt.args.ctx, tt.args.md, tt.args.in, tt.args.out, tt.args.opts...)
			got := tt.args.out
			if (err != nil) != tt.wantErr {
				t.Errorf("Invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewStream(t *testing.T) {
	type fields struct {
		Client      *http.Client
		Middleware  Middleware
		NameService giraffe.NameService
	}
	type args struct {
		ctx  context.Context
		sd   *giraffe.StreamDesc
		opts []giraffe.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    giraffe.ClientStream
		wantErr bool
	}{
		{
			name:    "Test_HappyPath",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:      tt.fields.Client,
				Middleware:  tt.fields.Middleware,
				NameService: tt.fields.NameService,
			}
			got, err := c.NewStream(tt.args.ctx, tt.args.sd, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_middleware(t *testing.T) {
	type mockMiddleware struct{ Middleware }
	mock := &mockMiddleware{}
	type fields struct {
		Client      *http.Client
		Middleware  Middleware
		NameService giraffe.NameService
	}
	tests := []struct {
		name   string
		fields fields
		want   Middleware
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				NameService: nil,
			},
			want: DefaultMiddleware,
		},
		{
			name: "Test_HappyPath",
			fields: fields{
				Middleware:  mock,
				NameService: nil,
			},
			want: mock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:      tt.fields.Client,
				Middleware:  tt.fields.Middleware,
				NameService: tt.fields.NameService,
			}
			if got := c.middleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_httpClient(t *testing.T) {
	type fields struct {
		Client      *http.Client
		Middleware  Middleware
		NameService giraffe.NameService
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				Client: nil,
			},
			want: http.DefaultClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:      tt.fields.Client,
				Middleware:  tt.fields.Middleware,
				NameService: tt.fields.NameService,
			}
			if got := c.httpClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		opts []ClientOption
	}
	tests := []struct {
		name string
		args args
		want giraffe.Client
	}{
		{
			name: "Test_HappyPath",
			args: args{
				opts: []ClientOption{
					WithClient(nil),
					WithClient(&http.Client{
						Timeout: 120 * time.Minute,
					}),
					WithNameService(nil),
					WithNameService(StaticAddress("192.168.100.162:8080")),
				},
			},
			want: &Client{
				Client: &http.Client{
					Timeout: 120 * time.Minute,
				},
				Middleware:  &BaseMiddleware{},
				NameService: StaticAddress("192.168.100.162:8080"),
				retryConf: RetryConfig{
					RetryInterval: defaultWaitDuration,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetryConfig(t *testing.T) {
	type args struct {
		conf RetryConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "happy path",
			args: args{
				conf: RetryConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithRetryConfig(tt.args.conf)
			got(&Client{})
		})
	}
}

func TestClient_getAllAddressesWithNS(t *testing.T) {
	type mockConfig struct {
		nameServiceNil bool
		getErr         bool
	}
	tests := []struct {
		name          string
		mockConfig    mockConfig
		wantAddresses []string
		wantErr       bool
	}{
		{
			name: "NameService nil",
			mockConfig: mockConfig{
				nameServiceNil: true,
			},
			wantErr: false,
		},
		{
			name: "error path",
			mockConfig: mockConfig{
				getErr: true,
			},
			wantErr: true,
		},
		{
			name: "happy path",
			mockConfig: mockConfig{
				getErr: false,
			},
			wantErr:       false,
			wantAddresses: []string{"127.0.0.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nameService giraffe.NameService
			if !tt.mockConfig.nameServiceNil {
				if tt.mockConfig.getErr {
					nameService = &errNameService{
						err: errors.New("mock error"),
					}
				} else {
					nameService = &errNameService{
						addrs: []string{"127.0.0.1"},
					}
				}

			}
			c := &Client{
				NameService: nameService,
			}
			gotAddresses, err := c.getAllAddressesWithNS(nil, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllAddressesWithNS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAddresses, tt.wantAddresses) {
				t.Errorf("getAllAddressesWithNS() gotAddresses = %v, want %v", gotAddresses, tt.wantAddresses)
			}
		})
	}
}

func TestClient_sendWithENS_retry_fail(t *testing.T) {
	req1, _ := http.NewRequest("POST", "http://127.0.0.1:80", bytes.NewReader([]byte("Hello")))
	t.Run("retry failed", func(t *testing.T) {
		mockNameService := &errNameService{
			addrs: []string{"127.0.0.1:80"},
		}
		c := &Client{
			Client:      http.DefaultClient,
			NameService: mockNameService,
			retryConf: RetryConfig{
				Enabled:       true,
				Retries:       2,
				RetryInterval: 10 * time.Millisecond,
			},
		}
		_, _ = c.sendWithENS(req1, nil)
	})
}

type testHandler1 struct {
	retryAfter string
}

func (s *testHandler1) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	reqBs, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(reqBs))

	resp.Header().Set("Retry-After", s.retryAfter)
	resp.WriteHeader(503)
}

func TestClient_sendWithENS_503_retry(t *testing.T) {
	server := httptest.NewServer(&testHandler1{
		retryAfter: "1",
	})
	defer server.Close()

	req1, _ := http.NewRequest("POST", server.URL, bytes.NewReader([]byte("Hello")))
	t.Run("retry failed", func(t *testing.T) {
		port := strings.Split(server.URL, ":")[2]
		mockNameService := &errNameService{
			addrs: []string{fmt.Sprintf("127.0.0.1:%s", port)},
		}
		c := &Client{
			Client:      http.DefaultClient,
			NameService: mockNameService,
			retryConf: RetryConfig{
				Enabled:       true,
				Retries:       2,
				RetryInterval: 10 * time.Millisecond,
			},
		}
		_, _ = c.sendWithENS(req1, nil)
	})
}

func TestClient_sendWithENS_503_retry_no_retryAfter(t *testing.T) {
	server := httptest.NewServer(&testHandler1{})
	defer server.Close()

	req1, _ := http.NewRequest("POST", server.URL, bytes.NewReader([]byte("Hello")))
	t.Run("retry failed", func(t *testing.T) {
		port := strings.Split(server.URL, ":")[2]
		mockNameService := &errNameService{
			addrs: []string{fmt.Sprintf("127.0.0.1:%s", port)},
		}
		c := &Client{
			Client:      http.DefaultClient,
			NameService: mockNameService,
			retryConf: RetryConfig{
				Enabled:       true,
				Retries:       2,
				RetryInterval: 10 * time.Millisecond,
			},
		}
		_, _ = c.sendWithENS(req1, nil)
	})
}
