package service

import (
	"test_task"
	"test_task/pkg/repository"
)


type Transactions interface {
	Create(transactions []test_task.Transaction) (int, error)

	GetByTransactionId(transactionId int) (test_task.Transaction, error)
	GetByTerminalIds(terminalIdParams []int) ([]test_task.Transaction, error)
	GetByStatus(statusParam string) ([]test_task.Transaction, error)
	GetByPaymentType(paymentTypeParam string) ([]test_task.Transaction, error)
	GetByDatePeriod(fromDateParam, toDateParam string) ([]test_task.Transaction, error)
	GetByPaymentNarrative(paymentNarrativeParam string) ([]test_task.Transaction, error)
}


type Service struct {
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transactions: NewTransactionsService(repos.Transaction),
	}
}