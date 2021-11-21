package main

import (
	"fmt"

	"github.com/gbubemismith/interview-accountapi-/accounts"
	"github.com/google/uuid"
)

func main() {
	//create new account data
	accountData := accounts.AccountData{
		Data: accounts.Data{
			ID:             uuid.NewString(),
			Type:           "accounts",
			Version:        1,
			OrganisationID: uuid.NewString(),
			Attributes: accounts.AccountAttributes{
				Country:      "GB",
				BaseCurrency: "GBP",
				BankID:       "400400",
				BankIDCode:   "GBDSC",
				Bic:          "NWBKGB22",
				Name: []string{
					"Gbubemi Smith",
				},
				AlternativeNames: []string{
					"Gbubemi Smith",
				},
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}

	accountFuntions := accounts.CreateAccountFuntions()

	result, err := accountFuntions.CreateAccount(&accountData)
	fmt.Println(result.Data)
	fmt.Println(err)

}
