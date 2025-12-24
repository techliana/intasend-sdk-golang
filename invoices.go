package intasend

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// InvoiceItem represents a single invoice in the IntaSend system
type InvoiceItem struct {
	InvoiceID      string    `json:"invoice_id"`
	State          string    `json:"state"`
	Provider       string    `json:"provider"`
	Charges        float64   `json:"charges"`    // Charges as number (e.g., 105.46)
	NetAmount      string    `json:"net_amount"` // Net amount as string (e.g., "2907.54")
	Currency       string    `json:"currency"`
	Value          float64   `json:"value"` // Value as number (e.g., 3013.0)
	Account        string    `json:"account"`
	APIRef         string    `json:"api_ref"`
	ClearingStatus string    `json:"clearing_status"` // Clearing status (e.g., "AVAILABLE")
	MpesaReference *string   `json:"mpesa_reference"`
	Host           string    `json:"host"`
	CardInfo       CardInfo  `json:"card_info"`   // Card information
	RetryCount     int       `json:"retry_count"` // Number of retry attempts
	FailedReason   *string   `json:"failed_reason"`
	FailedCode     *string   `json:"failed_code"`
	FailedCodeLink *string   `json:"failed_code_link"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// PaginatedInvoices represents the paginated response for listing invoices
type PaginatedInvoices struct {
	Count    int           `json:"count"`
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []InvoiceItem `json:"results"`
}

// ListInvoicesParams defines optional filters and pagination for listing invoices
type ListInvoicesParams struct {
	Page     *int   // optional page number
	PageSize *int   // optional page size
	State    string // optional filter by state (PENDING, COMPLETED, FAILED, etc.)
	Currency string // optional filter by currency
	APIRef   string // optional filter by API reference
}

// ListInvoices retrieves a paginated list of invoices with optional filters
func (c *Client) ListInvoices(params *ListInvoicesParams) (*PaginatedInvoices, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to list invoices")
	}

	// Build query parameters
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page != nil {
			queryParams["page"] = strconv.Itoa(*params.Page)
		}
		if params.PageSize != nil {
			queryParams["page_size"] = strconv.Itoa(*params.PageSize)
		}
		if params.State != "" {
			queryParams["state"] = params.State
		}
		if params.Currency != "" {
			queryParams["currency"] = params.Currency
		}
		if params.APIRef != "" {
			queryParams["api_ref"] = params.APIRef
		}
	}

	// Use the HTTP wrapper to make the request
	var result PaginatedInvoices
	err := c.GetJSON("api/v1/invoices/", queryParams, true, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetInvoice retrieves a single invoice by its ID
func (c *Client) GetInvoice(invoiceID string) (*InvoiceItem, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to get invoice")
	}
	if invoiceID == "" {
		return nil, fmt.Errorf("invoiceID is required")
	}

	// Use the common HTTP JSON helper
	endpoint := fmt.Sprintf("api/v1/invoices/%s/", url.PathEscape(invoiceID))
	var invoice InvoiceItem
	if err := c.GetJSON(endpoint, nil, true, &invoice); err != nil {
		return nil, err
	}

	return &invoice, nil
}

// Helper methods for InvoiceItem

// IsCompleted checks if the invoice is completed
func (i *InvoiceItem) IsCompleted() bool {
	return i.State == StatusCompleted || i.State == StatusComplete
}

// IsPending checks if the invoice is pending
func (i *InvoiceItem) IsPending() bool {
	return i.State == StatusPending
}

// IsFailed checks if the invoice failed
func (i *InvoiceItem) IsFailed() bool {
	return i.State == StatusFailed
}

// GetFailureReason returns the failure reason if the invoice failed
func (i *InvoiceItem) GetFailureReason() string {
	if i.FailedReason != nil {
		return *i.FailedReason
	}
	return ""
}

// GetMpesaReference returns the M-Pesa reference if available
func (i *InvoiceItem) GetMpesaReference() string {
	if i.MpesaReference != nil {
		return *i.MpesaReference
	}
	return ""
}

// GetNetAmountFloat converts the net amount string to float64
func (i *InvoiceItem) GetNetAmountFloat() (float64, error) {
	return strconv.ParseFloat(i.NetAmount, 64)
}

// GetCharges returns the charges as float64 (already a number)
func (i *InvoiceItem) GetCharges() float64 {
	return i.Charges
}

// GetValue returns the value as float64 (already a number)
func (i *InvoiceItem) GetValue() float64 {
	return i.Value
}
