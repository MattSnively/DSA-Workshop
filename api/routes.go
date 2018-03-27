package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/{test}",
		Index,
	},
	Route{
		"GetMetrics",
		"GET",
		"/metrics/{series}",
		QueryInfluxDB,
	},
}
