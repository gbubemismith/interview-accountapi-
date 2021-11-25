package accounts

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockBaseAddress = "http://localhost:3000"
	baseAddress     = os.Getenv("SERVER_URL")
	// baseAddress = "http://localhost:8080"
)

//Test Fetch Account
func TestFetchAccount(t *testing.T) {
	//Test Network Reaching endpoint
	t.Run("TestConnectionRefusedToFetchAccount", func(t *testing.T) {
		f, err := Init(mockBaseAddress).FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f")

		assert.NotNil(t, err)
		assert.Nil(t, f)
		assert.Contains(t, err.Error(), "connection refused")
	})

	//Test Invalid Id
	t.Run("TestInavlidIdErrorFetchAccount", func(t *testing.T) {
		f, err := Init(baseAddress).FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f")

		assert.NotNil(t, err)
		assert.Nil(t, f)
	})

	//Test Empty AccountId
	t.Run("TestEmptyAccountIdFetchAccount", func(t *testing.T) {
		f, err := Init(baseAddress).FetchAccount("")

		assert.NotNil(t, err)
		assert.Nil(t, f)
		assert.EqualError(t, err, "accountid must be provided")
	})

	//Test record not found error
	t.Run("TestRecordNotExistFetchAccount", func(t *testing.T) {
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f20"
		f, err := Init(baseAddress).FetchAccount(id)
		assert.Nil(t, f)
		assert.EqualError(t, err, fmt.Sprintf("record %s does not exist", id))
	})

	//Test a successful fetch of record
	t.Run("TestSuccessfulFetchAccount", func(t *testing.T) {
		a, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		f, err := Init(baseAddress).FetchAccount(a.Data.ID)

		assert.Nil(t, err)
		assert.NotNil(t, f)
		assert.Equal(t, a.Data.ID, f.Data.ID)

		//clean up
		Init(baseAddress).DeleteAccount(a.Data.ID, 0)
	})
}

//Test Create Account
func TestCreateAccount(t *testing.T) {

	//Test Network Reaching endpoint
	t.Run("TestConnectionRefusedToCreateAccount", func(t *testing.T) {
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		f, err := Init(mockBaseAddress).CreateAccount(&id, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User 2")

		assert.NotNil(t, err)
		assert.Nil(t, f)
		assert.Contains(t, err.Error(), "connection refused")
	})

	//Test invalid data passed to create account
	t.Run("TestInvalidPostDataCreateAccount", func(t *testing.T) {
		c, err := Init(baseAddress).CreateAccount(nil, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "")

		assert.Nil(t, c)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "validation failure")
	})

	//Test account creation with already used id
	t.Run("TestDuplicateIDCreateAccount", func(t *testing.T) {
		//ensure account does not exist
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		Init(baseAddress).DeleteAccount(id, 0)
		//create mock account
		_, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		c, err := Init(baseAddress).CreateAccount(&id, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User 2")

		assert.Nil(t, c)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")

		//clean up
		Init(baseAddress).DeleteAccount(id, 0)
	})

	t.Run("TestSuccessfulCreateAccount", func(t *testing.T) {
		//ensure account does not exist
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		Init(baseAddress).DeleteAccount(id, 0)

		c, err := Init(baseAddress).CreateAccount(&id, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User 2")

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, id, c.Data.ID)

		Init(baseAddress).DeleteAccount(id, 0)
	})
}

//Test Delete Account
func TestDeleteAccount(t *testing.T) {
	//Test Network Reaching endpoint
	t.Run("TestConnectionRefusedToDeleteAccount", func(t *testing.T) {
		id := "106e3db6-9ffa-4fb0-a649-210b324143c4"
		err := Init(mockBaseAddress).DeleteAccount(id, 0)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "connection refused")
	})

	//Test account does not exist
	t.Run("TestAccountIDDoesNotExist", func(t *testing.T) {
		id := "106e3db6-9ffa-4fb0-a649-210b324143c4"
		err := Init(baseAddress).DeleteAccount(id, 0)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "specified resource does not exist")
	})

	t.Run("TestVersionIncorrect", func(t *testing.T) {
		//ensure account does not exist, and then create a new account
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		_, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		e := Init(baseAddress).DeleteAccount(id, 1)

		assert.NotNil(t, e)
		assert.EqualError(t, e, "specified version incorrect")
	})

	t.Run("TestSuccessfulDeleteAccount", func(t *testing.T) {
		//ensure account does not exist, and then create a new account
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		_, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		e := Init(baseAddress).DeleteAccount(id, 0)

		assert.Nil(t, e)
	})
}

func mockCreateAccount() (*AccountData, error) {
	//creating mock account for testing purposes
	//id -> f8b9796e-c92a-4c36-96ff-a0e8cfc11f29
	//organisation id -> 62b875ee-bd02-4b3b-a724-55cd924527e2

	id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
	Init(baseAddress).DeleteAccount(id, 0)

	acc, err := Init(baseAddress).CreateAccount(&id, "62b875ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User")

	return acc, err
}
