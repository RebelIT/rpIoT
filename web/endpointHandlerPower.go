package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)

func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	if !config.ApiConfig.EndpointPower{
		returnErr(w,r,http.StatusNotFound)
		return
	}

	status, err := actions.SystemPower(action)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkPower(w,r,status)
}