package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	container *Container
}

func NewController(container *Container) *Controller {
	return &Controller{
		container: container,
	}
}

func (c *Controller) handleReset(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func (c *Controller) handleBalance(ctx *gin.Context) {
	accountID := ctx.Query("account_id")
	if accountID == "" {
		ctx.String(http.StatusBadRequest, "account_id querystring is required")
		return
	}

	accountBalance, err := c.container.accountService.GetAccountBalance(accountID)

	if err != nil {
		ctx.String(http.StatusNotFound, "0")
	}

	ctx.String(http.StatusOK, strconv.FormatUint(uint64(*accountBalance), 10))
}

func (c *Controller) handleEvent(ctx *gin.Context) {
	data, _ := ctx.GetRawData()

	var eventRequest EventRequest

	if err := json.Unmarshal(data, &eventRequest); err != nil {
		ctx.String(http.StatusInternalServerError, "anything wrong is not right")
		return
	}

	if eventRequest.isDeposit() {
		c.handleEventDeposit(ctx)
		return
	}

	if eventRequest.isWithdraw() {
		c.handleEventWithdraw(ctx)
		return
	}

	if eventRequest.isTransfer() {
		c.handleEventTransfer(ctx)
		return
	}

	ctx.String(http.StatusBadRequest, "unknown event type")
	return
}

func (c *Controller) handleEventDeposit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "deposit")
}

func (c *Controller) handleEventWithdraw(ctx *gin.Context) {
	ctx.String(http.StatusOK, "withdraw")
}

func (c *Controller) handleEventTransfer(ctx *gin.Context) {
	ctx.String(http.StatusOK, "transfer")
}
