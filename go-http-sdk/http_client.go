package go_http_sdk

import (
	"bytes"
	"io"
	"net/http"
)

type Option func(*HttpClient)

type HttpClient struct {
	Url    string
	Body   []byte
	Header map[string]string
	Client *http.Client
}

func NewHttpClient(url string, opts ...Option) *HttpClient {
	cli := &HttpClient{Url: url, Client: &http.Client{}}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func WithBody(body []byte) Option {
	return func(client *HttpClient) {
		client.Body = body
	}
}

func WithHeader(header map[string]string) Option {
	return func(client *HttpClient) {
		client.Header = header
	}
}

func (c *HttpClient) Post() ([]byte, error) {
	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *HttpClient) Get() ([]byte, error) {
	req, err := http.NewRequest("GET", c.Url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
