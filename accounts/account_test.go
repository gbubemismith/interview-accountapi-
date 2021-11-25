package accounts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test Fetch Account
func TestFetchAccount(t *testing.T) {

	//Test Invalid Id
	t.Run("TestInavlidIdErrorFetchAccount", func(t *testing.T) {
		f, err := Init().FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f")

		assert.NotNil(t, err)
		assert.Nil(t, f)
	})

	//connection refused

	//Test JSON unmarshal error
	t.Run("TestErrorUnmarshalFetchAccountResponseBody", func(t *testing.T) {
		f, err := Init().FetchAccount("")

		assert.Nil(t, f)
		assert.Contains(t, err.Error(), "cannot unmarshal")
	})

	//Test record not found error
	t.Run("TestRecordNotExistFetchAccount", func(t *testing.T) {
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f20"
		f, err := Init().FetchAccount(id)
		assert.Nil(t, f)
		assert.EqualError(t, err, fmt.Sprintf("record %s does not exist", id))
	})

	//Test a successful fetch of record
	t.Run("TestSuccessfulFetchAccount", func(t *testing.T) {
		a, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		f, err := Init().FetchAccount(a.Data.ID)

		assert.Nil(t, err)
		assert.NotNil(t, f)
		assert.Equal(t, a.Data.ID, f.Data.ID)

		//clean up
		Init().DeleteAccount(a.Data.ID, 0)
	})
}

//Test Create Account
func TestCreateAccount(t *testing.T) {
	//Test invalid data passed to create account
	t.Run("TestInvalidPostDataCreateAccount", func(t *testing.T) {
		c, err := Init().CreateAccount(nil, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "")

		assert.Nil(t, c)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "validation failure")
	})

	//Test account creation with already used id
	t.Run("TestDuplicateIDCreateAccount", func(t *testing.T) {
		//ensure account does not exist
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		Init().DeleteAccount(id, 0)
		//create mock account
		_, err := mockCreateAccount()
		if err != nil {
			t.Error("error creating mock account")
		}

		c, err := Init().CreateAccount(&id, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User 2")

		assert.Nil(t, c)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Account cannot be created as it violates a duplicate constraint")

		//clean up
		Init().DeleteAccount(id, 0)
	})

	t.Run("TestSuccessfulCreateAccount", func(t *testing.T) {
		//ensure account does not exist
		id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
		Init().DeleteAccount(id, 0)

		c, err := Init().CreateAccount(&id, "72b975ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User 2")

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, id, c.Data.ID)

		Init().DeleteAccount(id, 0)
	})
}

//Test Delete Account
func TestDeleteAccount(t *testing.T) {
	//Test account does not exist
	t.Run("TestAccountIDDoesNotExist", func(t *testing.T) {
		id := "106e3db6-9ffa-4fb0-a649-210b324143c4"
		err := Init().DeleteAccount(id, 0)

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

		e := Init().DeleteAccount(id, 1)

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

		e := Init().DeleteAccount(id, 0)

		assert.Nil(t, e)
	})
}

func mockCreateAccount() (*AccountData, error) {
	//creating mock account for testing purposes
	//id -> f8b9796e-c92a-4c36-96ff-a0e8cfc11f29
	//organisation id -> 62b875ee-bd02-4b3b-a724-55cd924527e2

	id := "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29"
	Init().DeleteAccount(id, 0)

	acc, err := Init().CreateAccount(&id, "62b875ee-bd02-4b3b-a724-55cd924527e2", "GB", "400400", "NWBKGB22", "GBDSC", "Test User")

	return acc, err
}
