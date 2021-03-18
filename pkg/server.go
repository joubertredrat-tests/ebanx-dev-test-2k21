package pkg

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.Default()

	router.POST("/reset", handleReset)
	router.Run(":8000")
}
