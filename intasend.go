package intasend

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// Client represents the IntaSend API client
type Client struct {
	PublishableKey string
	Token          string
	BaseURL        string

	HTTPClient *http.Client
	Test       bool
	ShowLogs   bool
}

// PaymentRequest represents the payment checkout request payload

// NewClient creates a new IntaSend API client
func NewClient(publishableKey, token string, test bool, showlogs bool) *Client {
	baseURL := "https://payment.intasend.com"
	// apiBaseURL := "https://api.intasend.com"
	if test {
		baseURL = "https://sandbox.intasend.com"
		// apiBaseURL = "https://sandbox.intasend.com"
	}

	return &Client{
		PublishableKey: publishableKey,
		Token:          token,
		BaseURL:        baseURL,
		ShowLogs:       showlogs,
		// APIBaseURL:     apiBaseURL,
		Test: test,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CreateCheckoutLink generates a secure checkout link for payment
func (c *Client) CreateCheckoutLink(req *PaymentRequest) (*PaymentResponse, error) {
	// Validate required fields
	if c.PublishableKey == "" {
		return nil, fmt.Errorf("publishable key is required")
	}

	// Set default values
	if req.Currency == "" {
		req.Currency = CurrencyKES
	}
	if req.CardTarrif == "" {
		req.CardTarrif = CUSTOMER_PAYS
	}
	if req.MobileTarrif == "" {
		req.MobileTarrif = CUSTOMER_PAYS
	}

	// Use the HTTP wrapper to make the request
	var paymentResp PaymentResponse
	err := c.PostJSON("api/v1/checkout/", req, false, true, &paymentResp)
	if err != nil {
		return nil, err
	}

	return &paymentResp, nil
}

// QuickCheckout creates a simple checkout link with minimal required fields
func (c *Client) QuickCheckout(phoneNumber, email string, amount float64, currency CurrencyType, comment, redirectURL string) (*PaymentResponse, error) {
	req := &PaymentRequest{
		PhoneNumber: phoneNumber,
		Email:       email,
		Amount:      amount,
		Currency:    currency,
		Comment:     comment,
		RedirectURL: redirectURL,
	}

	return c.CreateCheckoutLink(req)
}

// PaymentRequestBuilder provides a fluent interface for building payment requests
type PaymentRequestBuilder struct {
	request *PaymentRequest
}

// NewPaymentRequest creates a new payment request builder
func NewPaymentRequest() *PaymentRequestBuilder {
	return &PaymentRequestBuilder{
		request: &PaymentRequest{
			Currency:     CurrencyKES,
			CardTarrif:   CUSTOMER_PAYS,
			MobileTarrif: CUSTOMER_PAYS,
		},
	}
}

// WithPhoneNumber sets the customer's phone number
func (b *PaymentRequestBuilder) WithPhoneNumber(phoneNumber string) *PaymentRequestBuilder {
	b.request.PhoneNumber = phoneNumber
	return b
}

// WithEmail sets the customer's email
func (b *PaymentRequestBuilder) WithEmail(email string) *PaymentRequestBuilder {
	b.request.Email = email
	return b
}

// WithAmount sets the payment amount
func (b *PaymentRequestBuilder) WithAmount(amount float64) *PaymentRequestBuilder {
	b.request.Amount = amount
	return b
}

// WithCurrency sets the currency
func (b *PaymentRequestBuilder) WithCurrency(currency CurrencyType) *PaymentRequestBuilder {
	b.request.Currency = currency
	return b
}

// WithComment sets a comment for the payment
func (b *PaymentRequestBuilder) WithComment(comment string) *PaymentRequestBuilder {
	b.request.Comment = comment
	return b
}

// WithRedirectURL sets the redirect URL after payment
func (b *PaymentRequestBuilder) WithRedirectURL(url string) *PaymentRequestBuilder {
	b.request.RedirectURL = url
	return b
}

// WithAPIRef sets an API reference for tracking
func (b *PaymentRequestBuilder) WithAPIRef(ref string) *PaymentRequestBuilder {
	b.request.APIRef = ref
	return b
}

// WithCustomerInfo sets customer information
func (b *PaymentRequestBuilder) WithCustomerInfo(firstName, lastName string) *PaymentRequestBuilder {
	b.request.FirstName = firstName
	b.request.LastName = lastName
	return b
}

// WithBillingAddress sets billing address information
func (b *PaymentRequestBuilder) WithBillingAddress(country, address, city, state, zipCode string) *PaymentRequestBuilder {
	b.request.Country = country
	b.request.Address = address
	b.request.City = city
	b.request.State = state
	b.request.ZipCode = zipCode
	return b
}

// WithMethod restricts payment to a specific method
func (b *PaymentRequestBuilder) WithMethod(method string) *PaymentRequestBuilder {
	b.request.Method = PaymentMethodType(method)
	return b
}

// WithCardTariff sets who pays card processing fees
func (b *PaymentRequestBuilder) WithCardTariff(tariff string) *PaymentRequestBuilder {
	b.request.CardTarrif = TarriffType(tariff)
	return b
}

// WithMobileTariff sets who pays mobile payment fees
func (b *PaymentRequestBuilder) WithMobileTariff(tariff string) *PaymentRequestBuilder {
	b.request.MobileTarrif = TarriffType(tariff)
	return b
}

// Build returns the constructed payment request
func (b *PaymentRequestBuilder) Build() *PaymentRequest {
	return b.request
}

// ValidateEmail performs basic email validation
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	// Basic email validation
	return len(email) > 3 &&
		len(email) < 255 &&
		bytes.ContainsRune([]byte(email), '@') &&
		bytes.ContainsRune([]byte(email), '.')
}

// GetPaymentStatus retrieves the payment status using invoice ID
func (c *Client) GetPaymentStatus(invoiceID string) (*PaymentStatus, error) {
	// Validate required fields
	if c.PublishableKey == "" {
		return nil, fmt.Errorf("publishable key is required")
	}

	if invoiceID == "" {
		return nil, fmt.Errorf("invoice ID is required")
	}

	// Create request payload
	payload := map[string]string{
		"invoice_id": invoiceID,
	}

	// Use the HTTP wrapper to make the request
	var statusResp PaymentStatus
	err := c.PostJSON("api/v1/payment/status/", payload, true, true, &statusResp)
	if err != nil {
		return nil, err
	}

	return &statusResp, nil
}

// CheckPaymentStatus is an alias for GetPaymentStatus for consistency with Python SDK
func (c *Client) CheckPaymentStatus(invoiceID string) (*PaymentStatus, error) {
	return c.GetPaymentStatus(invoiceID)
}
