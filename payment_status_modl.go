package intasend

import "time"

// PaymentStatus represents the complete payment status response from IntaSend
type PaymentStatus struct {
	Invoice Invoice `json:"invoice"`
	Meta    Meta    `json:"meta"`
}

// Invoice represents the invoice details in status response
type Invoice struct {
	ID           string  `json:"id"`            // Invoice ID (e.g., "XMSLWOS")
	InvoiceID    string  `json:"invoice_id"`    // Same as ID, used for tracking
	State        string  `json:"state"`         // Payment state (PENDING, COMPLETED, FAILED, etc.)
	Provider     string  `json:"provider"`      // Payment provider (M-PESA, CARD, etc.)
	Charges      string  `json:"charges"`       // Processing charges as string (e.g., "0.00")
	NetAmount    float64 `json:"net_amount"`    // Net payment amount after charges
	Currency     string  `json:"currency"`      // Payment currency (KES, USD, etc.)
	Value        string  `json:"value"`         // Payment value as string
	Account      string  `json:"account"`       // Customer account (email or phone)
	APIRef       string  `json:"api_ref"`       // Your API reference for tracking
	Host         string  `json:"host"`          // IntaSend host URL
	FailedReason *string `json:"failed_reason"` // Failure reason (null if not failed)
	CreatedAt    string  `json:"created_at"`    // Creation timestamp (ISO format)
	UpdatedAt    string  `json:"updated_at"`    // Last update timestamp (ISO format)
}

// Meta represents additional metadata in the payment status response
type Meta struct {
	ID              string   `json:"id"`               // Meta record ID
	Customer        Customer `json:"customer"`         // Customer information
	CustomerComment string   `json:"customer_comment"` // Customer's comment/note
	CreatedAt       string   `json:"created_at"`       // Meta creation timestamp
	UpdatedAt       string   `json:"updated_at"`       // Meta update timestamp
}

// Customer represents customer information in the payment status
type Customer struct {
	ID          string `json:"id"`           // Customer ID (e.g., "ZOEW022")
	PhoneNumber string `json:"phone_number"` // Customer phone number
	Email       string `json:"email"`        // Customer email address
	FirstName   string `json:"first_name"`   // Customer first name
	LastName    string `json:"last_name"`    // Customer last name
	Country     string `json:"country"`      // Country code (e.g., "KE")
	Address     string `json:"address"`      // Customer address
	City        string `json:"city"`         // Customer city
	State       string `json:"state"`        // Customer state/province
	ZipCode     string `json:"zipcode"`      // Customer zip/postal code
	Provider    string `json:"provider"`     // Payment provider used
	CreatedAt   string `json:"created_at"`   // Customer record creation timestamp
	UpdatedAt   string `json:"updated_at"`   // Customer record update timestamp
}

// Payment status constants
const (
	StatusPending    = "PENDING"    // Payment is pending/waiting for completion
	StatusCompleted  = "COMPLETED"  // Payment completed successfully (some endpoints)
	StatusComplete   = "COMPLETE"   // Payment completed successfully (invoices endpoint)
	StatusFailed     = "FAILED"     // Payment failed
	StatusProcessing = "PROCESSING" // Payment is being processed
	StatusCancelled  = "CANCELLED"  // Payment was cancelled
)

// Payment provider constants
const (
	ProviderMPESA = "M-PESA"
	ProviderCard  = "CARD"
	ProviderBank  = "BANK"
)

// Helper methods for PaymentStatus

// IsCompleted checks if the payment status indicates successful completion
func (ps *PaymentStatus) IsCompleted() bool {
	return ps.Invoice.State == StatusCompleted
}

// IsPending checks if the payment status indicates it's still pending
func (ps *PaymentStatus) IsPending() bool {
	return ps.Invoice.State == StatusPending
}

// IsFailed checks if the payment status indicates failure
func (ps *PaymentStatus) IsFailed() bool {
	return ps.Invoice.State == StatusFailed
}

// IsProcessing checks if the payment is currently being processed
func (ps *PaymentStatus) IsProcessing() bool {
	return ps.Invoice.State == StatusProcessing
}

// IsCancelled checks if the payment was cancelled
func (ps *PaymentStatus) IsCancelled() bool {
	return ps.Invoice.State == StatusCancelled
}

// GetFailureReason returns the failure reason if the payment failed
func (ps *PaymentStatus) GetFailureReason() string {
	if ps.Invoice.FailedReason != nil {
		return *ps.Invoice.FailedReason
	}
	return ""
}

// GetPaymentAmount returns the net payment amount
func (ps *PaymentStatus) GetPaymentAmount() float64 {
	return ps.Invoice.NetAmount
}

// GetPaymentProvider returns the payment provider (M-PESA, CARD, etc.)
func (ps *PaymentStatus) GetPaymentProvider() string {
	return ps.Invoice.Provider
}

