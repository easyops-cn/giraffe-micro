package restv2

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/easyops-cn/giraffe-micro"
)

// Middleware 中间件定义
type Middleware interface {
	NewRequest(rule giraffe.HttpRule, in interface{}) (*http.Request, error)
	ParseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error
}

// Client REST Client对象
type Client struct {
	*http.Client
	Middleware  Middleware
	NameService giraffe.NameService
	retryConf   RetryConfig
}

// ClientOption Client 配置函数
type ClientOption func(c *Client)

// Invoke 单次请求方法
func (c *Client) Invoke(ctx context.Context, md *giraffe.MethodDesc, in interface{}, out interface{}, opts ...giraffe.CallOption) error {
	req, err := c.middleware().NewRequest(md.HttpRule, in)
	if err != nil {
		return err
	}

	req.Header.Set("giraffe-contract-name", md.Contract.GetName())
	req.Header.Set("giraffe-contract-version", md.Contract.GetVersion())
	req = req.WithContext(ctx)
	resp, err := c.Call(md.Contract, req, opts...)
	if resp != nil {
		defer func() {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}()
	}

	if err != nil {
		return err
	}
	return c.middleware().ParseResponse(md.HttpRule, resp, out)
}

// NewStream 流式请求方法(未实现)
func (c *Client) NewStream(context.Context, *giraffe.StreamDesc, ...giraffe.CallOption) (giraffe.ClientStream, error) {
	return nil, errors.New("not supported")
}

func (c *Client) middleware() Middleware {
	if c.Middleware != nil {
		return c.Middleware
	}
	return DefaultMiddleware
}

func (c *Client) httpClient() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

// Call 请求函数
func (c *Client) Call(contract giraffe.Contract, req *http.Request, opts ...giraffe.CallOption) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("request was nil")
	}

	options := &giraffe.CallOptions{
		Metadata: map[string][]string{},
	}
	for _, o := range opts {
		o(options)
	}

	for k, v := range options.Metadata {
		if v != nil && len(v) > 0 {
			req.Header.Set(k, v[0])
		}
	}
	if hostname := req.Header.Get("host"); hostname != "" {
		req.Host = hostname
	}
	if c.NameService == nil {
		// 原逻辑中NameService没有设置也会发送http请求，保留这个特征
		return c.httpClient().Do(req)
	}
	return c.sendWithENS(req, contract)
}

func (c *Client) sendWithENS(req *http.Request, contract giraffe.Contract) (resp *http.Response, err error) {
	// ENS服务发现
	addrList, err := c.getAllAddressesWithENS(req.Context(), contract)
	if err != nil {
		return
	}

	// 备份 http body
	var originalBody []byte
	req.URL.Scheme = "http"
	if req != nil && req.Body != nil {
		originalBody, _ = copyBody(req)
	}

	// 如果 retryConf.Enabled 为false, 获取 sendCount 的值为1，表示只执行一次，不会重试
	sendCount := c.retryConf.getSendCount()

	retryInterval := c.retryConf.RetryInterval
	addr := addrList[0]
	i := 0
	unavailableRetry := false
	for ; i < sendCount; i++ {
		// 当i为1时，表示已经是重试循环, 需要根据情况进行等待
		// 当 单节点连接被拒绝重试 或者 503重试, 需要等待一段时间后再发起请求
		// 多节点下，连接被拒绝，就不等待了，直接访问其他节点
		if i > 0 && (len(addrList) <= 0 || unavailableRetry) {
			time.Sleep(retryInterval)
		}
		// 服务重试，以轮询策略为节点选择策略
		if addr == "" {
			addr = addrList[i%len(addrList)]
		}
		// 根据执行次数, 轮询节点
		req.URL.Host = addr

		resp, err = c.httpClient().Do(req)
		if err != nil {
			// connection refuse 重试机制
			if errors.Is(err, syscall.ECONNREFUSED) {
				resetBody(req, originalBody)
				addr = "" // 当前节点的访问被拒绝, addr置空, 获取其他节点
				unavailableRetry = false
				continue
			}
			return // 非 connection refuse 错误直接退出
		}
		if err == nil && resp.StatusCode != http.StatusServiceUnavailable {
			return // 非 503 异常, 直接退出不重试
		}

		// 503异常的重试机制
		retryAfterStr := resp.Header.Get("Retry-After")
		retryAfter, _ := strconv.Atoi(retryAfterStr)
		if retryAfter <= 0 {
			// 没有设置 Retry-After , 重试机制失效, 直接退出
			return
		}
		// Retry-After单位为秒: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
		// TODO 对获取的retryAfter进行校验，不要让它因为错误的设置而等待超长的时间
		retryInterval = time.Duration(retryAfter) * time.Second
		unavailableRetry = true
		_ = resp.Body.Close() // 重试, 所以释放当前fd
		resetBody(req, originalBody)
	}
	return
}

func (c *Client) getAllAddressesWithENS(ctx context.Context, contract giraffe.Contract) (addresses []string, err error) {
	if c.NameService == nil {
		return
	}
	addresses, err = c.NameService.GetAllAddresses(ctx, contract)
	if err != nil {
		return
	}
	// 排序
	sort.Strings(addresses)
	return
}

// NewClient Client实例化函数
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		Client:      &http.Client{},
		Middleware:  DefaultMiddleware,
		NameService: nil,
	}
	for _, o := range opts {
		o(c)
	}
	c.retryConf.init()
	return c
}

// WithClient 注入 http.Client
func WithClient(client *http.Client) ClientOption {
	return func(c *Client) {
		if client == nil {
			client = &http.Client{}
		}
		c.Client = client
	}
}

// WithNameService 注入 NameService
func WithNameService(n giraffe.NameService) ClientOption {
	return func(c *Client) {
		c.NameService = n
	}
}

func WithRetryConfig(conf RetryConfig) ClientOption {
	return func(c *Client) {
		c.retryConf = conf
	}
}
