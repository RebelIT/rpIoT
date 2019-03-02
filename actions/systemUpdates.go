package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
	"io/ioutil"
	"strings"
)

func Update(action string)(cmdResp string, err error){
	if err := validateUpdateAction(action); err != nil{
		return "", err
	}

	go common.Cmd("apt-get", []string{action, "-y"})
	out := fmt.Sprintf("%s initiated", action)

	return out, nil
}

func Install(action string, pkg string)(cmdResp string, err error){
	if err := validateInstallAction(action); err != nil{
		return "", err
	}

	go common.Cmd("apt-get", []string{action ,pkg ,"-y"})
	out := fmt.Sprintf("%s initiated", action)

	return out, nil
}

func CheckUpdateLog()(log []string, error error){
	logData := []string{}

	data, err := ioutil.ReadFile(common.UPDATE_LOG_DIR)
	if err != nil{
		return nil, err
	}

	temp := strings.Split(string(data), "\n")

	for _, line := range temp{
		logData = append(logData, line)
	}

	return logData, nil
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
