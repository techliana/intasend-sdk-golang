package intasend

import (
	"encoding/json"
	"time"
)

// PaymentRequest represents the payment checkout request payload
type PaymentRequest struct {
	PhoneNumber  string      `json:"phone_number,omitempty"`
	Email        string      `json:"email,omitempty"`
	Amount       float64     `json:"amount,omitempty"`
	Currency     string      `json:"currency,omitempty"`
	Comment      string      `json:"comment,omitempty"`
	RedirectURL  string      `json:"redirect_url,omitempty"`
	APIRef       string      `json:"api_ref,omitempty"`
	FirstName    string      `json:"first_name,omitempty"`
	LastName     string      `json:"last_name,omitempty"`
	Country      string      `json:"country,omitempty"`
	Address      string      `json:"address,omitempty"`
	City         string      `json:"city,omitempty"`
	State        string      `json:"state,omitempty"`
	ZipCode      string      `json:"zipcode,omitempty"`
	Method       string      `json:"method,omitempty"`
	CardTarrif   TarriffType `json:"card_tarrif,omitempty"`
	MobileTarrif TarriffType `json:"mobile_tarrif,omitempty"`
}
type TarriffType string

const (
	BUSINESS_PAYS TarriffType = "BUSINESS-PAYS"
	CUSTOMER_PAYS TarriffType = "CUSTOMER-PAYS"
)

// PaymentResponse represents the API response for checkout creation
type PaymentResponse struct {
	URL         string                 `json:"url"`
	ID          string                 `json:"id"`
	Invoice     map[string]interface{} `json:"invoice"`
	CustomerID  string                 `json:"customer_id"`
	PaymentLink string                 `json:"payment_link"`
	QRCode      string                 `json:"qr_code"`
	Status      string                 `json:"status"`
	Message     string                 `json:"message"`
	Errors      map[string]interface{} `json:"errors,omitempty"`
}

// ErrorResponse represents API error response
type ErrorResponse struct {
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

// Currency constants
const (
	CurrencyKES = "KES"
	CurrencyUSD = "USD"
	CurrencyGBP = "GBP"
	CurrencyEUR = "EUR"
	CurrencyGHS = "GHS"
	CurrencyNGN = "NGN"
	CurrencyUGX = "UGX"
	CurrencyTZS = "TZS"
	CurrencyXAF = "XAF"
	CurrencyXOF = "XOF"
)

// Payment method constants
const (
	MethodMPESA = "M-PESA"
	MethodCard  = "CARD-PAYMENT"
)
// Wallet represents a wallet resource
type WalletResp struct {
	WalletID         string    `json:"wallet_id"`
	Label            string    `json:"label"`
	CanDisburse      bool      `json:"can_disburse"`
	Currency         string    `json:"currency"`
	WalletType       string    `json:"wallet_type"`
	CurrentBalance   float64   `json:"current_balance"`
	AvailableBalance float64   `json:"available_balance"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// PaginatedWallets is a standard paginated response for wallets
type PaginatedWallets struct {
	Count    int          `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []WalletResp `json:"results"`
}

type WalletRecordIDFilter string

const (
	WalletRecordIDFilterAll      WalletRecordIDFilter = ""
	WalletRecordIDFilterSpecific WalletRecordIDFilter = "specific"
)

type WalletUpdatedAtFilter string

const (
	WalletUpdatedAtFilterToday     WalletUpdatedAtFilter = "today"
	WalletUpdatedAtFilterYesterday WalletUpdatedAtFilter = "yesterday"
	WalletUpdatedAtFilterThisWeek  WalletUpdatedAtFilter = "week"
	WalletUpdatedAtFilterThisMonth WalletUpdatedAtFilter = "month"
	WalletUpdatedAtFilterThisYear  WalletUpdatedAtFilter = "year"
)

type WalletTypeFilter string

const (
	WalletTypeFilterSettlement WalletTypeFilter = "SETTLEMENT"
	WalletTypeFilterWorking    WalletTypeFilter = "WORKING"
)

// ListWalletsParams defines filters and pagination for listing wallets
type ListWalletsParams struct {
	CanDisburse *bool                 // optional
	Currency    string                // optional, one of supported currencies
	Label       string                // optional
	Page        *int                  // optional
	RecordID    WalletRecordIDFilter  // optional
	UpdatedAt   WalletUpdatedAtFilter // optional
	WalletType  WalletTypeFilter      // optional
}

func UnmarshalCollectionCallback(data []byte) (CollectionCallback, error) {
	var r CollectionCallback
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CollectionCallback) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CollectionCallback struct {
	InvoiceID      string    `json:"invoice_id"`
	State          string    `json:"state"`
	Provider       string    `json:"provider"`
	Charges        string    `json:"charges"`
	NetAmount      string    `json:"net_amount"`
	Currency       string    `json:"currency"`
	Value          string    `json:"value"`
	Account        string    `json:"account"`
	APIRef         string    `json:"api_ref"`
	Host           string    `json:"host"`
	FailedReason   string    `json:"failed_reason"`
	FailedCode     string    `json:"failed_code"`
	FailedCodeLink string    `json:"failed_code_link"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Challenge      string    `json:"challenge"`
}

type TransactionResp struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	TransactionID  string     `json:"transaction_id"`
	Invoice        *InvoiceTx `json:"invoice"`
	Currency       string     `json:"currency"`
	Value          float64    `json:"value"`
	RunningBalance float64    `json:"running_balance"`
	Narrative      string     `json:"narrative"`
	TransType      string     `json:"trans_type"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type InvoiceTx struct {
	InvoiceID      string      `json:"invoice_id"`
	State          string      `json:"state"`
	Provider       string      `json:"provider"`
	Charges        float64     `json:"charges"`
	NetAmount      string      `json:"net_amount"`
	Currency       string      `json:"currency"`
	Value          float64     `json:"value"`
	Account        string      `json:"account"`
	APIRef         string      `json:"api_ref"`
	ClearingStatus string      `json:"clearing_status"`
	MpesaReference interface{} `json:"mpesa_reference"`
	Host           string      `json:"host"`
	CardInfo       CardInfo    `json:"card_info"`
	RetryCount     int64       `json:"retry_count"`
	FailedReason   interface{} `json:"failed_reason"`
	FailedCode     interface{} `json:"failed_code"`
	FailedCodeLink interface{} `json:"failed_code_link"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type CardInfo struct {
	BinCountry *string `json:"bin_country"`
	CardType   *string `json:"card_type"`
}
