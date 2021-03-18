package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/repository"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/service"
	"github.com/sonyarouje/simdb/db"
)

func Run() {
	db, err := db.New("data")
	if err != nil {
		panic(err)
	}
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	container := NewContainer(accountService, accountRepository, db)

	controller := NewController(container)
	router := gin.Default()

	router.POST("/reset", controller.handleReset)
	router.GET("/balance", controller.handleBalance)
	router.POST("/event", controller.handleEvent)

	fmt.Println("Server running at http://0.0.0.0:8000")
	router.Run(":8000")
}
