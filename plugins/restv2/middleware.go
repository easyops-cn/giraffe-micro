package restv2

import (
	"net/http"

	"github.com/easyops-cn/giraffe-micro"
)

var DefaultMiddleware = &middleware{}

type middleware struct{}

func (m *middleware) NewRequest(rule giraffe.HttpRule, in interface{}) (*http.Request, error) {
	req, err := newRequest(rule, in)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (m *middleware) ParseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error {
	return parseResponse(rule, resp, out)
}
