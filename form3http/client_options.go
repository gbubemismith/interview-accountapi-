package form3http

import "net/http"

type clientOptions struct {
	headers http.Header
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

// func (c *clientOptions) SetBaseUrl(baseUrl string) ClientConfigure {
// 	c.baseUrl = baseUrl
// 	return c
// }

func (c *clientOptions) Configure() Client {
	client := httpClient{
		options: c,
	}

	return &client
}
