package account

import "fmt"

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AccountRepository interface {
	SaveAccount(account Account)
	GetAccounts()
}

type AccountService interface {
	SaveAccount(account Account)
	GetAccounts()
}

// String returns a string representation from this struct
func (c Account) String() string {
	return fmt.Sprintf("[%d, %d]", c.Id, c.Name)
}
