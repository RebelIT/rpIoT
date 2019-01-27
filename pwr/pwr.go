package pwr

import (
	"github.com/pkg/errors"
	"github.com/rebelit/rpIoT/common"
	"strings"
)

func LocalPowerAction(action string)(cmdResp string, err error){
	cmd, err := validatePowerAction(action)
	if err != nil{
		return "", err
	}

	out, err := common.Cmd("shutdown", []string{cmd,"+1"})
	if err != nil{
		return "", err
	}

	if !strings.Contains(out, "Shutdown scheduled for"){
		return "", errors.New(out)
	}

	return out, nil
}

func validatePowerAction(action string)(val string, err error){
	actual := ""
	if action == "shutdown"{
		actual = " -s"
	} else if action == "reboot"{
		actual = "-r"
	} else{
		return "", errors.New("power action is invalid")
	}
	return actual, nil
}
