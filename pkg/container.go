package pkg

import (
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/repository"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/service"
	"github.com/sonyarouje/simdb/db"
)

type Container struct {
	accountService    service.AccountServiceInterface
	accountRepository repository.AccountRepositoryInterface
	db                *db.Driver
}

func NewContainer(
	accountService *service.AccountService,
	accountRepository *repository.AccountRepository,
	db *db.Driver,
) *Container {
	return &Container{
		accountService:    accountService,
		accountRepository: accountRepository,
		db:                db,
	}
}
