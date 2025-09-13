package intasend

import (
	"encoding/json"
	"time"
)

// PaymentRequest represents the payment checkout request payload
type PaymentRequest struct {
	PhoneNumber  string            `json:"phone_number,omitempty"`
	Email        string            `json:"email,omitempty"`
	Amount       float64           `json:"amount,omitempty"`
	Currency     CurrencyType      `json:"currency,omitempty"`
	Comment      string            `json:"comment,omitempty"`
	RedirectURL  string            `json:"redirect_url,omitempty"`
	APIRef       string            `json:"api_ref,omitempty"`
	FirstName    string            `json:"first_name,omitempty"`
	LastName     string            `json:"last_name,omitempty"`
	Country      string            `json:"country,omitempty"`
	Address      string            `json:"address,omitempty"`
	City         string            `json:"city,omitempty"`
	State        string            `json:"state,omitempty"`
	ZipCode      string            `json:"zipcode,omitempty"`
	Method       PaymentMethodType `json:"method,omitempty"`
	CardTarrif   TarriffType       `json:"card_tarrif,omitempty"`
	MobileTarrif TarriffType       `json:"mobile_tarrif,omitempty"`
}
type TarriffType string

const (
	BUSINESS_PAYS TarriffType = "BUSINESS-PAYS"
	CUSTOMER_PAYS TarriffType = "CUSTOMER-PAYS"
)

// PaymentResponse represents the API response for checkout creation
type PaymentResponse struct {
	ID               string      `json:"id"`
	URL              string      `json:"url"`
	Signature        string      `json:"signature"`
	TrackingID       string      `json:"tracking_id"`
	Methods          []string    `json:"methods"`
	Layout           string      `json:"layout"`
	Styles           Styles      `json:"styles"`
	MerchantName     string      `json:"merchant_name"`
	MerchantID       string      `json:"merchant_id"`
	MerchantFullName string      `json:"merchant_full_name"`
	MerchantLogo     interface{} `json:"merchant_logo"`
	MerchantEmail    string      `json:"merchant_email"`
	MerchantOrigin   string      `json:"merchant_origin"`
	FirstName        string      `json:"first_name"`
	LastName         string      `json:"last_name"`
	PhoneNumber      interface{} `json:"phone_number"`
	Email            string      `json:"email"`
	Country          interface{} `json:"country"`
	Address          interface{} `json:"address"`
	City             interface{} `json:"city"`
	State            interface{} `json:"state"`
	Zipcode          interface{} `json:"zipcode"`
	APIRef           string      `json:"api_ref"`
	WalletID         interface{} `json:"wallet_id"`
	Method           string      `json:"method"`
	Channel          string      `json:"channel"`
	Host             string      `json:"host"`
	IsMobile         bool        `json:"is_mobile"`
	Version          interface{} `json:"version"`
	RedirectURL      string      `json:"redirect_url"`
	Amount           float64     `json:"amount"`
	Currency         string      `json:"currency"`
	Paid             bool        `json:"paid"`
	MobileTarrif     string      `json:"mobile_tarrif"`
	CardTarrif       string      `json:"card_tarrif"`
	BitcoinTarrif    string      `json:"bitcoin_tarrif"`
	AchTarrif        string      `json:"ach_tarrif"`
	BankTarrif       string      `json:"bank_tarrif"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

type Styles struct {
	ComponentBackgroundColor      string `json:"componentBackgroundColor"`
	UnselectedCardBackgroundColor string `json:"unselectedCardBackgroundColor"`
	SelectedCardBackgroundColor   string `json:"selectedCardBackgroundColor"`
	SelectedBorderColor           string `json:"selectedBorderColor"`
	UnselectedBorderColor         string `json:"unselectedBorderColor"`
	SelectedFontColor             string `json:"selectedFontColor"`
	UnselectedFontColor           string `json:"unselectedFontColor"`
	SelectedCardShadow            string `json:"selectedCardShadow"`
	UnselectedCardShadow          string `json:"unselectedCardShadow"`
	BorderRadius                  string `json:"borderRadius"`
	InputLabelColor               string `json:"inputLabelColor"`
	InputTextColor                string `json:"inputTextColor"`
	InputBackgroundColor          string `json:"inputBackgroundColor"`
	InputBorderColor              string `json:"inputBorderColor"`
	InputBorderRadius             string `json:"inputBorderRadius"`
	CtaBgColor                    string `json:"ctaBgColor"`
	CtaFontColor                  string `json:"ctaFontColor"`
	FontFamily                    string `json:"fontFamily"`
	FontWeight                    string `json:"fontWeight"`
}

// ErrorResponse represents API error response
type ErrorResponse struct {
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

type CurrencyType string

// Currency constants
const (
	CurrencyKES CurrencyType = "KES"
	CurrencyUSD CurrencyType = "USD"
	CurrencyGBP CurrencyType = "GBP"
	CurrencyEUR CurrencyType = "EUR"
	CurrencyGHS CurrencyType = "GHS"
	CurrencyNGN CurrencyType = "NGN"
	CurrencyUGX CurrencyType = "UGX"
	CurrencyTZS CurrencyType = "TZS"
	CurrencyXAF CurrencyType = "XAF"
	CurrencyXOF CurrencyType = "XOF"
)

type PaymentMethodType string

// Payment method constants
const (
	MethodMPESA PaymentMethodType = "M-PESA"
	MethodCard  PaymentMethodType = "CARD-PAYMENT"
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
