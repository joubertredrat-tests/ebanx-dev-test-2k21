package repository

import (
	"errors"

	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/app/entity"
	"github.com/sonyarouje/simdb/db"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountInsert   = errors.New("error on insert account")
	ErrAccountUpdate   = errors.New("error on update account")
)

type AccountRepositoryInterface interface {
	GetByAccountID(AccountID string) (*entity.Account, error)
	Insert(account entity.Account) error
	Update(account *entity.Account) error
}

type AccountRepository struct {
	db *db.Driver
}

func NewAccountRepository(db *db.Driver) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetByAccountID(AccountID string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Open(entity.Account{}).Where("account_id", "=", AccountID).First().AsEntity(&account)
	if err != nil {
		return nil, ErrAccountNotFound
	}
	return &account, nil
}

func (r *AccountRepository) Insert(account entity.Account) error {
	err := r.db.Insert(account)
	if err != nil {
		return ErrAccountInsert
	}
	return nil
}

func (r *AccountRepository) Update(account *entity.Account) error {
	err := r.db.Update(account)
	if err != nil {
		return ErrAccountUpdate
	}
	return nil
}
