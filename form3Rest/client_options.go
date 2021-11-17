package form3rest

import "net/http"

type clientOptions struct {
	headers http.Header
	baseUrl string
}

type ClientConfigure interface {
	SetHeaders(headers http.Header) ClientConfigure
	Configure() 
}
