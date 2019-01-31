package web

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//API response
type Response struct{
	Namespace	string		`json:"namespace"`
	Message		string		`json:"message"`
}

//Auth mux middleware
type Auth struct {
	TokenUsers map[string]string
}