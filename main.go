package main

import (
	"log"
	"net/http"
)

const CONFIG = "/etc/api/api_config.json"

func main(){
	c, err := readConfig()
	if err != nil{
		log.Printf("[PANIC] Config is required and there is an error : %s", err)
		panic(err)
	}
	router := NewRouter()
	log.Printf("Starting API :yay:")
	log.Fatal(http.ListenAndServe(":"+c.ApiPort, router))
}
