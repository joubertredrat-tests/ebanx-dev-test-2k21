package entity

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds on account balance amount")
)

type BalanceAmount struct {
	Value uint `json:"balance_amount"`
}

type Account struct {
	AccountID     string `json:"account_id"`
	BalanceAmount BalanceAmount
}

func (a Account) ID() (jsonField string, value interface{}) {
	value = a.AccountID
	jsonField = "account_id"
	return
}

func (a *Account) IncreaseBalanceAmount(BalanceAmount BalanceAmount) {
	a.BalanceAmount.Value = a.BalanceAmount.Value + BalanceAmount.Value
}

func (a *Account) DecreaseBalanceAmount(BalanceAmount BalanceAmount) error {
	if BalanceAmount.Value > a.BalanceAmount.Value {
		return ErrInsufficientFunds
	}

	a.BalanceAmount.Value = a.BalanceAmount.Value - BalanceAmount.Value
	return nil
}
