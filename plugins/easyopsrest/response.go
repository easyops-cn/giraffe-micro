package easyopsrest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/easyops-cn/giraffe-micro/gerr"
	"github.com/easyops-cn/giraffe-micro/status"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

type responseBody struct {
	Code        int             `json:"code"`
	CodeExplain string          `json:"codeExplain"`
	Error       string          `json:"error"`
	Data        json.RawMessage `json:"data"`
}

func parseResponse(resp *http.Response, out interface{}) error {
	body := new(responseBody)
	// 尝试解释 body, 解释失败则返回未知错误
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gerr.UnknownErrorf("%s, unexpected error when reading response body: %s", resp.Status, err.Error())
	}
	if err := json.Unmarshal(data, body); err != nil {
		return gerr.UnknownErrorf("%s, unexpected response body: %s", resp.Status, err.Error())
	}

	// 错误码不为0, 返回标准错误
	if body.Code != 0 {
		return gerr.FromProto(&status.Status{
			Code:    status.Code(body.Code),
			Error:   body.CodeExplain,
			Message: body.Error,
			//TODO 增加错误data
		})
	}

	// 状态码正常, 错误码正常, 解析 message
	if err := jsonpbUnmarshaler.Unmarshal(bytes.NewReader(body.Data), out.(proto.Message)); err != nil {
		return gerr.UnknownErrorf("%s, unexpected message when unmarshal body data: %s", resp.Status, err.Error())
	}
	return nil
}

var jsonpbUnmarshaler = jsonpb.Unmarshaler{
	AllowUnknownFields: true,
}
