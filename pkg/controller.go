package pkg

import (
	"net/http"

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

	ctx.String(http.StatusNotFound, "0")
}
