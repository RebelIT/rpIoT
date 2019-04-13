package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)

func getService(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	svc := vars["service"]

	status, err := actions.GetService(svc)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkService(w,r,status)
	return
}

func serviceAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]

	if !config.ApiConfig.EndpointService{
		returnErr(w,r,http.StatusNotFound)
		return
	}

	resp, err := actions.ServiceControl(service, action)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkService(w,r,resp)
	return
}