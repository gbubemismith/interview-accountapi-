package accounts

import (
	"errors"
	"net/http"

	"github.com/gbubemismith/interview-accountapi-/form3http"
)

const (
	baseAddress = "http://localhost:8080"
	url         = "/v1/organisation/accounts"
)

var (
	httpClient = getHttpClient()
)

//private business logic
func (a *accountFunction) create(body interface{}) (*AccountData, error) {
	response, err := httpClient.Post(url, nil, body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		var apiError ErrorResponse
		if err := response.UnmarshalJson(&apiError); err != nil {
			return nil, errors.New("error retrieving form 3 error message")
		}
		return nil, errors.New(apiError.ErrorMessage)
	}

	var result AccountData
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

//implementing a singleton for reusing the http client
func getHttpClient() form3http.Client {
	//create Content-Type: application/vnd.api+json as defined in form3 documentation
	headers := make(http.Header)
	headers.Set("Content-Type", "application/vnd.api+json")

	client := form3http.NewOptions().SetHeaders(headers).Configure()

	return client
}
