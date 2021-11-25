package main

import (
	"fmt"
	"os"

	"github.com/gbubemismith/interview-accountapi-/accounts"
)

func main() {
	//Examples on how to use the account function library
	//We have 3 functions available; create account, fetch account, delete account
	// Create Account: Account id can be nil if you want the system to generate one.
	//the other parameters for the create account are required
	//Fetch Account: The valid account id must be passed in order to retrive account details
	//Delete Account: The valid account id must be passed in order to delete an account

	//get base address
	baseAddress := os.Getenv("SERVER_URL")

	//if variable is not set on the environment or run the project locally
	if baseAddress == "" {
		baseAddress = "http://localhost:8080"
	}

	fmt.Println("BaseAddress", baseAddress)

	//Initialize account functions
	accountFuntions := accounts.Init(baseAddress)

	//create account
	result, err := accountFuntions.CreateAccount(nil, "f8b9796e-c92a-4c36-96ff-a0e8cfc11f29", "GB", "400400", "NWBKGB22", "GBDSC", "Oritse Smith")

	//fetch account
	// result, err := accountFuntions.FetchAccount("")

	//delete account
	// err := accountFuntions.DeleteAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*result)

}
