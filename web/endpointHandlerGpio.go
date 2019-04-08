package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/common"
	"github.com/rebelit/rpIoT/config"
	"net/http"
	"time"
)

func gpioToggle (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := common.StrToInt(vars["number"])

	if !config.ApiConfig.EndpointGpio{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	if err := actions.ValidateGpioPin(pin); err != nil{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	state, err := actions.GpioToggle(pin)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkGpioPinState(w,r,state)
	return
}

func gpioPullDown (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := common.StrToInt(vars["number"])


	if !config.ApiConfig.EndpointGpio{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	if err := actions.ValidateGpioPin(pin); err != nil{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	state, err := actions.GpioDown(pin)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkGpioPinState(w,r,state)
	return
}

func gpioPullUp (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := common.StrToInt(vars["number"])


	if !config.ApiConfig.EndpointGpio{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	if err := actions.ValidateGpioPin(pin); err != nil{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	state, err := actions.GpioUp(pin)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkGpioPinState(w,r,state)
	return
}

func gpioDepress (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := common.StrToInt(vars["number"])
	ms := common.StrToInt(vars["millisecond"])
	toggleTimes := 2

	if !config.ApiConfig.EndpointGpio{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	if err := actions.ValidateGpioPin(pin); err != nil{
		returnErr(w,r,http.StatusBadRequest)
		return
	}

	for i := 1;  i<=toggleTimes; i++ {
		_, err := actions.GpioToggle(pin)
		if err != nil{
			returnErr(w,r,http.StatusInternalServerError)
			return
		}

		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

	state, err := actions.GpioStatus(pin)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkGpioPinState(w,r,state)
	return
}

func getAllGpioStates (w http.ResponseWriter, r *http.Request) {
	states, err := actions.GpioAllStatus()
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkAllGpioPinStates(w,r,states)
	return
}

func getGpioState (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := common.StrToInt(vars["number"])

	state, err := actions.GpioStatus(pin)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkGpioPinState(w,r,state)
	return
}