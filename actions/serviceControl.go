package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
)

func GetService(service string)(state Service, error error){
	svcState := Service{}
	out, err := common.Cmd(string("systemctl"), []string{"check", service})
	if err != nil{
		if err.Error() == "exit status 3"{
			svcState.Status = "inactive or not installed"
			return svcState, nil //workaround funky systemctl return when not in a bash shell
		} else{
			return Service{}, err
		}
	}
	svcState.Status = out
	return svcState, nil
}

func ServiceControl(service string, action string)(state Service, err error){
	if err := validateServiceAction(action); err != nil{
		return Service{}, err
	}

	svcState := Service{}
	out, err := common.Cmd(string("service"), []string{service, action})
	if err != nil{
		return Service{}, err
	}

	if out != ""{
		svcState.Status = fmt.Sprintf("unable to %s %s check system log", action, service)
		return Service{}, fmt.Errorf(out)
	}

	svcState.Status = "success"
	return svcState, nil
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