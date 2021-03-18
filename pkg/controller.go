package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/entity"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/service"
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
		return
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
		var depositRequest EventDepositRequest
		if err := json.Unmarshal(data, &depositRequest); err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		account, err := c.container.accountService.MakeDeposit(
			depositRequest.Destination,
			entity.BalanceAmount{
				Value: uint(depositRequest.Amount),
			},
		)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
		}

		eventDepositResponse := EventDepositResponse{
			Destination: AccountResponse{
				ID:      account.AccountID,
				Balance: account.BalanceAmount.Value,
			},
		}
		response, err := json.Marshal(eventDepositResponse)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		ctx.String(http.StatusCreated, string(response))
		return
	}

	if eventRequest.isWithdraw() {
		var withdrawRequest EventWithdrawRequest
		if err := json.Unmarshal(data, &withdrawRequest); err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		account, err := c.container.accountService.MakeWithdraw(
			withdrawRequest.Origin,
			entity.BalanceAmount{
				Value: uint(withdrawRequest.Amount),
			},
		)
		if err != nil {
			if errors.Is(err, service.ErrAccountNotFound) {
				ctx.String(http.StatusNotFound, "0")
				return
			}
			if errors.Is(err, entity.ErrInsufficientFunds) {
				ctx.String(http.StatusUnprocessableEntity, err.Error())
				return
			}

			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		eventWithdrawResponse := EventtWithdrawResponse{
			Origin: AccountResponse{
				ID:      account.AccountID,
				Balance: account.BalanceAmount.Value,
			},
		}
		response, err := json.Marshal(eventWithdrawResponse)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		ctx.String(http.StatusCreated, string(response))
		return
	}

	if eventRequest.isTransfer() {
		var transferRequest EventTransferRequest
		if err := json.Unmarshal(data, &transferRequest); err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		accountOrigin, accountDestination, err := c.container.accountService.MakeTransfer(
			transferRequest.Origin,
			transferRequest.Destination,
			entity.BalanceAmount{
				Value: uint(transferRequest.Amount),
			},
		)
		if err != nil {
			if errors.Is(err, service.ErrAccountOriginNotFound) {
				ctx.String(http.StatusNotFound, "0")
				return
			}
			if errors.Is(err, service.ErrAccountDestinationNotFound) {
				ctx.String(http.StatusNotFound, "0")
				return
			}
			if errors.Is(err, entity.ErrInsufficientFunds) {
				ctx.String(http.StatusUnprocessableEntity, err.Error())
				return
			}

			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		eventTransferResponse := EventTransferResponse{
			Origin: AccountResponse{
				ID:      accountOrigin.AccountID,
				Balance: accountOrigin.BalanceAmount.Value,
			},
			Destination: AccountResponse{
				ID:      accountDestination.AccountID,
				Balance: accountDestination.BalanceAmount.Value,
			},
		}
		response, err := json.Marshal(eventTransferResponse)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "anything wrong is not right")
			return
		}

		ctx.String(http.StatusCreated, string(response))
		return
	}

	ctx.String(http.StatusBadRequest, "unknown event type")
	return
}
