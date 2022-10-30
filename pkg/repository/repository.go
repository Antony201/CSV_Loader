package repository

import (
	"github.com/jmoiron/sqlx"
	"test_task"
)

type Transaction interface {
	Create(transaction test_task.Transaction) (int, error)
	GetById(transactionId int) (test_task.Transaction, error)
	GetByTerminalId(terminalId int) (test_task.Transaction, error)
	GetByStatus(statusParam string) ([]test_task.Transaction, error)
	GetByPaymentType(paymentTypeParam string) ([]test_task.Transaction, error)
	GetByDatePeriod(fromDateParam, toDateParam string) ([]test_task.Transaction, error)
	GetByPaymentNarrative(paymentNarrativeParam string) ([]test_task.Transaction, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransaction(db),
	}
}