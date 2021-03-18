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
