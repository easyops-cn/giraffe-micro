package restv2

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func copyBody(req *http.Request) (bs []byte, err error) {
	bs, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	_ = req.Body.Close()
	resetBody(req, bs)

	return
}

func resetBody(req *http.Request, originalBody []byte) {
	req.Body = io.NopCloser(bytes.NewBuffer(originalBody))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBuffer(originalBody)), nil
	}
}
