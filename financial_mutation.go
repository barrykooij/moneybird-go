package moneybird

import (
	"encoding/json"
	"net/http"
)

type LinkBookingRequest struct {
	BookingType            string `json:"booking_type"`
	BookingID              string `json:"booking_id,omitempty"`
	PriceBase              string `json:"price_base,omitempty"`
	Price                  string `json:"price,omitempty"`
	Description            string `json:"description,omitempty"`
	PaymentBatchIdentifier string `json:"payment_batch_identifier,omitempty"`
	ProjectID              string `json:"project_id,omitempty"`
}

type FinancialMutationGateway struct {
	*Client
}

func (c *Client) FinancialMutation() *FinancialMutationGateway {
	return &FinancialMutationGateway{c}
}

func (c *FinancialMutationGateway) LinkBooking(financialMutationID string, request *LinkBookingRequest) error {
	var data []byte
	var err error

	data, err = json.Marshal(request)
	if err != nil {
		return err
	}

	// we have to use the callApi method here because Moneybird uses no parent key in the body json for this endpoint
	// making the whole envelope struct setup unusable
	res, err := c.callApi("PATCH", "financial_mutations/"+financialMutationID+"/link_booking", data)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return nil
	}

	return res.error()
}
