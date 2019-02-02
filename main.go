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
	if c.Ssl.Enabled{
		log.Printf("[INFO] Starting https API :yay:")
		log.Fatal(http.ListenAndServeTLS(":"+c.ApiPort, common.APPDIR+c.Ssl.CertFile, common.APPDIR+c.Ssl.KeyFile, router))
	}else{
		log.Printf("[INFO] Starting http API :yay:")
		log.Printf("[WARN] You are using an insecure connection")
		log.Fatal(http.ListenAndServe(":"+c.ApiPort, router))
	}
}