package go_http_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	defaultUsername = "barry"
	defaultPasswd   = "yan"
)

type SDK struct {
	Host   string
	User   string
	Passwd string
	header map[string]string
}

func NewSDK(host, userName, passWd string) (*SDK, error) {
	sdk := &SDK{
		Host:   host,
		User:   userName,
		Passwd: passWd,
	}
	if sdk.checkAuth() {
		sdk.header = map[string]string{"name": "barry yan"}
		return sdk, nil
	}
	return nil, errors.New("auth err")
}

func (s *SDK) checkAuth() bool {
	return s.User == defaultUsername && s.Passwd == defaultPasswd
}

func (s *SDK) Create(request CreateRequest) Err {
	path := "/create"

	bytes, err := json.Marshal(request)
	if err != nil {
		return ErrInnerErr
	}
	resp, err := NewHttpClient(fmt.Sprintf("%s%s", s.Host, path),
		WithBody(bytes),
		WithHeader(map[string]string{"name": "barry yan"})).Post()
	if err != nil {
		return ErrInnerErr
	}

	httpResp := &Err{}
	if err := json.Unmarshal(resp, httpResp); err != nil {
		return ErrInnerErr
	}
	if httpResp.Code != http.StatusOK {
		return *httpResp
	}
	return ErrOk
}

func (s *SDK) Get(key string) (string, Err) {
	path := "/get"

	resp, err := NewHttpClient(fmt.Sprintf("%s%s/%s", s.Host, path, key),
		WithHeader(map[string]string{"name": "barry yan"})).Get()
	if err != nil {
		return "", ErrInnerErr
	}
	return string(resp), ErrOk
}
