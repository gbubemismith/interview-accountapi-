package accounts

type accountFunc struct{}

type AccountFunc interface {
	CreateAccount(accountData *AccountData) (*AccountData, error)
	FetchAccount()
	DeleteAccount()
}

//Concrete implementations that can be accessed publicly
func (a *accountFunc) CreateAccount(accountData *AccountData) (*AccountData, error) {
	return a.create(accountData)
}

func (a *accountFunc) FetchAccount() {

}

func (a *accountFunc) DeleteAccount() {

}
