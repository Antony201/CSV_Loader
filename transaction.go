package test_task


type Transaction struct {
	Transactionid		int		`csv:"TransactionId" db:"transaction_id"`
	Requestid			int		`csv:"RequestId" db:"request_id"`
	Terminalid			int 	`csv:"TerminalId" db:"terminal_id"`
	PartnerObjectid		int		`csv:"PartnerObjectId" db:"partner_object_id"`
	AmountTotal			int		`csv:"AmountTotal" db:"amount_total"`
	AmountOriginal		int		`csv:"AmountOriginal" db:"amount_original"`
	CommissionPS		float32	`csv:"CommissionPS" db:"commision_ps"`
	CommissionClient	float32	`csv:"CommissionClient" db:"commission_client"`
	CommissionProvider	float32	`csv:"CommissionProvider" db:"commission_provider"`
	DateInput			string	`csv:"DateInput" db:"date_input"`
	DatePost			string	`csv:"DatePost" db:"date_post"`
	Status				string	`csv:"Status" db:"status"`
	PaymentType			string	`csv:"PaymentType" db:"payment_type"`
	PaymentNumber		string	`csv:"PaymentNumber" db:"payment_number"`
	Serviceid			int		`csv:"ServiceId" db:"service_id"`
	Service  			string	`csv:"Service" db:"service"`
	Payeeid				int		`csv:"PayeeId" db:"payee_id"`
	PayeeName			string	`csv:"PayeeName" db:"payee_name"`
	PayeeBankMfo		int		`csv:"PayeeBankMfo" db:"payee_bank_mfo"`
	PayeeBankAccount	string	`csv:"PayeeBankAccount" db:"payee_bank_account"`
	PaymentNarrative	string	`csv:"PaymentNarrative" db:"payment_narrative"`
}