package repository

import (
	"github.com/jmoiron/sqlx"
	loader "github.com/Antony201/CsvLoader"
)

type Transaction interface {
	Create(transaction loader.Transaction) (int, error)
	GetById(transactionId int) (loader.Transaction, error)
	GetByTerminalId(terminalId int) (loader.Transaction, error)
	GetByStatus(statusParam string) ([]loader.Transaction, error)
	GetByPaymentType(paymentTypeParam string) ([]loader.Transaction, error)
	GetByDatePeriod(fromDateParam, toDateParam string) ([]loader.Transaction, error)
	GetByPaymentNarrative(paymentNarrativeParam string) ([]loader.Transaction, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransaction(db),
	}
}