package main

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func PaymentStripe(request *RequestStripe) (*stripe.Charge, error) {
	chargeParams := &stripe.ChargeParams{
		Amount:      request.Amount,
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: request.Description,
	}
	if len(request.SourceID) > 0 {
		chargeParams.SetSource(request.SourceID)
	}
	charge, err := charge.New(chargeParams)
	if err != nil {
		fmt.Println("Could not process payment: ", err)
		return nil, err
	}
	return charge, nil
}

// SetStripeKey - Helper func to set our Stripe key for auth
func SetStripeKey(stripeKey string) string {
	// set our stripe SECRET key to auth with the API
	stripe.Key = stripeKey
	return stripe.Key
}
