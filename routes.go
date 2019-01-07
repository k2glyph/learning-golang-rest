package main

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc HandlerWrap
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		HandlerWrap{hello},
	},
	Route{
		"POST",
		"/payment/paypal",
		HandlerWrap{PaymentFromPaypal},
	},
	Route{
		"POST",
		"/payment/stripe",
		HandlerWrap{PaymentFromStripe},
	},
}
