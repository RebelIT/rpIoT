package main

import (
	"flag"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/config"
	"github.com/rebelit/rpIoT/web"
	"log"
	"net/http"
)

func main(){
	f := flag.String("config", "", "path to config.json to override default settings.")
	flag.Parse()

	config.DoApiSettings(*f)

	router := web.NewRouter()
	msg := "[INFO] Starting API on "+common.GetHostname()+" on port:"+config.ApiConfig.ApiPort+" :yay:"
	if config.ApiConfig.SslEnabled{
		log.Printf(msg)
		common.SendSlack(msg)
		log.Fatal(http.ListenAndServeTLS(":"+config.ApiConfig.ApiPort, config.ApiConfig.SslCertName, config.ApiConfig.SslKeyName, router))
	}else{
		log.Printf(msg)
		log.Printf("[WARN] You are using an insecure connection")
		common.SendSlack(msg)
		log.Fatal(http.ListenAndServe(":"+config.ApiConfig.ApiPort, router))
	}
}