package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
}

func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig(action)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := command(string(action), []string{"+1"}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}
}

func yumAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig(action)
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := command(string("apt-get"), []string{action , "-y"}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}
}

func serviceAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]

	enabled, err := checkConfig("service")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validateServiceAction(action); err != nil{
		log.Printf("[WARN] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := command(string("service"), []string{service , action}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}
}

func displayAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	enabled, err := checkConfig("display")
	if err != nil{
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !enabled{
		log.Printf("[WARN] %s : %s", r.URL.Path, "API function disabled")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newState, err := validateDisplayAction(action)
	if err != nil{
		log.Printf("[WARN] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := command(string("vcgencmd"), []string{"display_power", newState}); err != nil{
		log.Printf("[ERROR] %s : %s", r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}
}
