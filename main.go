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
	if c.Ssl.Enabled{
		log.Printf("[INFO] Starting https API :yay:")
		log.Fatal(http.ListenAndServeTLS(":"+c.ApiPort, c.Ssl.CertFile, c.Ssl.KeyFile, router))
	}else{
		log.Printf("[INFO] Starting http API :yay:")
		log.Printf("[WARN] You are using an insecure connection")
		log.Fatal(http.ListenAndServe(":"+c.ApiPort, router))
	}
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
