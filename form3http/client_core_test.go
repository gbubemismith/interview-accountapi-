package form3http

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/vnd.api+json")
	client.options.headers = commonHeaders

	//Execution
	resultHeaders := client.getRequestHeaders()

	//Validation
	if len(resultHeaders) != 1 {
		t.Error("we expect 1 header")
	}
}
