package main

import (
	"log"
	"net/http"
)

const CONFIG = "./api_config.json"

func main(){
	router := NewRouter()
	log.Printf("Starting API :yay:")
	log.Fatal(http.ListenAndServe(":6661", router))
}
