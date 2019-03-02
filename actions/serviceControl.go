package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
)

func ServiceControl(service string, action string)(cmdResp string, err error){
	if err := validateServiceAction(action); err != nil{
		return "", err
	}

	out, err := common.Cmd(string("service"), []string{service, action})
	if err != nil{
		return "", err
	}

	if out != ""{
		return "", fmt.Errorf(out)
	}

	return out, nil
}

func validateServiceAction(action string) error{
	validated := false
	if action == "start"{
		validated = true
	}
	if action == "stop"{
		validated = true
	}
	if action == "restart" {
		validated = true
	}

	if !validated{
		return fmt.Errorf("service control %s is invalid", action)
	}
	return nil
}