// GetCurrency returns the payment currency
func (ps *PaymentStatus) GetCurrency() string {
	return ps.Invoice.Currency
}

// GetCustomerEmail returns the customer's email address
func (ps *PaymentStatus) GetCustomerEmail() string {
	return ps.Meta.Customer.Email
}

// GetCustomerPhone returns the customer's phone number
func (ps *PaymentStatus) GetCustomerPhone() string {
	return ps.Meta.Customer.PhoneNumber
}

// GetCustomerName returns the customer's full name
func (ps *PaymentStatus) GetCustomerName() string {
	firstName := ps.Meta.Customer.FirstName
	lastName := ps.Meta.Customer.LastName
	if firstName != "" && lastName != "" {
		return firstName + " " + lastName
	} else if firstName != "" {
		return firstName
	} else if lastName != "" {
		return lastName
	}
	return ""
}

// GetInvoiceID returns the invoice ID for tracking
func (ps *PaymentStatus) GetInvoiceID() string {
	return ps.Invoice.InvoiceID
}

// GetAPIReference returns your API reference if provided
func (ps *PaymentStatus) GetAPIReference() string {
	return ps.Invoice.APIRef
}

// GetCharges returns the processing charges
func (ps *PaymentStatus) GetCharges() string {
	return ps.Invoice.Charges
}

// IsInFinalState checks if the payment is in a final state (completed, failed, or cancelled)
func (ps *PaymentStatus) IsInFinalState() bool {
	return ps.IsCompleted() || ps.IsFailed() || ps.IsCancelled()
}

// GetCreatedAt returns the payment creation time as a parsed time.Time
func (ps *PaymentStatus) GetCreatedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, ps.Invoice.CreatedAt)
}

// GetUpdatedAt returns the payment update time as a parsed time.Time
func (ps *PaymentStatus) GetUpdatedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, ps.Invoice.UpdatedAt)
}

// GetStatusSummary returns a human-readable status summary
func (ps *PaymentStatus) GetStatusSummary() string {
	switch ps.Invoice.State {
	case StatusCompleted:
		return "Payment completed successfully"
	case StatusPending:
		return "Payment is pending completion"
	case StatusFailed:
		reason := ps.GetFailureReason()
		if reason != "" {
			return "Payment failed: " + reason
		}
		return "Payment failed"
	case StatusProcessing:
		return "Payment is being processed"
	case StatusCancelled:
		return "Payment was cancelled"
	default:
		return "Unknown payment status: " + ps.Invoice.State
	}
}

// ToMap converts the PaymentStatus to a map for easy serialization/logging
func (ps *PaymentStatus) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"invoice_id":     ps.Invoice.InvoiceID,
		"state":          ps.Invoice.State,
		"provider":       ps.Invoice.Provider,
		"amount":         ps.Invoice.NetAmount,
		"currency":       ps.Invoice.Currency,
		"customer_email": ps.Meta.Customer.Email,
		"customer_phone": ps.Meta.Customer.PhoneNumber,
		"customer_name":  ps.GetCustomerName(),
		"api_ref":        ps.Invoice.APIRef,
		"failed_reason":  ps.GetFailureReason(),
		"created_at":     ps.Invoice.CreatedAt,
		"updated_at":     ps.Invoice.UpdatedAt,
	}
}

// Sample JSON response structure for documentation
/*
{
    "invoice": {
        "id": "XMSLWOS",
        "invoice_id": "XMSLWOS",
        "state": "PENDING",
        "provider": "M-PESA",
        "charges": "0.00",
        "net_amount": 10.36,
        "currency": "KES",
        "value": "10.36",
        "account": "test@example.com",
        "api_ref": "ISL_faa26ef9-eb08-4353-b125-ec6a8f022815",
        "host": "https://sandbox.intasend.com",
        "failed_reason": null,
        "created_at": "2021-04-11T08:37:15.781977+03:00",
        "updated_at": "2021-04-11T08:37:15.782011+03:00"
    },
    "meta": {
        "id": "5aec8e0b-8d96-429b-98b7-5361198160bd",
        "customer": {
            "id": "ZOEW022",
            "phone_number": "",
            "email": "test@example.com",
            "first_name": "FELIX",
            "last_name": "CHERUIYOT",
            "country": "KE",
            "address": "Westlands",
            "city": "Nairobi",
            "state": "Nairobi",
            "zipcode": "2020",
            "provider": "M-PESA",
            "created_at": "2020-08-06T16:24:06.247397+03:00",
            "updated_at": "2021-04-11T08:37:15.755013+03:00"
        },
        "customer_comment": "",
        "created_at": "2021-04-11T08:37:15.810438+03:00",
        "updated_at": "2021-04-11T08:37:15.810475+03:00"
    }
}
*/
