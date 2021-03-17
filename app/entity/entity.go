package entity

type Amount struct {
	Value uint `json:"amount_value"`
}

type Account struct {
	AccountID string `json:"account_id"`
	Amount    Amount
}

func (c Account) ID() (jsonField string, value interface{}) {
	value = c.AccountID
	jsonField = "account_id"
	return
}
