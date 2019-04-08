package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/config"
	"log"
	"net/http"
)

//Http Generic Response functions
func returnOk(w http.ResponseWriter, r *http.Request, resp Response){
	code := http.StatusOK
	common.SendMetric(r.URL.Path, code, r.Method)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnBad(w http.ResponseWriter, r *http.Request, resp Response, err error){
	code := http.StatusBadRequest
	common.SendMetric(r.URL.Path, code, r.Method)

	resp.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnDisabled(w http.ResponseWriter, r *http.Request, resp Response){
	code := http.StatusForbidden
	common.SendMetric(r.URL.Path, code, r.Method)

	resp.Message = "endpoint is disabled"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnInternalError(w http.ResponseWriter, r *http.Request, resp Response, err error){
	code := http.StatusInternalServerError
	common.SendMetric(r.URL.Path, code, r.Method)

	resp.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func returnUnauthorized(w http.ResponseWriter, r *http.Request, resp Response, err error){
	code := http.StatusUnauthorized
	common.SendMetric(r.URL.Path, code, r.Method)

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

func getSystemStats (w http.ResponseWriter, r *http.Request){
	resp := actions.Sysinfo{}

	resp.Host = actions.GetHostStats()
	resp.Cpu = actions.GetCpuStats()
	resp.Mem = actions.GetMemStats()
	resp.Disk = actions.GetDiskStats()

	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

func getUpdates (w http.ResponseWriter, r *http.Request){	//Returns everything in apt/history.log
	resp := Response{}
	resp.Namespace = r.URL.Path

	data, err  := actions.CheckUpdateLog()
	if err != nil{
		returnInternalError(w, r, resp, err)
	}

	resp.Message = "ok"
	resp.Data = data
	returnOk(w,r,resp)
}

func getService(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	svc := vars["service"]

	resp := Response{}
	resp.Namespace = r.URL.Path

	status, err := actions.GetService(svc)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = status
	returnOk(w,r,resp)
	return
}

func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if !config.ApiConfig.EndpointPower{
		returnDisabled(w,r,resp)
		return
	}

	err := actions.SystemPower(action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "ok"
	returnOk(w,r,resp)
}

func updateAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if !config.ApiConfig.EndpointPower{
		returnDisabled(w,r,resp)
		return
	}

	cmdOut, err := actions.Update(action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}

func installAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	pkg := vars["package"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if !config.ApiConfig.EndpointPackages{
		returnDisabled(w,r,resp)
		return
	}

	cmdOut, err := actions.Install(action, pkg)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}

func serviceAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if !config.ApiConfig.EndpointService{
		returnDisabled(w,r,resp)
		return
	}

	cmdOut, err := actions.ServiceControl(service, action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}