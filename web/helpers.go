package web

import (
	"encoding/json"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)

func validateAuth(username string,token string)(bool){
	if config.ApiConfig.DefaultUser == username && config.ApiConfig.DefaultUserToken == token{
		return true
	}

	return false
}

// helpers for all http handlers
func returnErr(w http.ResponseWriter, r *http.Request, status int){
	common.SendMetric(r.URL.Path, status, r.Method)
	w.WriteHeader(status)
	return
}

// helpers per handler function
func returnOkHdmiState(w http.ResponseWriter, r *http.Request, resp actions.Display){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkGpioPinState(w http.ResponseWriter, r *http.Request, resp actions.Pin){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkAllGpioPinStates(w http.ResponseWriter, r *http.Request, resp actions.GpioStates){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
}