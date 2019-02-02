package actions

import (
	"github.com/pkg/errors"
	"github.com/rebelit/rpIoT/common"
	"strings"
)

func Update(action string)(cmdResp string, err error){
	if err := validateUpdateAction(action); err != nil{
		return "", err
	}

	out, err := common.Cmd("apt-get", []string{action, "-y"})
	if err != nil{
		return "", err
	}

	if !strings.Contains(out, "Shutdown scheduled for"){
		return "", errors.New(out)
	}

	return out, nil
}

func Install(action string, pkg string)(cmdResp string, err error){
	if err := validateInstallAction(action); err != nil{
		return "", err
	}

	out, err := common.Cmd("apt-get", []string{action ,pkg ,"-y"})
	if err != nil{
		return "", err
	}

	return out, nil
}

func validateUpdateAction(action string) error{
	if action == "update"{
		return nil
	} else if action == "upgrade"{
		return nil
	} else{
		return errors.New("update action is invalid")
	}
	return nil
}

func validateInstallAction(action string) error{
	if action == "install"{
		return nil
	} else if action == "remove"{
		return nil
	} else{
		return errors.New("action "+action+" is invalid")
	}
	return nil
}
