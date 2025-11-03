package moneybird

import "net/http"

type FinancialStatement struct {
	ID                           string                                  `json:"id,omitempty"`
	FinancialAccountID           string                                  `json:"financial_account_id"`
	Reference                    string                                  `json:"reference"`
	OfficialDate                 string                                  `json:"official_date,omitempty"`
	OfficialBalance              string                                  `json:"official_balance,omitempty"`
	ImporterService              string                                  `json:"importer_service,omitempty"`
	FinancialMutations           []*FinancialMutation                    `json:"financial_mutations,omitempty"`
	ImporterKey                  string                                  `json:"importer_key,omitempty"`                   // used in request
	FinancialMutationsAttributes map[string]*FinancialMutationAttributes `json:"financial_mutations_attributes,omitempty"` // used in request
}

type FinancialMutationAttributes struct {
	Date                         string `json:"date,omitempty"`
	ValutationDate               string `json:"valutation_date,omitempty"`
	Message                      string `json:"message,omitempty"`
	Amount                       string `json:"amount,omitempty"`
	Code                         string `json:"code,omitempty"`
	ContraAccountName            string `json:"contra_account_name,omitempty"`
	ContraAccountNumber          string `json:"contra_account_number,omitempty"`
	BatchReference               string `json:"batch_reference,omitempty"`
	Offset                       string `json:"offset,omitempty"`
	AccountServicerTransactionId string `json:"account_servicer_transaction_id,omitempty"`
}

type FinancialMutation struct {
	ID                           string                                   `json:"id,omitempty"`
	AdministrationID             string                                   `json:"administration_id,omitempty"`
	Amount                       string                                   `json:"amount,omitempty"`
	Code                         string                                   `json:"code,omitempty"`
	Date                         string                                   `json:"date,omitempty"`
	Message                      string                                   `json:"message,omitempty"`
	ContraAccountName            string                                   `json:"contra_account_name,omitempty"`
	ContraAccountNumber          string                                   `json:"contra_account_number,omitempty"`
	State                        string                                   `json:"state,omitempty"`
	AmountOpen                   string                                   `json:"amount_open,omitempty"`
	SepaFields                   interface{}                              `json:"sepa_fields,omitempty"`
	BatchReference               string                                   `json:"batch_reference,omitempty"`
	FinancialAccountID           string                                   `json:"financial_account_id,omitempty"`
	Currency                     string                                   `json:"currency,omitempty"`
	OriginalAmount               string                                   `json:"original_amount,omitempty"`
	CreatedAt                    string                                   `json:"created_at,omitempty"`
	UpdatedAt                    string                                   `json:"updated_at,omitempty"`
	Version                      int64                                    `json:"version,omitempty"`
	FinancialStatementID         string                                   `json:"financial_statement_id,omitempty"`
	ProcessedAt                  string                                   `json:"processed_at,omitempty"`
	AccountServicerTransactionID string                                   `json:"account_servicer_transaction_id,omitempty"`
	Payments                     []*FinancialMutationPayment              `json:"payments,omitempty"`
	LedgerAccountBookings        []*FinancialMutationLedgerAccountBooking `json:"ledger_account_bookings,omitempty"`
}

type FinancialMutationPayment struct {
	ID                    string `json:"id,omitempty"`
	AdministrationID      string `json:"administration_id,omitempty"`
	InvoiceType           string `json:"invoice_type,omitempty"`
	InvoiceID             string `json:"invoice_id,omitempty"`
	FinancialAccountID    string `json:"financial_account_id,omitempty"`
	UserID                int    `json:"user_id,omitempty"`
	PaymentTransactionID  string `json:"payment_transaction_id,omitempty"`
	TransactionIdentifier string `json:"transaction_identifier,omitempty"`
	Price                 string `json:"price,omitempty"`
	PriceBase             string `json:"price_base,omitempty"`
	PaymentDate           string `json:"payment_date,omitempty"`
	CreditInvoiceID       string `json:"credit_invoice_id,omitempty"`
	FinancialMutationID   string `json:"financial_mutation_id,omitempty"`
	LedgerAccountID       string `json:"ledger_account_id,omitempty"`
	LinkedPaymentID       string `json:"linked_payment_id,omitempty"`
	ManualPaymentAction   string `json:"manual_payment_action,omitempty"`
	CreatedAt             string `json:"created_at,omitempty"`
	UpdatedAt             string `json:"updated_at,omitempty"`
}

type FinancialMutationLedgerAccountBooking struct {
	ID                  string `json:"id,omitempty"`
	AdministrationID    string `json:"administration_id,omitempty"`
	FinancialMutationID string `json:"financial_mutation_id,omitempty"`
	LedgerAccountID     string `json:"ledger_account_id,omitempty"`
	ProjectID           string `json:"project_id,omitempty"`
	Description         string `json:"description,omitempty"`
	Price               string `json:"price,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
}

type FinancialStatementGateway struct {
	*Client
}

func (c *Client) FinancialStatement() *FinancialStatementGateway {
	return &FinancialStatementGateway{c}
}

func (c *FinancialStatementGateway) Create(financialStatement *FinancialStatement) (*FinancialStatement, error) {
	res, err := c.execute("POST", "financial_statements", &envelope{FinancialStatement: financialStatement})
	if err != nil {
		return financialStatement, err
	}

	switch res.StatusCode {
	case http.StatusCreated:
		return res.financialStatement()
	}

	return nil, res.error()
}

func (c *FinancialStatementGateway) Delete(financialStatement *FinancialStatement) error {
	res, err := c.execute("DELETE", "financial_statements/"+financialStatement.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
