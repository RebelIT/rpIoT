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
type Alive struct{
	Status	string		`json:"status"`
}

type Auth struct{
	Message		string		`json:"message"`
}