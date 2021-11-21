package accounts

type accountFunction struct{}

type AccountFunc interface {
	CreateAccount(accountData *AccountData) (*AccountData, error)
	FetchAccount(test string)
	DeleteAccount(test string)
}

func CreateAccountFuntions() *accountFunction {
	return &accountFunction{}
}

//Concrete implementations that can be accessed publicly
func (a *accountFunction) CreateAccount(accountData *AccountData) (*AccountData, error) {
	return a.create(accountData)
}

func (a *accountFunction) FetchAccount(test string) {

}

func (a *accountFunction) DeleteAccount(test string) {

}
