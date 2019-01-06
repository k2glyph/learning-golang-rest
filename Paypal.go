package main

import (
	"fmt"

	paypalsdk "github.com/logpacker/PayPal-Go-SDK"
)

func DirectPaypalPayment(c *gin.Context) {
	client, err := SetAccessTokenForClient(ClientID, SecretID)
	if err != nil {
		panic(err)
		return
	}
	// create direct paypal payment
	amount := paypalsdk.Amount{
		Total:    "7.00",
		Currency: "USD",
	}
	redirectURI := "http://example.com/redirect-uri"
	cancelURI := "http://example.com/cancel-uri"
	description := "Description for direct Paypal payment"
	paymentResult, err := client.CreateDirectPaypalPayment(amount, redirectURI, cancelURI, description)

	//fmt.Println(paymentResult)
	c.JSON(200, gin.H{
		Response: paymentResult,
	})

	// Execute approved payment
	paymentID := paymentResult.ID
	fmt.Println(paymentID)
	// payerID := paymentResult.payer_id
	// executeResult, err := c.ExecuteApprovedPayment(paymentID, payerID)

}
func SetAccessTokenForClient(clientId, secretId string) (client *paypalsdk.Client, err interface{}) {
	client, err = paypalsdk.NewClient(clientId, secretId, paypalsdk.APIBaseSandBox)
	if err != nil {
		return client, err
	}
	accessToken, err := client.GetAccessToken()
	if err != nil {
		return client, err
	}
	client.SetAccessToken(accessToken.Token)
	return client, err
}
