package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
	"strings"
)

func HdmiPower(action string)(cmdResp string, err error){
	state, err := validateDisplayAction(action)
	if err != nil{
		return "", err
	}

	out, err := common.Cmd(string("vcgencmd"), []string{"display_power", state})
	if err != nil{
		return "", err
	}

	if !strings.Contains(out, "display_power="+state){
		return "", fmt.Errorf(out)
	}

	return out, nil
}

func GetHdmiPower()(state string, error error){
	out, err := common.Cmd(string("vcgencmd"), []string{"display_power"})
	if err != nil{
		return "", err
	}

	return strings.Split(out, "=")[1], nil
}

func validateDisplayAction(state string)(action string, err error){
	actual := ""
	if state == "on"{
		actual = "1"
	} else if state == "off"{
		actual = "0"
	} else{
		return "", fmt.Errorf("monitor control %s is invalid", action)
	}
	return actual, nil
}

