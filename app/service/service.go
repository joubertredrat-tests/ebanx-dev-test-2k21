package service

import (
	"errors"

	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/entity"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/repository"
)

var (
	ErrHouston            = errors.New("anything wrong is not right")
	ErrAccountNotFound    = errors.New("account not found")
	ErrMakeDepositAccount = errors.New("error on make deposit in account")
)

type AccountServiceInterface interface {
	GetAccountBalance(AccountID string) (*uint, error)
	MakeDeposit(AccountID string, Amount entity.BalanceAmount) (*entity.Account, error)
}

type AccountService struct {
	repo repository.AccountRepositoryInterface
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) GetAccountBalance(AccountID string) (*uint, error) {
	account, err := s.repo.GetByAccountID(AccountID)
	if err != nil {
		if errors.Is(err, repository.ErrAccountNotFound) {
			return nil, ErrAccountNotFound
		}

		return nil, ErrHouston
	}

	return &account.BalanceAmount.Value, nil
}

func (s *AccountService) MakeDeposit(AccountID string, BalanceAmount entity.BalanceAmount) (*entity.Account, error) {
	account, _ := s.repo.GetByAccountID(AccountID)
	if account == nil {
		return s.makeDepositNewAccount(AccountID, BalanceAmount)
	}

	account.IncreaseAmount(BalanceAmount)

	err := s.repo.Update(account)
	if err != nil {
		return nil, ErrMakeDepositAccount
	}
	return account, nil
}

func (s *AccountService) makeDepositNewAccount(AccountID string, BalanceAmount entity.BalanceAmount) (*entity.Account, error) {
	account := entity.Account{
		AccountID:     AccountID,
		BalanceAmount: BalanceAmount,
	}
	err := s.repo.Insert(account)
	if err != nil {
		return nil, ErrMakeDepositAccount
	}
	return &account, nil
}
