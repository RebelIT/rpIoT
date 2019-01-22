package main

import (
	"net/http"
	"time"
)

//HTTP router
type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//API response for all 200 ok
type Response struct{
	Namespace	string		`json:"namespace"`
	Message		string		`json:"message"`
	Timestamp	time.Time	`json:"timestamp"`
}

//Config File functions toggle on/off
type Function struct{
	Name		string `json:"name"`
	Enabled 	bool `json:"enabled"`
}

type Config struct{
	ApiPort	string `json:"api_port"`
	Functions		[]Function
}