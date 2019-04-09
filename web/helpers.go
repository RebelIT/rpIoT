package web

import (
	"encoding/json"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/config"
	"log"
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

func returnUnauthorized(w http.ResponseWriter, r *http.Request, resp Auth, err error){
	common.SendMetric(r.URL.Path, http.StatusUnauthorized, r.Method)

	resp.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}


// helpers per handler function
func returnOkAlive(w http.ResponseWriter, r *http.Request, resp Alive){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkHdmiState(w http.ResponseWriter, r *http.Request, resp actions.Display){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkGpioPinState(w http.ResponseWriter, r *http.Request, resp actions.Pin){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkAllGpioPinStates(w http.ResponseWriter, r *http.Request, resp actions.GpioStates){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkSystemInfo(w http.ResponseWriter, r *http.Request, resp actions.Sysinfo){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkAptLog(w http.ResponseWriter, r *http.Request, resp actions.AptLog){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkAptInstallUpdate(w http.ResponseWriter, r *http.Request, resp actions.AptInstallUpdate){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkPower(w http.ResponseWriter, r *http.Request, resp actions.Power){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}

func returnOkService(w http.ResponseWriter, r *http.Request, resp actions.Service){
	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
	return
}