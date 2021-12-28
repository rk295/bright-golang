package mock

import (
	"io"
	"net/http"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func DefaultDo(r io.ReadCloser, sc int) func(req *http.Request) (*http.Response, error) {
	return func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: sc,
			Body:       r,
		}, nil
	}
}
