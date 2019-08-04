package restv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"github.com/easyops-cn/giraffe-micro"
)

var jsonpbUnmarshaler = jsonpb.Unmarshaler{
	AllowUnknownFields: true,
}

func parseResponse(rule giraffe.HttpRule, resp *http.Response, out interface{}) error {
	bodyData, err := ioutil.ReadAll(resp.Body) // 读取 response body
	if err != nil {
		return fmt.Errorf("unexpected error when reading response body: %s", err.Error())
	}

	// 如果指定了 response message 字段, 则提取字段
	if body := rule.GetResponseBody(); body != "" {
		m := map[string]json.RawMessage{}
		if err := json.Unmarshal(bodyData, &m); err != nil {
			return fmt.Errorf("unmarshal response body failed: %s", err)
		}
		bodyData = m[body]
	}

	// 错误码正常, 解析 message
	if err := jsonpbUnmarshaler.Unmarshal(bytes.NewReader(bodyData), out.(proto.Message)); err != nil {
		return fmt.Errorf("unexpected message when unmarshal body data: %s", err.Error())
	}
	return nil
}
