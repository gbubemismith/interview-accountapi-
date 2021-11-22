package accounts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gbubemismith/interview-accountapi-/form3http"
)

const (
	baseAddress = "http://localhost:8080"
	apiPath     = "/v1/organisation/accounts"
)

var (
	httpClient = getHttpClient()
)

//private business logic to create an account
//method sends a post request to form3's endpoint
func (a *accountFunction) create(body interface{}) (*AccountData, error) {

	response, err := httpClient.Post(apiPath, body)
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

//method retrives a created account
func (a *accountFunction) getAccount(accountId string) (*AccountData, error) {

	response, err := httpClient.Get(apiPath + "/" + accountId)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
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

//method deletes an account resource
func (a *accountFunction) deleteResource(accountId string, version int64) error {
	//build delete url with query paramater
	deleteUrl := apiPath + "/" + accountId + "?version=" + strconv.Itoa(int(version))

	response, err := httpClient.Delete(deleteUrl)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		if response.StatusCode == http.StatusNotFound {
			return errors.New("specified resource does not exist")
		}

		if response.StatusCode == http.StatusNotFound {
			return errors.New("specified resource does not exist")
		}

		return errors.New("something unusual occured, please try again later")
	}

	return nil
}

//implementing a singleton for reusing the http client
func getHttpClient() form3http.Client {
	//create Content-Type: application/vnd.api+json as defined in form3 documentation
	//using httpclient buit for form3's purpose

	headers := make(http.Header)
	headers.Set("Content-Type", "application/vnd.api+json")

	client := form3http.NewOptions().SetBaseUrl(baseAddress).SetHeaders(headers).Configure()

	return client
}
