package moneybird

import "testing"

func TestFinancialStatementGateway_Create(t *testing.T) {

	gateway := testClient.FinancialStatement()

	_, err := gateway.Create(&FinancialStatement{
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

}
