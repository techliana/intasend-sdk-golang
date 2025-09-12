package intasend

import (
	"fmt"
	"net/url"
	"strconv"
)

// ListWallets retrieves a paginated list of wallets with optional filters
func (c *Client) ListWallets(params *ListWalletsParams) (*PaginatedWallets, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to list wallets")
	}

	// Build query parameters
	queryParams := make(map[string]string)
	if params != nil {
		if params.CanDisburse != nil {
			queryParams["can_disburse"] = strconv.FormatBool(*params.CanDisburse)
		}
		if params.Currency != "" {
			queryParams["currency"] = params.Currency
		}
		if params.Label != "" {
			queryParams["label"] = params.Label
		}
		if params.Page != nil {
			queryParams["page"] = strconv.Itoa(*params.Page)
		}
		if params.RecordID != "" {
			queryParams["record_id"] = string(params.RecordID)
		}
		if params.UpdatedAt != "" {
			queryParams["updated_at"] = string(params.UpdatedAt)
		}
		if params.WalletType != "" {
			queryParams["wallet_type"] = string(params.WalletType)
		}
	}

	// Use the HTTP wrapper to make the request
	var result PaginatedWallets
	err := c.GetJSON("api/v1/wallets/", queryParams, true, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// WalletTransactionsParams defines optional pagination for wallet transactions
type WalletTransactionsParams struct {
	Page *int // optional page number
}

// ListWalletTransactions retrieves transactions performed under a specific wallet
func (c *Client) ListWalletTransactions(walletID string, params *WalletTransactionsParams) (*TransactionResp, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("token is required to list wallet transactions")
	}
	if walletID == "" {
		return nil, fmt.Errorf("walletID is required")
	}

	// Build query params
	queryParams := make(map[string]string)
	if params != nil && params.Page != nil {
		queryParams["page"] = strconv.Itoa(*params.Page)
	}

	// Use the common HTTP JSON helper
	endpoint := fmt.Sprintf("api/v1/wallets/%s/transactions/", url.PathEscape(walletID))
	var txns TransactionResp
	if err := c.GetJSON(endpoint, queryParams, true, &txns); err != nil {
		return nil, err
	}
	return &txns, nil
}
