package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
	"strings"
)

func HdmiToggleState(action string)(displayState Display, err error){
	suggestedState, err := validateDisplayAction(action)
	if err != nil{
		return Display{}, err
	}

	out, err := common.Cmd(string("vcgencmd"), []string{"display_power", suggestedState})
	if err != nil{
		return Display{}, err
	}

	if !strings.Contains(out, "display_power="+suggestedState){
		return Display{}, fmt.Errorf(out)
	}

	state := Display{}
	state.HdmiState = common.StrToInt(strings.Split(out,"=")[1])

	return state, nil
}

func GetHdmiState()(displayState Display, error error){
	out, err := common.Cmd(string("vcgencmd"), []string{"display_power"})
	if err != nil{
		return Display{}, err
	}

	state := Display{}
	state.HdmiState = common.StrToInt(strings.Split(out,"=")[1])

	return state, nil
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

