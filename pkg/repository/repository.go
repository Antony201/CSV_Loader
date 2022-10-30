package repository

import (
	"github.com/jmoiron/sqlx"
	"test_task"
)

type Transaction interface {
	Create(transaction test_task.Transaction) (int, error)
	GetById(transactionId int) (test_task.Transaction, error)
	GetByTerminalId(terminalId int) (test_task.Transaction, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransaction(db),
	}
}