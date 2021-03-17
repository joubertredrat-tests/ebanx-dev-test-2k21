package entity

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

func (a *Account) IncreaseAmount(BalanceAmount BalanceAmount) {
	a.BalanceAmount.Value = a.BalanceAmount.Value + BalanceAmount.Value
}

func (a *Account) DecreaseAmount(BalanceAmount BalanceAmount) {
	a.BalanceAmount.Value = a.BalanceAmount.Value - BalanceAmount.Value
}
