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

	resp, err := client.SendIntaSendXBPush(&intasend.IntaSendXBPushRequest{
		Amount:       "500",
		Currency:     intasend.CurrencyTZS,
		PhoneNumber:  "255755974217",
		APIRef:       "random_ref19",
		WalletID:     "YPOZ6GK",
		MobileTarrif: intasend.CUSTOMER_PAYS,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	return
	// Quick checkout for simple payments
	paymentRequest := intasend.PaymentRequest{
		PhoneNumber: "",
		FirstName:   "Felix",
		LastName:    "United",
		Email:       "k4NpI@example.com",
		Currency:    intasend.CurrencyKES,
		RedirectURL: "https://intasend.com",
		Method:      intasend.MethodCard,
		CardTarrif:  intasend.CUSTOMER_PAYS,
		APIRef:      "random_ref1",
		Comment:     "testing payment",
		Amount:      50,
	}
	response, err := client.CreateCheckoutLink(
		&paymentRequest,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Payment URL: %s\n", response.URL)
	fmt.Printf("Payment ID: %s\n", response.ID)

	/* wallets, err := client.ListWallets(&intasend.ListWalletsParams{
		WalletType: intasend.WalletTypeFilterSettlement,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(wallets)
		for count := 0; count < len(wallets.Results); count++ {
			tx := wallets.Results[count]
			fmt.Println(tx)
			// start := 0
			transactions, err := client.ListWalletTransactions(tx.WalletID, &intasend.WalletTransactionsParams{
				// Page: &start,
			})
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Println(transactions)
			}
		}
	} */

}
