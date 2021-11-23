package accounts

import (
	"github.com/google/uuid"
)

//priavte struct with priavte fields
type accountFunction struct{}

//public interface
type AccountFunc interface {
	CreateAccount(organisationId string, country string, bankId string, bic string, bankId_code string, accountName string) (*AccountData, error)
	FetchAccount(accountId string) (*AccountData, error)
	DeleteAccount(accountId string, version int64) error
}

//Initialize
//returns accountFunction struct which implements AccountFunc interface
func Init() AccountFunc {
	return &accountFunction{}
}

//Concrete implementations that can be accessed publicly
//Create account method just creates an account necessary fields that should be provided create a normal account or simple account
//organisationId is a unique uuid that should be provided
//country ISO 3166-1 code used to identify the domicile of the account e.g 'GB', 'FR'
//bankid Local country bank identifier
//bic SWIFT BIC in either 8 or 11 character format, nil if not required
//bankId_code Identifies the type of bank ID being used
//accountName Name of the account holder
func (a *accountFunction) CreateAccount(organisationId string, country string, bankId string, bic string, bankId_code string, accountName string) (*AccountData, error) {
	//create unique id
	id := uuid.NewString()
	//build account attributes
	attr := AccountAttributes{
		BankID:     bankId,
		Bic:        bic,
		Country:    &country,
		BankIDCode: bankId_code,
		Name:       []string{accountName},
	}
	//build data struct
	data := Data{
		Attributes:     &attr,
		ID:             id,
		OrganisationID: organisationId,
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

//Concrete implementation of delete account
//accountId ID of a account to delete
//version Version number of record
func (a *accountFunction) DeleteAccount(accountId string, version int64) error {
	return a.deleteResource(accountId, version)
}
