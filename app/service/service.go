package service

import (
	"errors"

	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/entity"
	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/repository"
)

var (
	ErrHouston                    = errors.New("anything wrong is not right")
	ErrAccountNotFound            = errors.New("account not found")
	ErrAccountOriginNotFound      = errors.New("origin account not found")
	ErrAccountDestinationNotFound = errors.New("destination account not found")
	ErrMakeDepositAccount         = errors.New("error on make deposit in account")
	ErrMakeWithdrawAccount        = errors.New("error on make withdraw in account")
	ErrMakeTransferAccounts       = errors.New("error on make tranfer between accounts")
)

type AccountServiceInterface interface {
	GetAccountBalance(AccountID string) (*uint, error)
	MakeDeposit(AccountID string, BalanceAmount entity.BalanceAmount) (*entity.Account, error)
	MakeWithdraw(AccountID string, BalanceAmount entity.BalanceAmount) (*entity.Account, error)
	MakeTransfer(AccountOriginID, AccountDestinationID string, BalanceAmount entity.BalanceAmount) (*entity.Account, *entity.Account, error)
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

	account.IncreaseBalanceAmount(BalanceAmount)

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

func (s *AccountService) MakeWithdraw(AccountID string, BalanceAmount entity.BalanceAmount) (*entity.Account, error) {
	account, _ := s.repo.GetByAccountID(AccountID)
	if account == nil {
		return nil, ErrAccountNotFound
	}

	if err := account.DecreaseBalanceAmount(BalanceAmount); err != nil {
		return nil, err
	}

	if err := s.repo.Update(account); err != nil {
		return nil, ErrMakeWithdrawAccount
	}

	return account, nil
}

func (s *AccountService) MakeTransfer(
	AccountOriginID, AccountDestinationID string,
	BalanceAmount entity.BalanceAmount,
) (*entity.Account, *entity.Account, error) {
	accountOrigin, _ := s.repo.GetByAccountID(AccountOriginID)
	if accountOrigin == nil {
		return nil, nil, ErrAccountOriginNotFound
	}
	accountDestination, _ := s.repo.GetByAccountID(AccountDestinationID)
	if accountDestination == nil {
		return nil, nil, ErrAccountDestinationNotFound
	}

	if err := accountOrigin.DecreaseBalanceAmount(BalanceAmount); err != nil {
		return nil, nil, err
	}
	accountDestination.IncreaseBalanceAmount(BalanceAmount)

	if err := s.repo.Update(accountOrigin); err != nil {
		return nil, nil, ErrMakeTransferAccounts
	}
	if err := s.repo.Update(accountDestination); err != nil {
		return nil, nil, ErrMakeTransferAccounts
	}

	return accountOrigin, accountDestination, nil
}
