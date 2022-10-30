package service

import (
	"test_task"
	"test_task/pkg/repository"
)

type TransactionsService struct {
	repo repository.Transaction
}

func NewTransactionsService(repo repository.Transaction) *TransactionsService {
	return &TransactionsService{repo: repo}
}

func (s *TransactionsService) Create(transactions []test_task.Transaction) (int, error) {

	for _, transaction := range transactions {
		_, err := s.repo.Create(transaction) // creating transaction in db

		if err != nil {
			return 0, err
		}
	}

	return len(transactions), nil
}

func (s *TransactionsService) GetByTransactionId(transactionId int) (test_task.Transaction, error) {
	return s.repo.GetById(transactionId)
}

func (s *TransactionsService) GetByTerminalIds(terminalIdParams []int) ([]test_task.Transaction, error) {
	resultTransactions := make([]test_task.Transaction, 0)

	for _, terminalId := range terminalIdParams {
		transaction, err := s.repo.GetByTerminalId(terminalId)
		if err != nil {
			continue
		}

		resultTransactions = append(resultTransactions, transaction)
	}
	return resultTransactions, nil
}

func (s *TransactionsService) GetByStatus(statusParam string) ([]test_task.Transaction, error) {
	return s.repo.GetByStatus(statusParam)
}

func (s *TransactionsService) GetByPaymentType(paymentTypeParam string) ([]test_task.Transaction, error) {
	return s.repo.GetByPaymentType(paymentTypeParam)
}

func (s *TransactionsService) GetByDatePeriod(fromDateParam, toDateParam string) ([]test_task.Transaction, error) {
	return s.repo.GetByDatePeriod(fromDateParam, toDateParam)
}

func (s *TransactionsService) GetByPaymentNarrative(paymentNarrativeParam string) ([]test_task.Transaction, error) {
	return s.repo.GetByPaymentNarrative(paymentNarrativeParam)
}