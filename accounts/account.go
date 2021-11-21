package accounts

import (
	"github.com/google/uuid"
)

//priavte struct with priavte fields
type accountFunction struct {
	organisationId string
	country        string
	bankId         string
	bic            string
	bankId_code    string
	accountName    string
}

//public interface
type AccountFunc interface {
	CreateAccount() (*AccountData, error)
	FetchAccount(accountId string) (*AccountData, error)
	DeleteAccount(test string)
}

//organisationId is a unique uuid that should be provided
//country ISO 3166-1 code used to identify the domicile of the account e.g 'GB', 'FR'
//bankid Local country bank identifier
//bic SWIFT BIC in either 8 or 11 character format, nil if not required
//bankId_code Identifies the type of bank ID being used
//accountName Name of the account holder
func CreateAccountFuntions(organisationId string, country string, bankId string, bic string, bankId_code string, accountName string) AccountFunc {
	return &accountFunction{
		organisationId: organisationId,
		country:        country,
		bankId:         bankId,
		bic:            bic,
		bankId_code:    bankId_code,
		accountName:    accountName,
	}
}

//Concrete implementations that can be accessed publicly
//Create account method just creates an account necessary fields that should be provided create a normal account or simple account
func (a *accountFunction) CreateAccount() (*AccountData, error) {
	//create unique id
	id := uuid.NewString()
	//build account attributes
	attr := AccountAttributes{
		BankID:     a.bankId,
		Bic:        a.bic,
		Country:    &a.country,
		BankIDCode: a.bankId_code,
		Name:       []string{a.accountName},
	}
	//build data struct
	data := Data{
		Attributes:     &attr,
		ID:             id,
		OrganisationID: a.organisationId,
		Type:           "accounts",
	}

	accountData := AccountData{
		Data: &data,
	}

	return a.create(accountData)
}

//Concrete implementation of fetch account
//accountId ID of the Account resource to fetch
func (a *accountFunction) FetchAccount(accountId string) (*AccountData, error) {
	return a.getAccount(accountId)
}

func (a *accountFunction) DeleteAccount(test string) {

}
