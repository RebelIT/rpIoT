package actions

import (
	"github.com/pkg/errors"
	"github.com/rebelit/rpIoT/common"
	"strings"
)

func LocalPowerAction(action string)(cmdResp string, err error){
	state, err := validateDisplayAction(action)
	if err != nil{
		return "", err
	}

	out, err := common.Cmd(string("vcgencmd"), []string{"display_power", state})
	if err != nil{
		return "", err
	}

	if !strings.Contains(out, "display_power="+state){
		return "", errors.New(out)
	}

	return out, nil
}


func validateDisplayAction(state string)(action string, err error){
	actual := ""
	if state == "on"{
		actual = "1"
	} else if state == "off"{
		actual = "0"
	} else{
		return "", errors.New("monitor control "+action+" is invalid")
	}
	return actual, nil
}

