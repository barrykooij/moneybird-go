package moneybird

import "net/http"

type FinancialStatement struct {
	ID                           string                                  `json:"id,omitempty"`
	FinancialAccountID           string                                  `json:"financial_account_id"`
	Reference                    string                                  `json:"reference"`
	OfficialDate                 string                                  `json:"official_date,omitempty"`
	OfficialBalance              string                                  `json:"official_balance,omitempty"`
	ImporterService              string                                  `json:"importer_service,omitempty"`
	FinancialMutations           []FinancialMutation                     `json:"financial_mutations,omitempty"`
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
	ID                           string `json:"id"`
	AdministrationID             string `json:"administration_id"`
	Amount                       string `json:"amount"`
	Code                         string `json:"code"`
	Date                         string `json:"date"`
	Message                      string `json:"message"`
	ContraAccountName            string `json:"contra_account_name"`
	ContraAccountNumber          string `json:"contra_account_number"`
	State                        string `json:"state"`
	AmountOpen                   string `json:"amount_open"`
	SepaFields                   string `json:"sepa_fields"`
	BatchReference               string `json:"batch_reference"`
	FinancialAccountID           string `json:"financial_account_id"`
	Currency                     string `json:"currency"`
	OriginalAmount               string `json:"original_amount"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	Version                      int64  `json:"version"`
	FinancialStatementID         string `json:"financial_statement_id"`
	ProcessedAt                  string `json:"processed_at"`
	AccountServicerTransactionID string `json:"account_servicer_transaction_id"`
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
