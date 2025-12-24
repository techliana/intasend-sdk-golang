package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/techliana/intasend-sdk-golang"
)

// challenge:="intasendwebhook"
func main() {
	// load environment variables
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	publishable := os.Getenv("PUBLISHABLE_KEY")
	token := os.Getenv("TOKEN")
	client := intasend.NewClient(publishable, token, false, true)

	// resp, err := client.SendIntaSendXBPush(&intasend.IntaSendXBPushRequest{
	// 	Amount:       "500",
	// 	Currency:     intasend.CurrencyTZS,
	// 	PhoneNumber:  "255755974217",
	// 	APIRef:       "random_ref19",
	// 	WalletID:     "YPOZ6GK",
	// 	MobileTarrif: intasend.CUSTOMER_PAYS,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	// return
	// // Quick checkout for simple payments
	// paymentRequest := intasend.PaymentRequest{
	// 	PhoneNumber: "",
	// 	FirstName:   "Felix",
	// 	LastName:    "United",
	// 	Email:       "k4NpI@example.com",
	// 	Currency:    intasend.CurrencyKES,
	// 	RedirectURL: "https://intasend.com",
	// 	Method:      intasend.MethodCard,
	// 	CardTarrif:  intasend.CUSTOMER_PAYS,
	// 	APIRef:      "random_ref1",
	// 	Comment:     "testing payment",
	// 	Amount:      50,
	// }
	// response, err := client.CreateCheckoutLink(
	// 	&paymentRequest,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Payment URL: %s\n", response.URL)
	// fmt.Printf("Payment ID: %s\n", response.ID)

	// Example 1: List all invoices
	invoices, err := client.ListInvoices(nil)
	if err != nil {
		fmt.Printf("Error listing invoices: %s\n", err)
	} else {
		fmt.Printf("Total invoices: %d\n", invoices.Count)
		for _, invoice := range invoices.Results {
			fmt.Printf("Invoice ID: %s, State: %s, Net Amount: %s %s, Value: %.2f, Charges: %.2f\n",
				invoice.InvoiceID, invoice.State, invoice.NetAmount, invoice.Currency,
				invoice.Value, invoice.Charges)
		}
	}

	// Example 2: List invoices with filters
	page := 1
	pageSize := 10
	filteredInvoices, err := client.ListInvoices(&intasend.ListInvoicesParams{
		Page:     &page,
		PageSize: &pageSize,
		// State:    "COMPLETED",
		// Currency: "KES",
	})
	if err != nil {
		fmt.Printf("Error listing filtered invoices: %s\n", err)
	} else {
		fmt.Printf("Completed KES invoices: %d\n", filteredInvoices.Count)
	}

	// Example 3: Get a specific invoice by ID
	// Replace "INVOICE_ID_HERE" with an actual invoice ID from your account
	// invoice, err := client.GetInvoice("INVOICE_ID_HERE")
	// if err != nil {
	// 	fmt.Printf("Error getting invoice: %s\n", err)
	// } else {
	// 	fmt.Printf("Invoice: %+v\n", invoice)
	// 	if invoice.IsCompleted() {
	// 		fmt.Println("Invoice is completed!")
	// 	}
	// 	if invoice.IsFailed() {
	// 		fmt.Printf("Invoice failed: %s\n", invoice.GetFailureReason())
	// 	}
	// }

	// Example 4: List all transactions
	// transactions, err := client.ListTransactions(nil)
	// if err != nil {
	// 	fmt.Printf("Error listing transactions: %s\n", err)
	// } else {
	// 	fmt.Printf("Total transactions: %d\n", transactions.Count)
	// 	for _, tx := range transactions.Results {
	// 		fmt.Printf("Transaction ID: %s, Type: %s, Status: %s, Amount: %.2f %s\n",
	// 			tx.TransactionID, tx.TransType, tx.Status, tx.Value, tx.Currency)
	// 		if tx.Invoice != nil {
	// 			fmt.Printf("  Invoice ID: %s\n", tx.Invoice.InvoiceID)
	// 		}
	// 	}
	// }

	// Example 5: List transactions with filters
	// txPage := 1
	// txPageSize := 20
	// filteredTxns, err := client.ListTransactions(&intasend.ListTransactionsParams{
	// 	Page:      &txPage,
	// 	PageSize:  &txPageSize,
	// 	Currency:  "KES",
	// 	TransType: intasend.TransTypeDeposit,
	// 	Status:    intasend.TransStatusCompleted,
	// })
	// if err != nil {
	// 	fmt.Printf("Error listing filtered transactions: %s\n", err)
	// } else {
	// 	fmt.Printf("Filtered transactions: %d\n", filteredTxns.Count)
	// 	for _, tx := range filteredTxns.Results {
	// 		if tx.IsCompleted() && tx.IsDeposit() {
	// 			fmt.Printf("Completed deposit: %s - %.2f %s\n",
	// 				tx.TransactionID, tx.Value, tx.Currency)
	// 		}
	// 	}
	// }

	// Example 6: Get a specific transaction by ID
	// transaction, err := client.GetTransaction("TRANSACTION_ID_HERE")
	// if err != nil {
	// 	fmt.Printf("Error getting transaction: %s\n", err)
	// } else {
	// 	fmt.Printf("Transaction: %+v\n", transaction)
	// 	fmt.Printf("Running Balance: %.2f\n", transaction.RunningBalance)
	// 	if transaction.IsCompleted() {
	// 		fmt.Println("✓ Transaction completed")
	// 	}
	// }

	// Example 7: List transactions by wallet ID
	// walletTxns, err := client.ListTransactions(&intasend.ListTransactionsParams{
	// 	WalletID: "YOUR_WALLET_ID",
	// })
	// if err != nil {
	// 	fmt.Printf("Error listing wallet transactions: %s\n", err)
	// } else {
	// 	fmt.Printf("Wallet transactions: %d\n", walletTxns.Count)
	// }

	// Example 8: List transactions by date range
	// dateTxns, err := client.ListTransactions(&intasend.ListTransactionsParams{
	// 	DateFrom: "2024-01-01",
	// 	DateTo:   "2024-12-31",
	// 	Currency: "KES",
	// })
	// if err != nil {
	// 	fmt.Printf("Error listing transactions by date: %s\n", err)
	// } else {
	// 	fmt.Printf("Transactions in date range: %d\n", dateTxns.Count)
	// }

	// Example 9: List wallets (existing code)
	// wallets, err := client.ListWallets(&intasend.ListWalletsParams{
	// 	WalletType: intasend.WalletTypeFilterSettlement,
	// })
	// if err != nil {
	// 	fmt.Printf("Error: %s\n", err)
	// } else {
	// 	fmt.Println(wallets)
	// 	for count := 0; count < len(wallets.Results); count++ {
	// 		tx := wallets.Results[count]
	// 		fmt.Println(tx)
	// 		// start := 0
	// 		transactions, err := client.ListWalletTransactions(tx.WalletID, &intasend.WalletTransactionsParams{
	// 			// Page: &start,
	// 		})
	// 		if err != nil {
	// 			fmt.Printf("Error: %s\n", err)
	// 		} else {
	// 			fmt.Println(transactions)
	// 		}
	// 	}
	// }

}
