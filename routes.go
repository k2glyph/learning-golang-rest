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
}
