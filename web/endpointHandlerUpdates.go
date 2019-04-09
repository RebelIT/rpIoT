package web

import (
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/actions"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)

func getUpdates (w http.ResponseWriter, r *http.Request){	//Returns everything in apt/history.log
	logData, err  := actions.CheckUpdateLog()
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
	}

	returnOkAptLog(w,r,logData)
	return
}

func updateAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]

	if !config.ApiConfig.EndpointPower{
		returnErr(w,r,http.StatusNotFound)
		return
	}

	status, err := actions.Update(action)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkAptInstallUpdate(w,r,status)
	return
}

func installAction (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	pkg := vars["package"]

	if !config.ApiConfig.EndpointPackages{
		returnErr(w,r,http.StatusNotFound)
		return
	}

	status, err := actions.Install(action, pkg)
	if err != nil{
		returnErr(w,r,http.StatusInternalServerError)
		return
	}

	returnOkAptInstallUpdate(w,r,status)
	return
}