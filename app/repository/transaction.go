package repository

import (
	"bytes"
	"time"

	"github.com/niltonkummer/desafio-conductor/app/model"
	"github.com/niltonkummer/desafio-conductor/app/pdf"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction interface {
	GetTransactions(accountID uuid.UUID) (model.Transactions, error)
	GeneratePDF(accountID uuid.UUID) ([]byte, error)
}

type transactionRepo struct {
	DB *gorm.DB
}

func (t *transactionRepo) GetTransactions(accountID uuid.UUID) (transactions model.Transactions, err error) {
	tx := t.DB.
		Where("conta_id = ?", accountID).
		Find(&transactions)

	return transactions, tx.Error
}

func (t *transactionRepo) GeneratePDF(accountID uuid.UUID) ([]byte, error) {
	transactions, err := t.GetTransactions(accountID)
	if err != nil {
		return nil, err
	}

	var list pdf.Rows
	for _, tx := range transactions {
		list = append(list, tx)
	}

	var writer bytes.Buffer
	err = pdf.GeneratePDF(&writer, time.Now(), time.UTC, list)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func CreateTransactionRepository(db *gorm.DB) Transaction {
	return &transactionRepo{
		DB: db,
	}
}
