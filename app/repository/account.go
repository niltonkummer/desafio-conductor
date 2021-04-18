package repository

import (
	"github.com/niltonkummer/desafio-conductor/app/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Account interface {
	GetAccount(id uuid.UUID) (*model.Account, error)
	ListAccounts() (model.Accounts, error)
}

type accountRepo struct {
	DB *gorm.DB
}

func (a *accountRepo) GetAccount(id uuid.UUID) (account *model.Account, err error) {

	tx := a.DB.First(&account, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return
}

func (a *accountRepo) ListAccounts() (accounts model.Accounts, err error) {

	tx := a.DB.Find(&accounts)

	return accounts, tx.Error

}

func CreateAccountRepository(db *gorm.DB) Account {
	return &accountRepo{
		DB: db,
	}
}
