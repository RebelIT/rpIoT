package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)
func displayGet (w http.ResponseWriter, r *http.Request) {
	state, err := actions.GetHdmiState()
	if err != nil{
		returnErr(w,r, http.StatusInternalServerError)
		return
	}

	returnOkHdmiState(w,r,state)
	return
}

func displayAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	if !config.ApiConfig.EndpointDisplay{
		returnErr(w,r,http.StatusNotFound)
		return
	}

	state, err := actions.HdmiToggleState(action)
	if err != nil{
		returnErr(w,r, http.StatusInternalServerError)
		return
	}

	returnOkHdmiState(w,r,state)
	return
}