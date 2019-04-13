package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
	"io/ioutil"
	"strings"
)

func Update(action string)(status AptInstallUpdate, err error){
	if err := validateUpdateAction(action); err != nil{
		return AptInstallUpdate{}, err
	}

	go common.Cmd("apt-get", []string{action, "-y"})
	out := AptInstallUpdate{}
	out.Status = fmt.Sprintf("%s async action in progress, GET - /api/apt for status", action)

	return out, nil
}

func Install(action string, pkg string)(status AptInstallUpdate, err error){
	if err := validateInstallAction(action); err != nil{
		return AptInstallUpdate{}, err
	}

	go common.Cmd("apt-get", []string{action ,pkg ,"-y"})
	out := AptInstallUpdate{}
	out.Status = fmt.Sprintf("%s %s async action in progress, GET - /api/apt for status", action, pkg)

	return out, nil
}

func CheckUpdateLog()(log AptLog, error error){
	aptLog := AptLog{}
	aptLogdata := AptLog{}.Log

	data, err := ioutil.ReadFile(common.UPDATE_LOG_DIR)
	if err != nil{
		return AptLog{}, err
	}

	temp := strings.Split(string(data), "\n")

	for _, line := range temp{
		aptLogdata = append(aptLogdata, line)
	}

	aptLog.Log = aptLogdata

	return aptLog, nil
}

func validateUpdateAction(action string) error{
	if action == "update"{
		return nil
	} else if action == "upgrade"{
		return nil
	} else{
		return fmt.Errorf("update action is invalid")
	}
	return nil
}

func validateInstallAction(action string) error{
	if action == "install"{
		return nil
	} else if action == "remove"{
		return nil
	} else{
		return fmt.Errorf("action %s is invalid", action)
	}
	return nil
}
