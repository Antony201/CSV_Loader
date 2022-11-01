package service

import (
	loader "github.com/Antony201/CsvLoader"
	"github.com/Antony201/CsvLoader/pkg/repository"
	"github.com/gocarina/gocsv"
	"mime/multipart"
)

type TransactionsService struct {
	repo repository.Transaction
}

func NewTransactionsService(repo repository.Transaction) *TransactionsService {
	return &TransactionsService{repo: repo}
}

func (s *TransactionsService) LoadFileToDb(file multipart.File) {
	// create chan for file chunks
	chunks_chan := make(chan loader.Transaction)

	go func() {
		err := gocsv.UnmarshalToChan(file, chunks_chan)
		if err != nil {
			return
		}
	}()


	for transaction := range chunks_chan {
		_, err := s.repo.Create(transaction)

		if err != nil {
			return
		}
	}

	defer file.Close()
}

func (s *TransactionsService) GetByTransactionId(transactionId int) (loader.Transaction, error) {
	return s.repo.GetById(transactionId)
}

func (s *TransactionsService) GetByTerminalIds(terminalIdParams []int) ([]loader.Transaction, error) {
	resultTransactions := make([]loader.Transaction, 0)

	for _, terminalId := range terminalIdParams {
		transaction, err := s.repo.GetByTerminalId(terminalId)
		if err != nil {
			continue
		}

		resultTransactions = append(resultTransactions, transaction)
	}
	return resultTransactions, nil
}

func (s *TransactionsService) GetByStatus(statusParam string) ([]loader.Transaction, error) {
	return s.repo.GetByStatus(statusParam)
}

func (s *TransactionsService) GetByPaymentType(paymentTypeParam string) ([]loader.Transaction, error) {
	return s.repo.GetByPaymentType(paymentTypeParam)
}

func (s *TransactionsService) GetByDatePeriod(fromDateParam, toDateParam string) ([]loader.Transaction, error) {
	return s.repo.GetByDatePeriod(fromDateParam, toDateParam)
}

func (s *TransactionsService) GetByPaymentNarrative(paymentNarrativeParam string) ([]loader.Transaction, error) {
	return s.repo.GetByPaymentNarrative(paymentNarrativeParam)
}