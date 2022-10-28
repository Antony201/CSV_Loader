package service

import "test_task/pkg/repository"

type TransactionList interface {

}

type TransactionItem interface {

}


type Service struct {
	TransactionList
	TransactionItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}