package main

import (
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/web"
	"log"
	"net/http"
)

func main(){
	c, err := common.ReadConfig()
	if err != nil{
		log.Printf("[PANIC] Config is required and there is an error : %s", err)
		panic(err)
	}

	router := web.NewRouter()
	msg := "[INFO] Starting API on "+common.GetHostname()+" :yay:"
	if c.Ssl.Enabled{
		log.Printf(msg)
		common.SendSlack(msg)
		log.Fatal(http.ListenAndServeTLS(":"+c.ApiPort, common.APPDIR+c.Ssl.CertFile, common.APPDIR+c.Ssl.KeyFile, router))
	}else{
		log.Printf(msg)
		log.Printf("[WARN] You are using an insecure connection")
		common.SendSlack(msg)
		log.Fatal(http.ListenAndServe(":"+c.ApiPort, router))
	}
}