package main

import (
	"fmt"

	"github.com/gbubemismith/interview-accountapi-/accounts"
)

func main() {

	//Initialize account functions
	accountFuntions := accounts.Init()

	//create account
	result, err := accountFuntions.CreateAccount(nil, "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29", "GB", "400400", "NWBKGB22", "GBDSC", "Oritse Smith")

	//fetch account
	// result, err := accountFuntions.FetchAccount("f8b9796e-c92a-4c36-96ff-a0e8cfc11f29")

	//delete account
	// err := accountFuntions.DeleteAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)

	if err != nil {
		fmt.Println(err)
	}

	resp := *result.Data

	fmt.Println(*resp.Attributes)

}
