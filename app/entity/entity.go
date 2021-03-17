package entity

type Amount struct {
	Value uint `json:"amount_value"`
}

type Account struct {
	AccountID string `json:"account_id"`
	Amount    Amount
}

func (a Account) ID() (jsonField string, value interface{}) {
	value = a.AccountID
	jsonField = "account_id"
	return
}

func (a *Account) IncreaseAmount(Amount Amount) {
	a.Amount.Value = a.Amount.Value + Amount.Value
}

func (a *Account) DecreaseAmount(Amount Amount) {
	a.Amount.Value = a.Amount.Value - Amount.Value
}
