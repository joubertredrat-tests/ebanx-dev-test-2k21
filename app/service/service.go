package service

import (
	"errors"

	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/entity"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/repository"
)

var (
	ErrMakeDepositAccount = errors.New("error on make deposit in account")
)

type AccountServiceInterface interface {
	MakeDeposit(AccountID string, Amount entity.Amount) (*entity.Account, error)
}

type AccountService struct {
	repo repository.AccountRepositoryInterface
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) MakeDeposit(AccountID string, Amount entity.Amount) (*entity.Account, error) {
	account, _ := s.repo.GetByAccountID(AccountID)
	if account == nil {
		return s.makeDepositNewAccount(AccountID, Amount)
	}

	account.IncreaseAmount(Amount)

	err := s.repo.Update(account)
	if err != nil {
		return nil, ErrMakeDepositAccount
	}
	return account, nil
}

func (s *AccountService) makeDepositNewAccount(AccountID string, Amount entity.Amount) (*entity.Account, error) {
	account := entity.Account{
		AccountID: AccountID,
		Amount:    Amount,
	}
	err := s.repo.Insert(account)
	if err != nil {
		return nil, ErrMakeDepositAccount
	}
	return &account, nil
}
