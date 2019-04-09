package web

import (
	"github.com/rebelit/rpIoT/actions"
	"net/http"
)

func getSystemStats (w http.ResponseWriter, r *http.Request){
	resp, _ := actions.GetSystemInfo()
	returnOkSystemInfo(w,r,resp)
	return
}

func getStatus (w http.ResponseWriter, r *http.Request){
	resp := Alive{}
	resp.Status = "im running fine..."

	returnOkAlive(w,r,resp)
}