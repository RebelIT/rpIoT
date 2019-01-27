package main

import (
	"github.com/rebelit/rpIoT/common"
	"log"
	"net/http"
)

func main(){
	c, err := common.ReadConfig()
	if err != nil{
		log.Printf("[PANIC] Config is required and there is an error : %s", err)
		panic(err)
	}
	router := NewRouter()
	log.Printf("Starting API :yay:")
	log.Fatal(http.ListenAndServe(":"+c.ApiPort, router))
}

//Main structs
//HTTP router
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
