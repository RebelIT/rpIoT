package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/apt"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/gpio"
	"github.com/rebelit/rpIoT/hdmi"
	"github.com/rebelit/rpIoT/pwr"
	"github.com/rebelit/rpIoT/svc"
	"github.com/rebelit/rpIoT/sys"
	"log"
	"net/http"
	"strconv"
)

//Http Response helper functions
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

//Namespace handlers GETS
func getStatus (w http.ResponseWriter, r *http.Request){
	resp := Response{}
	resp.Namespace = r.URL.Path
	resp.Message = "alive"

	returnOk(w,r,resp)
}

func getSystemStats (w http.ResponseWriter, r *http.Request){
	resp := sys.Sysinfo{}

	resp.Host = sys.GetHostStats()
	resp.Cpu = sys.GetCpuStats()
	resp.Mem = sys.GetMemStats()
	resp.Disk = sys.GetDiskStats()

	common.SendMetric(r.URL.Path, http.StatusOK, r.Method)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[ERROR] %s : %s\n", r.URL.Path, err)
	}
}

//Namespace handlers POSTS
func powerAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if err := common.CheckEnabled("power"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	cmdOut, err := pwr.LocalPowerAction(action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}

func updateAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if err := common.CheckEnabled("packages"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	cmdOut, err := apt.LocalUpdateAction(action)
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

	if err := common.CheckEnabled("packages"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	cmdOut, err := apt.LocalInstallAction(action, pkg)
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

	if err := common.CheckEnabled("service"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	cmdOut, err := svc.LocalServiceAction(service, action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}

func displayAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if err := common.CheckEnabled("display"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	cmdOut, err := hdmi.LocalPowerAction(action)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = cmdOut
	returnOk(w,r,resp)
}

func gpioSwitch (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)


	if err := common.CheckEnabled("gpio"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := gpio.ValidateGpioPin(pin); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := gpio.GpioToggle(pin); err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = "success"
	returnOk(w,r,resp)
}

func gpioPullDown (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if err := common.CheckEnabled("gpio"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := gpio.ValidateGpioPin(pin); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	pinState, err := gpio.GpioDown(pin)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	returnOk(w,r,resp)
}

func gpioPullUp (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["number"]

	resp := Response{}
	resp.Namespace = string(r.URL.Path)

	if err := common.CheckEnabled("gpio"); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	if err := gpio.ValidateGpioPin(pin); err != nil{
		returnBad(w,r,resp, err)
		return
	}

	pinState, err := gpio.GpioUp(pin)
	if err != nil{
		returnInternalError(w,r,resp, err)
		return
	}

	resp.Message = strconv.Itoa(int(pinState)) //Returns the pin readout
	returnOk(w,r,resp)
}