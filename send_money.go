package intasend

import (
	"fmt"
	"time"
)

// SendMoneyProvider represents the payment provider for send-money
type SendMoneyProvider string

const (
	ProviderMPESAB2C    SendMoneyProvider = "MPESA-B2C"
	ProviderMPESAB2B    SendMoneyProvider = "MPESA-B2B"
	ProviderBankTransfer SendMoneyProvider = "BANK"
	ProviderIntaSendXB  SendMoneyProvider = "INTASEND-XB"
)

// SendMoneyAccountType represents the account type for a transaction
type SendMoneyAccountType string

const (
	AccountTypeTillNumber  SendMoneyAccountType = "TillNumber"
	AccountTypePaybill     SendMoneyAccountType = "PayBill"
	AccountTypeBankAccount SendMoneyAccountType = "BankAccount"
	AccountTypePhone       SendMoneyAccountType = "PhoneNumber"
)

// RequiresApproval indicates whether the batch requires manual approval
type RequiresApproval string

const (
	ApprovalYes RequiresApproval = "YES"
	ApprovalNo  RequiresApproval = "NO"
)

// SendMoneyTransaction represents a single recipient in a send-money batch
type SendMoneyTransaction struct {
	Name             string               `json:"name"`
	Account          string               `json:"account"`
	IDNumber         string               `json:"id_number,omitempty"`
	Amount           string               `json:"amount"`
	BankCode         string               `json:"bank_code,omitempty"`
	CategoryName     string               `json:"category_name,omitempty"`
	Narrative        string               `json:"narrative,omitempty"`
	AccountType      SendMoneyAccountType `json:"account_type,omitempty"`
	AccountReference string               `json:"account_reference,omitempty"`
}

// SendMoneyRequest represents the payload for initiating a send-money batch
type SendMoneyRequest struct {
	Currency        CurrencyType          `json:"currency"`
	Provider        SendMoneyProvider     `json:"provider"`
	DeviceID        string                `json:"device_id,omitempty"`
	CallbackURL     string                `json:"callback_url,omitempty"`
	BatchReference  string                `json:"batch_reference,omitempty"`
	Transactions    []SendMoneyTransaction `json:"transactions"`
	Country         string                `json:"country,omitempty"`
	RequiresApproval RequiresApproval     `json:"requires_approval,omitempty"`
}

// SendMoneyTransactionStatus represents the status of a single transaction in the response
type SendMoneyTransactionStatus struct {
	TransactionID      string      `json:"transaction_id"`
	Status             string      `json:"status"`
	StatusCode         string      `json:"status_code"`
	StatusDescription  string      `json:"status_description"`
	RequestReferenceID string      `json:"request_reference_id"`
	Name               string      `json:"name"`
	Account            string      `json:"account"`
	IDNumber           interface{} `json:"id_number"`
	BankCode           interface{} `json:"bank_code"`
	Amount             float64     `json:"amount"`
	Narrative          string      `json:"narrative"`
	IdempotencyKey     interface{} `json:"idempotency_key"`
}

// SendMoneyResponse represents the API response for initiating a send-money batch
type SendMoneyResponse struct {
	FileID               string                       `json:"file_id"`
	DeviceID             interface{}                  `json:"device_id"`
	TrackingID           string                       `json:"tracking_id"`
	BatchReference       string                       `json:"batch_reference"`
	Status               string                       `json:"status"`
	StatusCode           string                       `json:"status_code"`
	Nonce                string                       `json:"nonce"`
	Wallet               WalletResp                   `json:"wallet"`
	Transactions         []SendMoneyTransactionStatus `json:"transactions"`
	ChargeEstimate       float64                      `json:"charge_estimate"`
	TotalAmountEstimate  float64                      `json:"total_amount_estimate"`
	TotalAmount          float64                      `json:"total_amount"`
	TransactionsCount    int                          `json:"transactions_count"`
	CreatedAt            time.Time                    `json:"created_at"`
	UpdatedAt            time.Time                    `json:"updated_at"`
}

// InitiateSendMoney initiates a send-money (disbursement) batch
func (c *Client) InitiateSendMoney(req *SendMoneyRequest) (*SendMoneyResponse, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to initiate send-money")
	}
	if req == nil {
		return nil, fmt.Errorf("request payload is required")
	}
	if req.Provider == "" {
		return nil, fmt.Errorf("provider is required")
	}
	if req.Currency == "" {
		return nil, fmt.Errorf("currency is required")
	}
	if len(req.Transactions) == 0 {
		return nil, fmt.Errorf("at least one transaction is required")
	}

	endpoint := fmt.Sprintf("%s/api/v1/send-money/initiate/", c.APIBaseURL)
	var resp SendMoneyResponse
	if err := c.PostJSON(endpoint, req, true, false, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
