package moneybird

import "testing"

func TestFinancialStatementGatewayCreateAndDelete(t *testing.T) {

	gateway := testClient.FinancialStatement()

	// create
	financialStatement, err := gateway.Create(&FinancialStatement{
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

	// delete
	err = gateway.Delete(financialStatement)
	if err != nil {
		t.Error(err)
	}

}
