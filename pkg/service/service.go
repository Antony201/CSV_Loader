package service

import (
	loader "github.com/Antony201/CsvLoader"
	"github.com/Antony201/CsvLoader/pkg/repository"
	"mime/multipart"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Transactions interface {
	LoadFileToDb(file multipart.File)

	GetByTransactionId(transactionId int) (loader.Transaction, error)
	GetByTerminalIds(terminalIdParams []int) ([]loader.Transaction, error)
	GetByStatus(statusParam string) ([]loader.Transaction, error)
	GetByPaymentType(paymentTypeParam string) ([]loader.Transaction, error)
	GetByDatePeriod(fromDateParam, toDateParam string) ([]loader.Transaction, error)
	GetByPaymentNarrative(paymentNarrativeParam string) ([]loader.Transaction, error)
}


type Service struct {
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transactions: NewTransactionsService(repos.Transaction),
	}
}