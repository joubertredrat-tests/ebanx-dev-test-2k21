package entity

type Amount struct {
	value uint
}

type Account struct {
	id     string
	amount Amount
}
