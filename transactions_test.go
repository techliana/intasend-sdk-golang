package intasend

import (
	"encoding/json"
	"testing"
)

func TestTransactionUnmarshal(t *testing.T) {
	// Sample transaction data based on IntaSend API response structure
	jsonData := `{
		"transaction_id": "TXN123456",
		"invoice": {
			"invoice_id": "INV789",
			"state": "COMPLETE",
			"provider": "M-PESA",
			"charges": 10.5,
			"net_amount": "989.50",
			"currency": "KES",
			"value": 1000.0,
			"account": "254712345678",
			"api_ref": "REF123",
			"clearing_status": "AVAILABLE",
			"mpesa_reference": "ABC123XYZ",
			"host": "https://api.intasend.com",
			"card_info": {
				"bin_country": null,
				"card_type": null
			},
			"retry_count": 0,
			"failed_reason": null,
			"failed_code": null,
			"failed_code_link": null,
			"created_at": "2024-12-24T10:00:00+03:00",
			"updated_at": "2024-12-24T10:05:00+03:00"
		},
		"currency": "KES",
		"value": 1000.0,
		"running_balance": 5000.0,
		"narrative": "Payment received from customer",
		"trans_type": "DEPOSIT",
		"status": "COMPLETED",
		"created_at": "2024-12-24T10:00:00+03:00",
		"updated_at": "2024-12-24T10:05:00+03:00"
	}`

	var transaction Result
	err := json.Unmarshal([]byte(jsonData), &transaction)
	if err != nil {
		t.Fatalf("Failed to unmarshal transaction: %v", err)
	}

	// Validate transaction fields
	if transaction.TransactionID != "TXN123456" {
		t.Errorf("Expected TransactionID to be TXN123456, got %s", transaction.TransactionID)
	}

	if transaction.Currency != "KES" {
		t.Errorf("Expected Currency to be KES, got %s", transaction.Currency)
	}

	if transaction.Value != 1000.0 {
		t.Errorf("Expected Value to be 1000.0, got %.2f", transaction.Value)
	}

	if transaction.RunningBalance != 5000.0 {
		t.Errorf("Expected RunningBalance to be 5000.0, got %.2f", transaction.RunningBalance)
	}

	if transaction.TransType != "DEPOSIT" {
		t.Errorf("Expected TransType to be DEPOSIT, got %s", transaction.TransType)
	}

	if transaction.Status != "COMPLETED" {
		t.Errorf("Expected Status to be COMPLETED, got %s", transaction.Status)
	}

	// Validate invoice
	if transaction.Invoice == nil {
		t.Fatal("Expected Invoice to be present")
	}

	if transaction.Invoice.InvoiceID != "INV789" {
		t.Errorf("Expected Invoice.InvoiceID to be INV789, got %s", transaction.Invoice.InvoiceID)
	}

	if transaction.Invoice.State != "COMPLETE" {
		t.Errorf("Expected Invoice.State to be COMPLETE, got %s", transaction.Invoice.State)
	}

	// Test helper methods
	if !transaction.IsCompleted() {
		t.Error("Expected transaction to be completed")
	}

	if !transaction.IsDeposit() {
		t.Error("Expected transaction to be a deposit")
	}

	if transaction.IsPending() {
		t.Error("Expected transaction not to be pending")
	}

	if transaction.IsFailed() {
		t.Error("Expected transaction not to be failed")
	}

	if transaction.IsWithdrawal() {
		t.Error("Expected transaction not to be a withdrawal")
	}

	// Test GetInvoiceID helper
	invoiceID := transaction.GetInvoiceID()
	if invoiceID != "INV789" {
		t.Errorf("Expected GetInvoiceID to return INV789, got %s", invoiceID)
	}

	t.Logf("Successfully unmarshaled transaction: %+v", transaction)
}

func TestTransactionRespUnmarshal(t *testing.T) {
	// Sample paginated transactions response
	jsonData := `{
		"count": 2,
		"next": null,
		"previous": null,
		"results": [
			{
				"transaction_id": "TXN001",
				"invoice": null,
				"currency": "KES",
				"value": 500.0,
				"running_balance": 5500.0,
				"narrative": "Withdrawal to bank",
				"trans_type": "WITHDRAWAL",
				"status": "COMPLETED",
				"created_at": "2024-12-24T09:00:00+03:00",
				"updated_at": "2024-12-24T09:05:00+03:00"
			},
			{
				"transaction_id": "TXN002",
				"invoice": null,
				"currency": "KES",
				"value": 1000.0,
				"running_balance": 6500.0,
				"narrative": "Transfer from wallet",
				"trans_type": "TRANSFER",
				"status": "PENDING",
				"created_at": "2024-12-24T11:00:00+03:00",
				"updated_at": "2024-12-24T11:00:00+03:00"
			}
		]
	}`

	var response TransactionResp
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal transaction response: %v", err)
	}

	if response.Count != 2 {
		t.Errorf("Expected Count to be 2, got %d", response.Count)
	}

	if len(response.Results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(response.Results))
	}

	// Check first transaction
	tx1 := response.Results[0]
	if !tx1.IsWithdrawal() {
		t.Error("Expected first transaction to be a withdrawal")
	}

	// Check second transaction
	tx2 := response.Results[1]
	if !tx2.IsPending() {
		t.Error("Expected second transaction to be pending")
	}

	if !tx2.IsTransfer() {
		t.Error("Expected second transaction to be a transfer")
	}

	t.Logf("Successfully unmarshaled %d transactions", len(response.Results))
}

