package accounts

const Uri = "/v1/organisation/accounts"

type accountFunc struct{}

type AccountFunc interface {
	CreateAccount()
	FetchAccount()
	DeleteAccount()
}

func (a *accountFunc) CreateAccount() {

}

func (a *accountFunc) FetchAccount() {

}

func (a *accountFunc) DeleteAccount() {

}
