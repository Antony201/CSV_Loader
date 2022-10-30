package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test_task"
)

type TransactionPostgres struct {
	db *sqlx.DB
}


func NewTransaction(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{
		db: db,
	}
}

func (r *TransactionPostgres) Create(transaction test_task.Transaction) (int, error) {
	var transactionId int

	createTransactionQuery := fmt.Sprintf("INSERT INTO %s (transaction_id, " +
		"request_id, terminal_id, partner_object_id, amount_total, amount_original, commision_ps, commission_client," +
		"commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service," +
		"payee_id, payee_name, payee_bank_mfo, payee_bank_account, payment_narrative) VALUES ($1, $2, $3, $4, $5, $6, $7," +
		"$8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) RETURNING transaction_id", transactionsTable)

	row := r.db.QueryRow(createTransactionQuery, transaction.Transactionid, transaction.Requestid,
		transaction.Terminalid, transaction.PartnerObjectid, transaction.AmountTotal, transaction.AmountOriginal,
		transaction.CommissionPS, transaction.CommissionClient, transaction.CommissionProvider, transaction.DateInput,
		transaction.DatePost, transaction.Status, transaction.PaymentType, transaction.PaymentNumber,
		transaction.Serviceid, transaction.Service, transaction.Payeeid, transaction.PayeeName,
		transaction.PayeeBankMfo, transaction.PayeeBankAccount, transaction.PaymentNarrative)

	if err := row.Scan(&transactionId); err != nil {
		return 0, err
	}

	return transactionId, nil
}

func (r *TransactionPostgres) GetById(transactionId int) (test_task.Transaction, error) {
	var transaction test_task.Transaction

	query := fmt.Sprintf("SELECT transaction_id, request_id, terminal_id, partner_object_id, " +
		"amount_total, amount_original, commision_ps, commission_client, commission_provider, date_input, " +
		"date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, " +
		"payee_bank_mfo, payee_bank_account, payment_narrative FROM %s WHERE transaction_id=$1",
		transactionsTable)

	err := r.db.Get(&transaction, query, transactionId)

	return transaction, err
}

func (r *TransactionPostgres) GetByTerminalId(terminalId int) (test_task.Transaction, error) {
	var transaction test_task.Transaction

	query := fmt.Sprintf("SELECT transaction_id, request_id, terminal_id, partner_object_id, " +
		"amount_total, amount_original, commision_ps, commission_client, commission_provider, date_input, " +
		"date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, " +
		"payee_bank_mfo, payee_bank_account, payment_narrative FROM %s WHERE terminal_id=$1",
		transactionsTable)

	err := r.db.Get(&transaction, query, terminalId)

	return transaction, err
}