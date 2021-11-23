package accounts

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestFetchAccount(t *testing.T) {
	//Test network issue reaching the form 3 api
	t.Run("TestErrorFetchingFromForm3", func(t *testing.T) {
		//enable mocking
		StartMockServer()

		AddMock(Mock{
			Method: http.MethodGet,
			Url:    "http://localhost:8080/v1/organisation/accounts/f8b9796e-c92a-4c36-96ff-a0e8cfc11f29",
			Error:  errors.New("error reaching the fetch endpoint"),
		})

		a := accountFunction{}
		response, err := a.FetchAccount("b80a0a46-7b68-472d-963e-1a7d2cc28a70")

		//validation
		if response != nil {
			t.Error("no record is expected")
		}

		if err == nil {
			t.Error("expected error")
		}

		if err.Error() != "error reaching the fetch endpoint" {
			t.Error("invalid error message gotten")
		}
	})

	//Test JSON unmarshal error
	t.Run("TestErrorUnmarshalFetchAccountResponseBody", func(t *testing.T) {
		//enable mocking
		AddMock(Mock{
			Method:     http.MethodGet,
			Url:        "http://localhost:8080/v1/organisation/accounts/f8b9796e-c92a-4c36-96ff-a0e8cfc11f29",
			StatusCode: http.StatusOK,
			ResponseBody: `
			{
				"data": {
					"attributes": {
						"alternative_names": null,
						"bank_id": "400400",
						"bank_id_code": "GBDSC",
						"base_currency": "GBP",
						"bic": "NWBKGB22",
						"country": "GB",
						"name": "Samantha Holders",
						"secondary_identification": "A1B2C3D4"
					},
					"created_on": "2021-11-21T20:29:38.867Z",
					"id": 1000,
					"modified_on": "2021-11-21T20:29:38.867Z",
					"organisation_id": 1000,
					"type": "accounts",
					"version": 0
				}
			}
			`,
		})

		a := accountFunction{}
		response, err := a.FetchAccount("b80a0a46-7b68-472d-963e-1a7d2cc28a70")

		//validation
		if response != nil {
			t.Error("no record is expected")
		}

		if err == nil {
			t.Error("expected error")
		}

		if err.Error() != "json unmarshal error" {
			t.Error("invalid error message gotten")
		}
	})

	//Test record not found error
	t.Run("TestRecordNotExistFetchAccount", func(t *testing.T) {
		//enable mocking
		AddMock(Mock{
			Method:     http.MethodGet,
			Url:        "http://localhost:8080/v1/organisation/accounts/f8b9796e-c92a-4c36-96ff-a0e8cfc11f20",
			StatusCode: http.StatusNotFound,
			RequestBody: `
			{
				"error_message": "record f8b9796e-c92a-4c36-96ff-a0e8cfc11f20 does not exist"
			}
			`,
			Error: errors.New("record f8b9796e-c92a-4c36-96ff-a0e8cfc11f20 does not exist"),
		})

		a := accountFunction{}
		response, err := a.FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f20")

		//validation
		if err != nil {
			t.Error(fmt.Sprintf("no error expected, error received '%s'", err.Error()))
		}

		if response.Data == nil {
			t.Error("invalid error message gotten")
		}

		if err.Error() != "record f8b9796e-c92a-4c36-96ff-a0e8cfc11f20 does not exist" {
			t.Error("invalid error message gotten")
		}
	})

	//Test a successful fetch of record
	t.Run("TestSuccessfulFetchAccount", func(t *testing.T) {
		// enable mocking
		AddMock(Mock{
			Method:     http.MethodGet,
			Url:        "http://localhost:8080/v1/organisation/accounts/f8b9796e-c92a-4c36-96ff-a0e8cfc11f29",
			StatusCode: http.StatusOK,
			ResponseBody: `
			"data": {
				"attributes": {
					"alternative_names": null,
					"bank_id": "400400",
					"bank_id_code": "GBDSC",
					"base_currency": "GBP",
					"bic": "NWBKGB22",
					"country": "GB",
					"name": [
						"Samantha Holders"
					],
					"secondary_identification": "A1B2C3D4"
				},
				"created_on": "2021-11-21T20:29:38.867Z",
				"id": "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29",
				"modified_on": "2021-11-21T20:29:38.867Z",
				"organisation_id": "62b875ee-bd02-4b3b-a724-55cd924527e2",
				"type": "accounts",
				"version": 0
			}
			`,
		})

		a := accountFunction{}
		response, err := a.FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f29")

		//validation
		if err != nil {
			t.Error(fmt.Sprintf("no error expected, error received '%s'", err.Error()))
		}

		if response == nil {
			t.Error("invalid error message gotten")
		}

		fmt.Println("Check::", response)

		if response.Data.ID != "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29" {
			t.Error("invalid account retrived")
		}

	})
}

// func TestCreateAccount(t *testing.T) {

// 	t.Run("TestErrorSendingPostRequestToApi", func(t *testing.T) {

// 		a := accountFunction{}
// 		response, err := a.create()
// 	})

// 	t.Run("TestErrorUnmarshalCreateAccountResponseBody", func(t *testing.T) {
// 		a := accountFunction{}
// 		response, err := a.create()
// 	})

// 	t.Run("TestErrorUnmarshalCreateAccountResponseBody", func(t *testing.T) {
// 		a := accountFunction{}
// 		response, err := a.create()
// 	})
// }
