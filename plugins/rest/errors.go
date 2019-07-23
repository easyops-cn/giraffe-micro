package rest

import (
	"net/http"
)

type restError http.Response

func (r *restError) Error() string {
	return r.Status
}

func (r *restError) HttpResponse() *http.Response {
	return (*http.Response)(r)
}

func isErrorResponse(resp *http.Response) *restError {
	if resp.StatusCode < 400 {
		return nil
	}
	return (*restError)(resp)
}
