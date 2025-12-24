package intasend

import (
	"encoding/json"
	"testing"
)

func TestInvoiceItemUnmarshal(t *testing.T) {
	// Real API response data
	jsonData := `{
		"invoice_id":"Y23RDWZ",
		"state":"COMPLETE",
		"provider":"INTASEND-XB-PUSH",
		"charges":105.46,
		"net_amount":"2907.54",
		"currency":"UGX",
		"value":3013.0,
		"account":"256759739706",
		"api_ref":"Dbrs5_C_22044",
		"clearing_status":"AVAILABLE",
		"mpesa_reference":null,
		"host":"159.65.204.159",
		"card_info":{
			"bin_country":null,
			"card_type":null
		},
		"retry_count":0,
		"failed_reason":null,
		"failed_code":null,
		"failed_code_link":null,
		"created_at":"2025-12-24T12:57:26.794896+03:00",
		"updated_at":"2025-12-24T12:58:31.464570+03:00"
	}`

	var invoice InvoiceItem
	err := json.Unmarshal([]byte(jsonData), &invoice)
	if err != nil {
		t.Fatalf("Failed to unmarshal invoice: %v", err)
	}

	// Verify the fields
	if invoice.InvoiceID != "Y23RDWZ" {
		t.Errorf("Expected InvoiceID 'Y23RDWZ', got '%s'", invoice.InvoiceID)
	}

	if invoice.State != "COMPLETE" {
		t.Errorf("Expected State 'COMPLETE', got '%s'", invoice.State)
	}

	if invoice.Charges != 105.46 {
		t.Errorf("Expected Charges 105.46, got %f", invoice.Charges)
	}

	if invoice.NetAmount != "2907.54" {
		t.Errorf("Expected NetAmount '2907.54', got '%s'", invoice.NetAmount)
	}

	if invoice.Value != 3013.0 {
		t.Errorf("Expected Value 3013.0, got %f", invoice.Value)
	}

	if invoice.Currency != "UGX" {
		t.Errorf("Expected Currency 'UGX', got '%s'", invoice.Currency)
	}

	if invoice.ClearingStatus != "AVAILABLE" {
		t.Errorf("Expected ClearingStatus 'AVAILABLE', got '%s'", invoice.ClearingStatus)
	}

	if invoice.RetryCount != 0 {
		t.Errorf("Expected RetryCount 0, got %d", invoice.RetryCount)
	}

	// Test helper methods
	netAmount, err := invoice.GetNetAmountFloat()
	if err != nil {
		t.Errorf("Failed to convert NetAmount to float: %v", err)
	}
	if netAmount != 2907.54 {
		t.Errorf("Expected NetAmount float 2907.54, got %f", netAmount)
	}

	if !invoice.IsCompleted() {
		t.Error("Expected invoice to be completed")
	}

	t.Logf("Successfully unmarshaled invoice: %+v", invoice)
}

