package moneybird

import "net/http"

type PurchaseInvoice struct {
	ID                    string                   `json:"id,omitempty"`
	AdministrationID      string                   `json:"administration_id,omitempty"`
	ContactID             string                   `json:"contact_id,omitempty"`
	Contact               Contact                  `json:"contact,omitempty"`
	Reference             string                   `json:"reference,omitempty"`
	Date                  string                   `json:"date,omitempty"`
	DueDate               string                   `json:"due_date,omitempty"`
	EntryNumber           int                      `json:"entry_number,omitempty"`
	State                 string                   `json:"state,omitempty"`
	Currency              string                   `json:"currency,omitempty"`
	ExchangeRate          string                   `json:"exchange_rate,omitempty"`
	RevenueInvoice        bool                     `json:"revenue_invoice,omitempty"`
	PricesAreInclTax      bool                     `json:"prices_are_incl_tax,omitempty"`
	Origin                string                   `json:"origin,omitempty"`
	PaidAt                string                   `json:"paid_at,omitempty"`
	TaxNumber             string                   `json:"tax_number,omitempty"`
	TotalPriceExclTax     string                   `json:"total_price_excl_tax,omitempty"`
	TotalPriceExclTaxBase string                   `json:"total_price_excl_tax_base,omitempty"`
	TotalPriceInclTax     string                   `json:"total_price_incl_tax,omitempty"`
	TotalPriceInclTaxBase string                   `json:"total_price_incl_tax_base,omitempty"`
	CreatedAt             string                   `json:"created_at,omitempty"`
	UpdatedAt             string                   `json:"updated_at,omitempty"`
	Version               int64                    `json:"version,omitempty"`
	Details               []*PurchaseInvoiceDetail `json:"details_attributes,omitempty"`
	Payments              []any                    `json:"payments,omitempty"`
	Attachments           []any                    `json:"attachments,omitempty"`
	Events                []*PurchaseInvoiceEvent  `json:"events,omitempty"`
}

type PurchaseInvoiceDetail struct {
	ID                                string   `json:"id,omitempty"`
	TaxRateID                         string   `json:"tax_rate_id,omitempty"`
	LedgerAccountID                   string   `json:"ledger_account_id,omitempty"`
	ProjectID                         string   `json:"project_id,omitempty"`
	ProductID                         string   `json:"product_id,omitempty"`
	Amount                            string   `json:"amount,omitempty"`
	AmountDecimal                     string   `json:"amount_decimal,omitempty"`
	Description                       string   `json:"description,omitempty"`
	Price                             string   `json:"price,omitempty"`
	Period                            string   `json:"period,omitempty"`
	RowOrder                          int      `json:"row_order,omitempty"`
	TotalPriceExclTaxWithDiscount     string   `json:"total_price_excl_tax_with_discount,omitempty"`
	TotalPriceExclTaxWithDiscountBase string   `json:"total_price_excl_tax_with_discount_base,omitempty"`
	TaxReportReference                []string `json:"tax_report_reference,omitempty"`
	MandatoryTaxText                  string   `json:"mandatory_tax_text,omitempty"`
	CreatedAt                         string   `json:"created_at,omitempty"`
	UpdatedAt                         string   `json:"updated_at,omitempty"`
}

type PurchaseInvoiceEvent struct {
	UserID         string         `json:"user_id,omitempty"`
	Action         string         `json:"action,omitempty"`
	LinkEntityID   string         `json:"link_entity_id,omitempty"`
	LinkEntityType string         `json:"link_entity_type,omitempty"`
	Data           map[string]any `json:"data,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
}

type PurchaseInvoiceGateway struct {
	*Client
}

func (c *Client) PurchaseInvoice() *PurchaseInvoiceGateway {
	return &PurchaseInvoiceGateway{c}
}

// Create creates the invoice in Moneybird
func (c *PurchaseInvoiceGateway) Create(invoice *PurchaseInvoice) (*PurchaseInvoice, error) {
	res, err := c.execute("POST", "documents/purchase_invoices", &envelope{PurchaseInvoice: invoice})
	if err != nil {
		return invoice, err
	}

	switch res.StatusCode {
	case http.StatusCreated:
		return res.purchaseInvoice()
	}

	return nil, res.error()
}
