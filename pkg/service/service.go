package service

import (
	"test_task"
	"test_task/pkg/repository"
)


type Transactions interface {
	Create(transactions []test_task.Transaction) (int, error) // sending transaction id and error(maybe)
	GetByTransactionId(transactionId int) (test_task.Transaction, error)

}


type Service struct {
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transactions: NewTransactionsService(repos.Transaction),
	}
}