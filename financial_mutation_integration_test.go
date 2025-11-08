package moneybird

import "testing"

func TestLinkBooking(t *testing.T) {

	// we need a statement with mutations to test linking bookings
	fsGateway := testClient.FinancialStatement()

	// create
	financialStatement, err := fsGateway.Create(&FinancialStatement{
		FinancialAccountID: "469657091992192361",
		Reference:          "TestFinancialStatement_Create",
		FinancialMutationsAttributes: map[string]*FinancialMutationAttributes{
			"1": {
				Date:                         "2025-11-02",
				Message:                      "Test mutation",
				Amount:                       "123,45",
				AccountServicerTransactionId: "TEST-12345",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// delete the created statement on defer
	defer func() {
		err = fsGateway.Delete(financialStatement)
		if err != nil {
			t.Error(err)
		}
	}()

	// link booking on first mutation of statement
	fmGateway := testClient.FinancialMutation()
	err = fmGateway.LinkBooking(financialStatement.FinancialMutations[0].ID, &LinkBookingRequest{
		BookingType: "LedgerAccount",
		BookingID:   "210552408975083053",
		PriceBase:   "5.12",
		Description: "Linked via integration test",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log("LinkBooking executed successfully")
}
