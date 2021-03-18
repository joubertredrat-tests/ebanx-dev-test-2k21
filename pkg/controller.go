package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleReset(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
