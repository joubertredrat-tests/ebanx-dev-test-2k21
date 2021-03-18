package pkg

const (
	EVENT_TYPE_DEPOSIT  = "deposit"
	EVENT_TYPE_WITHDRAW = "withdraw"
	EVENT_TYPE_TRANSFER = "transfer"
)

type EventRequest struct {
	Type string `json:"type"`
}

func (f *EventRequest) isDeposit() bool {
	return EVENT_TYPE_DEPOSIT == f.Type
}

func (f *EventRequest) isWithdraw() bool {
	return EVENT_TYPE_WITHDRAW == f.Type
}

func (f *EventRequest) isTransfer() bool {
	return EVENT_TYPE_TRANSFER == f.Type
}

type AccountResponse struct {
	ID      string `json:"id"`
	Balance uint   `json:"balance"`
}

type EventDepositRequest struct {
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
}

type EventDepositResponse struct {
	Destination AccountResponse `json:"destination"`
}

type EventWithdrawRequest struct {
	Origin string `json:"origin"`
	Amount int    `json:"amount"`
}

type EventtWithdrawResponse struct {
	Origin AccountResponse `json:"origin"`
}

type EventTransferRequest struct {
	Origin      string `json:"origin"`
	Amount      int    `json:"amount"`
	Destination string `json:"destination"`
}

type EventTransferResponse struct {
	Origin      AccountResponse `json:"origin"`
	Destination AccountResponse `json:"destination"`
}
