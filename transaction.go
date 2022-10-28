package test_task


type Transaction struct {
	Transactionid		int		`json:"transactionid"`
	Requestid			int		`json:"requestid"`
	Terminalid			int 	`json:"terminalid"`
	PartnerObjectid		int		`json:"partner_objectid"`
	AmountTotal			int		`json:"amount_total"`
	AmountOriginal		int		`json:"amount_original"`
	CommissionPS		int		`json:"commission_ps"`
	CommissionClient	int		`json:"commission_client"`
	CommissionProvider	int		`json:"commission_provider"`
	DateInput			string	`json:"date_input"`
	DatePost			string	`json:"date_post"`
	Status				string	`json:"status"`
	PaymentType			string	`json:"payment_type"`
	PaymentNumber		string	`json:"payment_number"`
	Serviceid			int		`json:"serviceid"`
	Service  			string	`json:"service"`
	Payeeid				int		`json:"payeeid"`
	PayeeName			string	`json:"payee_name"`
	PayeeBankMfo		int		`json:"payee_bank_mfo"`
	PayeeBankAccount	string	`json:"payee_bank_account"`
	PaymentNarrative	string	`json:"payment_narrative"`
}