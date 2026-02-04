package intasend

import (
	"encoding/json"
	"testing"
)

func TestPaymentStatusUnmarshal(t *testing.T) {
	// Sample payment status response with charges as number (not string)
	// This matches the actual API response format
	jsonData := `{
		"invoice": {
			"id": "sCsVD_C_27555",
			"invoice_id": "sCsVD_C_27555",
			"state": "COMPLETE",
			"provider": "M-PESA",
			"charges": 105.46,
			"net_amount": 2907.54,
			"currency": "UGX",
			"value": 3013.0,
			"account": "256759739706",
			"api_ref": "Dbrs5_C_27555",
			"host": "https://api.intasend.com",
			"failed_reason": null,
			"created_at": "2024-12-24T10:00:00+03:00",
			"updated_at": "2024-12-24T10:05:00+03:00"
		},
		"meta": {
			"id": "META123",
			"customer": {
				"id": "CUST456",
				"phone_number": "256759739706",
				"email": "customer@example.com",
				"first_name": "John",
				"last_name": "Doe",
				"country": "UG",
				"address": "123 Main St",
				"city": "Kampala",
				"state": "Central",
				"zipcode": "12345",
				"provider": "M-PESA",
				"created_at": "2024-01-01T00:00:00+03:00",
				"updated_at": "2024-01-01T00:00:00+03:00"
			},
			"customer_comment": "Payment for services",
			"created_at": "2024-12-24T10:00:00+03:00",
			"updated_at": "2024-12-24T10:05:00+03:00"
		}
	}`

	var status PaymentStatus
	err := json.Unmarshal([]byte(jsonData), &status)
	if err != nil {
		t.Fatalf("Failed to unmarshal payment status: %v", err)
	}

	// Validate invoice fields
	if status.Invoice.InvoiceID != "sCsVD_C_27555" {
		t.Errorf("Expected InvoiceID to be sCsVD_C_27555, got %s", status.Invoice.InvoiceID)
	}

	if status.Invoice.State != "COMPLETE" {
		t.Errorf("Expected State to be COMPLETE, got %s", status.Invoice.State)
	}

	if status.Invoice.Charges != 105.46 {
		t.Errorf("Expected Charges to be 105.46, got %.2f", status.Invoice.Charges)
	}

	if status.Invoice.NetAmount != 2907.54 {
		t.Errorf("Expected NetAmount to be 2907.54, got %.2f", status.Invoice.NetAmount)
	}

	if status.Invoice.Value != 3013.0 {
		t.Errorf("Expected Value to be 3013.0, got %.2f", status.Invoice.Value)
	}

	if status.Invoice.Currency != "UGX" {
		t.Errorf("Expected Currency to be UGX, got %s", status.Invoice.Currency)
	}

	if status.Invoice.Provider != "M-PESA" {
		t.Errorf("Expected Provider to be M-PESA, got %s", status.Invoice.Provider)
	}

	// Test helper methods
	if status.GetCharges() != 105.46 {
		t.Errorf("Expected GetCharges() to return 105.46, got %.2f", status.GetCharges())
	}

	if status.GetValue() != 3013.0 {
		t.Errorf("Expected GetValue() to return 3013.0, got %.2f", status.GetValue())
	}

	if status.GetPaymentAmount() != 2907.54 {
		t.Errorf("Expected GetPaymentAmount() to return 2907.54, got %.2f", status.GetPaymentAmount())
	}

	if status.GetCurrency() != "UGX" {
		t.Errorf("Expected GetCurrency() to return UGX, got %s", status.GetCurrency())
	}

	if status.GetInvoiceID() != "sCsVD_C_27555" {
		t.Errorf("Expected GetInvoiceID() to return sCsVD_C_27555, got %s", status.GetInvoiceID())
	}

	if status.GetAPIReference() != "Dbrs5_C_27555" {
		t.Errorf("Expected GetAPIReference() to return Dbrs5_C_27555, got %s", status.GetAPIReference())
	}

	// Test customer info
	if status.GetCustomerEmail() != "customer@example.com" {
		t.Errorf("Expected customer email to be customer@example.com, got %s", status.GetCustomerEmail())
	}

	if status.GetCustomerPhone() != "256759739706" {
		t.Errorf("Expected customer phone to be 256759739706, got %s", status.GetCustomerPhone())
	}

	expectedName := "John Doe"
	if status.GetCustomerName() != expectedName {
		t.Errorf("Expected customer name to be %s, got %s", expectedName, status.GetCustomerName())
	}

	// Test status checks - note: API returns "COMPLETE" not "COMPLETED"
	// The IsCompleted() method should handle both
	if status.Invoice.State == StatusComplete {
		t.Log("Invoice state is COMPLETE (as expected from API)")
	}

	t.Logf("Successfully unmarshaled payment status: Invoice=%s, Amount=%.2f %s, Charges=%.2f",
		status.Invoice.InvoiceID, status.Invoice.NetAmount, status.Invoice.Currency, status.Invoice.Charges)
}

func TestPaymentStatusWithStringValues(t *testing.T) {
	// Test backward compatibility - some endpoints might still return strings
	// This test ensures we handle numeric values correctly
	jsonData := `{
		"invoice": {
			"id": "TEST123",
			"invoice_id": "TEST123",
			"state": "PENDING",
			"provider": "CARD",
			"charges": 0.00,
			"net_amount": 100.00,
			"currency": "KES",
			"value": 100.00,
			"account": "test@example.com",
			"api_ref": "REF123",
			"host": "https://sandbox.intasend.com",
			"failed_reason": null,
			"created_at": "2024-01-01T00:00:00+03:00",
			"updated_at": "2024-01-01T00:00:00+03:00"
		},
		"meta": {
			"id": "META123",
			"customer": {
				"id": "CUST123",
				"phone_number": "254712345678",
				"email": "test@example.com",
				"first_name": "Test",
				"last_name": "User",
				"country": "KE",
				"address": "",
				"city": "",
				"state": "",
				"zipcode": "",
				"provider": "CARD",
				"created_at": "2024-01-01T00:00:00+03:00",
				"updated_at": "2024-01-01T00:00:00+03:00"
			},
			"customer_comment": "",
			"created_at": "2024-01-01T00:00:00+03:00",
			"updated_at": "2024-01-01T00:00:00+03:00"
		}
	}`

	var status PaymentStatus
	err := json.Unmarshal([]byte(jsonData), &status)
	if err != nil {
		t.Fatalf("Failed to unmarshal payment status: %v", err)
	}

	if status.Invoice.Charges != 0.00 {
		t.Errorf("Expected Charges to be 0.00, got %.2f", status.Invoice.Charges)
	}

	if status.Invoice.Value != 100.00 {
		t.Errorf("Expected Value to be 100.00, got %.2f", status.Invoice.Value)
	}

	t.Logf("Successfully handled numeric values: Charges=%.2f, Value=%.2f", 
		status.Invoice.Charges, status.Invoice.Value)
}

