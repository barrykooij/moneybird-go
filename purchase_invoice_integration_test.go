package moneybird

import (
	"testing"
	"time"
)

func TestPurchaseInvoiceGatewayCRUD(t *testing.T) {
	var err error
	// create contact
	contact := &Contact{
		Email:     "johndoe@email.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	contact, err = testClient.Contact().Create(contact)
	if err != nil {
		t.Fatalf("ContactGateway.Create: %s", err)
	}

	// delete contact (deferred)
	defer func() {
		err = testClient.Contact().Delete(contact)
		if err != nil {
			t.Errorf("ContactGateway.Delete: %s", err)
		}
	}()

	gateway := testClient.PurchaseInvoice()

	// create purchase invoice
	purchaseInvoice := &PurchaseInvoice{
		ContactID: contact.ID,
		Reference: "test-reference",
		Date:      time.Now().Format("2006-01-02"),
		Details: []*PurchaseInvoiceDetail{
			{
				Amount:      "1",
				Price:       "10.00",
				Description: "Test Service",
				TaxRateID:   "469726274244314641",
			},
		},
	}
	purchaseInvoice, err = gateway.Create(purchaseInvoice)
	if err != nil {
		t.Fatalf("PurchaseInvoiceGateway.Create: %s", err) // abandon test if invoice creation fails
	}

	// update purchase invoice
	purchaseInvoice.Reference = "updated-reference"
	purchaseInvoice, err = gateway.Update(purchaseInvoice)
	if err != nil {
		t.Errorf("PurchaseInvoiceGateway.Update: %s", err)
	}

	if purchaseInvoice.Reference != "updated-reference" {
		t.Error("PurchaseInvoiceGateway.Update: reference was not properly updated")
	}

	// get purchase invoice
	purchaseInvoice, err = gateway.Get(purchaseInvoice.ID)
	if err != nil {
		t.Errorf("PurchaseInvoiceGateway.Get: %s", err)
	}

	if purchaseInvoice.Contact.ID != contact.ID {
		t.Errorf("PurchaseInvoiceGateway.Get: invoice contact ID does not match, got %#v", purchaseInvoice.Contact.ID)
	}
}
