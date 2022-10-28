package repository

import "github.com/jmoiron/sqlx"

type TransactionList interface {

}

type TransactionItem interface {

}

type Repository struct {
	TransactionList
	TransactionItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}