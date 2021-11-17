package form3rest

import (
	"net/http"
)

//struct to add configurations needs to make a request
type restClient struct {
	options *clientOptions
	client  *http.Client
}

//interface for the common rest verbs, other applications would interact with this directly
type HttpClient interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

//Public method starts with uppercase
//Performs a get operation
func (c *restClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

//Performs a post operations, can make a post request with body of any type
func (c *restClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

//Performs a delete operation
func (c *restClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
