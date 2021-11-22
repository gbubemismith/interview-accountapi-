package form3http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//do
func (c *httpClient) do(method string, url string, body interface{}) (*Response, error) {

	addHeaders := c.getRequestHeaders()

	requestBody, err := c.marshalRequestBody(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("creating request failed")
	}

	request.Header = addHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resultResp := Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		body:       responseBody,
	}

	return &resultResp, nil
}

func (c *httpClient) getRequestHeaders() http.Header {
	result := make(http.Header)

	//Add default headers to request, content-type, content
	for header, value := range c.options.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) marshalRequestBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	return json.Marshal(body)
}

//
func (c *httpClient) getHttpClient() *http.Client {

	c.client = &http.Client{}
	return c.client
}
