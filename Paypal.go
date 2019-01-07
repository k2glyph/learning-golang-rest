package main

import (
	paypalsdk "github.com/logpacker/PayPal-Go-SDK"
)

func DirectPaypalPayment() {
	c, err := paypalsdk.NewClient("AZfxEQ2LMtKTtkmHzUXIVcFrN_ddV0J-7EQ1SQc5Z1OrqqEB7RzxcMWKNrsOr-JO9C4H1IawgOyPhCQe", "secretID", paypalsdk.APIBaseSandBox)
	if err != nil {
		panic(err)
	}

	// Retrieve access token
	_, err = c.GetAccessToken()
	if err != nil {
		panic(err)
	}

	// Create credit card payment
	p := paypalsdk.Payment{
		Intent: "sale",
		Payer: &paypalsdk.Payer{
			PaymentMethod: "credit_card",
			FundingInstruments: []paypalsdk.FundingInstrument{{
				CreditCard: &paypalsdk.CreditCard{
					Number:      "4111111111111111",
					Type:        "visa",
					ExpireMonth: "11",
					ExpireYear:  "2020",
					CVV2:        "777",
					FirstName:   "John",
					LastName:    "Doe",
				},
			}},
		},
		Transactions: []paypalsdk.Transaction{{
			Amount: &paypalsdk.Amount{
				Currency: "USD",
				Total:    "7.00",
			},
			Description: "My Payment",
		}},
		RedirectURLs: &paypalsdk.RedirectURLs{
			ReturnURL: "http://...",
			CancelURL: "http://...",
		},
	}
	_, err = c.CreatePayment(p)
	if err != nil {
		panic(err)
	}
}
