package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Namespace handlers
func getStatus (w http.ResponseWriter, r *http.Request){
	resp := Response{}
	resp.Namespace = r.URL.Path
	resp.Message = "alive"
	resp.Timestamp = time.Now()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig(action)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := command(string(action), []string{"+1"}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}

	outputs := Response{}
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "success"
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func yumAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig(action)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := command(string("apt-get"), []string{action , "-y"}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}

	outputs := Response{}
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "success"
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func serviceAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]

	enabled, err := checkConfig("service")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := validateServiceAction(action); err != nil{
		log.Printf("[WARN] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := command(string("service"), []string{service , action}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}

	outputs := Response{}
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "success"
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func displayAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig("display")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	newState, err := validateDisplayAction(action)
	if err != nil{
		log.Printf("[WARN] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := command(string("vcgencmd"), []string{"display_power", newState}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}

	outputs := Response{}
	outputs.Namespace = string(r.URL.Path)
	outputs.Message = "success"
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func gpioSwitch (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	enabled, err := checkConfig("gpio")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	if err := gpioToggle(pin); err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
	}

	outputs := Response{}
	outputs.Message = "success"
	outputs.Namespace = string(r.URL.Path)
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func gpioPullDown (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	enabled, err := checkConfig("gpio")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	pinState, err := gpioDown(pin)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
	}

	outputs := Response{}
	outputs.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	outputs.Namespace = string(r.URL.Path)
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}

func gpioPullUp (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	enabled, err := checkConfig("gpio")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		sendMetric(r.URL.Path,http.StatusNotFound, r.Method)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		sendMetric(r.URL.Path,http.StatusBadRequest, r.Method)
		return
	}

	pinState, err := gpioUp(pin)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
	}

	outputs := Response{}
	outputs.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	outputs.Namespace = string(r.URL.Path)
	outputs.Timestamp = time.Time(time.Now())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(outputs); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		sendMetric(r.URL.Path,http.StatusInternalServerError, r.Method)
		return
	}
	sendMetric(r.URL.Path,http.StatusOK, r.Method)
}
