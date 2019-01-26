package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)

//Http Response helper functions
func returnOk(w http.ResponseWriter, r *http.Request, resp Response){
	code := http.StatusOK
	sendMetric(r.URL.Path, code, r.Method)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnBad(w http.ResponseWriter, r *http.Request, resp Response, err error){
	code := http.StatusBadRequest
	sendMetric(r.URL.Path, code, r.Method)

	resp.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnInternalError(w http.ResponseWriter, r *http.Request, resp Response, err error){
	code := http.StatusInternalServerError
	sendMetric(r.URL.Path, code, r.Method)

	resp.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

//Namespace handlers
func getStatus (w http.ResponseWriter, r *http.Request){
	resp := Response{}
	resp.Namespace = r.URL.Path
	resp.Message = "alive"

	returnOk(w,r,resp)
}

func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig(action)
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	if err := command(string(action), []string{"+1"}); err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func yumAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig(action)
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	if err := command(string("apt-get"), []string{action , "-y"}); err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func serviceAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig("service")
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	if err := validateServiceAction(action); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := command(string("service"), []string{service, action}); err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func displayAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig("display")
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	newState, err := validateDisplayAction(action)
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := command(string("vcgencmd"), []string{"display_power", newState}); err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func gpioSwitch (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)


	enabled, err := checkConfig("gpio")
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	if err := gpioToggle(pin); err != nil{
		returnInternalError(w,r,resp, err)
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func gpioPullDown (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig("gpio")
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	pinState, err := gpioDown(pin)
	if err != nil{
		returnInternalError(w,r,resp, err)
	}

	resp.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	returnOk(w,r,resp)
}

func gpioPullUp (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	enabled, err := checkConfig("gpio")
	if err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if !enabled{
		returnBad(w,r,resp, errors.New("API function disabled"))
		return
	}

	pinState, err := gpioUp(pin)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	returnOk(w,r,resp)
}