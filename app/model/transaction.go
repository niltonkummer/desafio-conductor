package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Model
	ContaID   *string  `json:"conta-id"`
	Conta     *Account `json:"-"`
	Descricao string   `json:"descricao"`
	Valor     float64  `json:"valor"`
}
type Transactions []Transaction

func (t Transaction) Description() string {
	return t.Descricao
}

func (t Transaction) Amount() decimal.Decimal {
	return decimal.NewFromFloat(t.Valor)
}

func (t Transaction) CreatedAt() time.Time {
	return t.Model.CreatedAt
}
