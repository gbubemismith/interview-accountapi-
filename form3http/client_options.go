package form3http

import "net/http"

type clientOptions struct {
	headers http.Header
	baseUrl string
}

type ClientConfigure interface {
	SetHeaders(headers http.Header) ClientConfigure
	Configure() Client
}

func NewOptions() ClientConfigure {
	options := &clientOptions{}
	return options
}

func (c *clientOptions) SetHeaders(headers http.Header) ClientConfigure {
	c.headers = headers
	return c
}

func (c *clientOptions) Configure() Client {
	client := httpClient{
		options: c,
	}

	return &client
}
