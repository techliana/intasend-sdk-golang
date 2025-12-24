package intasend

import (
	"fmt"
	"net/url"
	"strconv"
)

// ListTransactionsParams defines optional filters and pagination for listing transactions
type ListTransactionsParams struct {
	Page       *int   // optional page number
	PageSize   *int   // optional page size
	WalletID   string // optional filter by wallet ID
	Currency   string // optional filter by currency (KES, USD, etc.)
	TransType  string // optional filter by transaction type (DEPOSIT, WITHDRAWAL, etc.)
	Date       string // optional filter by date
	DateFrom   string // optional filter by start date (YYYY-MM-DD)
	DateTo     string // optional filter by end date (YYYY-MM-DD)
	Status     string // optional filter by status
	RecordID   string // optional filter by record ID
	UpdatedAt  string // optional filter by updated_at (today, yesterday, week, month, year)
}

// Transaction type constants
const (
	TransTypeDeposit    = "DEPOSIT"
	TransTypeWithdrawal = "WITHDRAWAL"
	TransTypeTransfer   = "TRANSFER"
	TransTypeCharge     = "CHARGE"
	TransTypeRefund     = "REFUND"
)

// Transaction status constants
const (
	TransStatusPending   = "PENDING"
	TransStatusCompleted = "COMPLETED"
	TransStatusFailed    = "FAILED"
)

// ListTransactions retrieves a paginated list of all transactions with optional filters
func (c *Client) ListTransactions(params *ListTransactionsParams) (*TransactionResp, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to list transactions")
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
		if params.WalletID != "" {
			queryParams["wallet_id"] = params.WalletID
		}
		if params.Currency != "" {
			queryParams["currency"] = params.Currency
		}
		if params.TransType != "" {
			queryParams["trans_type"] = params.TransType
		}
		if params.Date != "" {
			queryParams["date"] = params.Date
		}
		if params.DateFrom != "" {
			queryParams["date_from"] = params.DateFrom
		}
		if params.DateTo != "" {
			queryParams["date_to"] = params.DateTo
		}
		if params.Status != "" {
			queryParams["status"] = params.Status
		}
		if params.RecordID != "" {
			queryParams["record_id"] = params.RecordID
		}
		if params.UpdatedAt != "" {
			queryParams["updated_at"] = params.UpdatedAt
		}
	}

	// Use the HTTP wrapper to make the request
	var result TransactionResp
	err := c.GetJSON("api/v1/transactions/", queryParams, true, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransaction retrieves a single transaction by its ID
func (c *Client) GetTransaction(transactionID string) (*Result, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to get transaction")
	}
	if transactionID == "" {
		return nil, fmt.Errorf("transactionID is required")
	}

	// Use the common HTTP JSON helper
	endpoint := fmt.Sprintf("api/v1/transactions/%s/", url.PathEscape(transactionID))
	var transaction Result
	if err := c.GetJSON(endpoint, nil, true, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Helper methods for Result (Transaction)

// IsCompleted checks if the transaction is completed
func (r *Result) IsCompleted() bool {
	return r.Status == TransStatusCompleted
}

// IsPending checks if the transaction is pending
func (r *Result) IsPending() bool {
	return r.Status == TransStatusPending
}

// IsFailed checks if the transaction failed
func (r *Result) IsFailed() bool {
	return r.Status == TransStatusFailed
}

// IsDeposit checks if the transaction is a deposit
func (r *Result) IsDeposit() bool {
	return r.TransType == TransTypeDeposit
}

// IsWithdrawal checks if the transaction is a withdrawal
func (r *Result) IsWithdrawal() bool {
	return r.TransType == TransTypeWithdrawal
}

// IsTransfer checks if the transaction is a transfer
func (r *Result) IsTransfer() bool {
	return r.TransType == TransTypeTransfer
}

// GetInvoiceID returns the invoice ID if the transaction has an associated invoice
func (r *Result) GetInvoiceID() string {
	if r.Invoice != nil {
		return r.Invoice.InvoiceID
	}
	return ""
}

