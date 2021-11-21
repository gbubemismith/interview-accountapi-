package main

import (
	"encoding/json"
	"fmt"

	"github.com/gbubemismith/interview-accountapi-/accounts"
)

func main() {
	//create new account data
	// accountData := accounts.AccountData{
	// 	Data: &accounts.Data{
	// 		ID:             uuid.NewString(),
	// 		Type:           "accounts",
	// 		Version:        1,
	// 		OrganisationID: uuid.NewString(),
	// 		Attributes: &accounts.AccountAttributes{
	// 			Country:      "GB",
	// 			BaseCurrency: "GBP",
	// 			BankID:       "400400",
	// 			BankIDCode:   "GBDSC",
	// 			Bic:          "NWBKGB22",
	// 			Name: []string{
	// 				"Gbubemi Smith",
	// 			},
	// 			AlternativeNames: []string{
	// 				"Gbubemi Smith",
	// 			},
	// 			AccountClassification:   "Personal",
	// 			JointAccount:            false,
	// 			AccountMatchingOptOut:   false,
	// 			SecondaryIdentification: "A1B2C3D4",
	// 		},
	// 	},
	// }

	//create account
	accountFuntions := accounts.CreateAccountFuntions("b80a0a46-7b68-472d-963e-1a7d2cc28a70", "GB", "400400", "NWBKGB22", "GBDSC", "Oritse Smith")

	// result, err := accountFuntions.CreateAccount()

	result, err := accountFuntions.FetchAccount("e1ec4c2f-a687-4391-934c-055980e0d77")

	if err != nil {
		fmt.Println(err)
	}

	y, _ := json.Marshal(result)
	fmt.Println(string(y))

}
