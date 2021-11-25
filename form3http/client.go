package form3http

import (
	"net/http"
)

//struct to add configurations needs to make a request
type httpClient struct {
	options *clientOptions
	client  *http.Client
}

//interface for the common rest verbs, other applications would interact with this directly
type Client interface {
	Get(url string) (*Response, error)
	Post(url string, body interface{}) (*Response, error)
	Delete(url string) (*Response, error)
}

//Public method starts with uppercase
//Performs a get operation
func (c *httpClient) Get(url string) (*Response, error) {
	return c.do(http.MethodGet, url, nil)
}

//Performs a post operations, can make a post request with body of any type
func (c *httpClient) Post(url string, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, body)
}

//Performs a delete operation
func (c *httpClient) Delete(url string) (*Response, error) {
	return c.do(http.MethodDelete, url, nil)
}
