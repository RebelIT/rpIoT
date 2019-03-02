package actions

import (
	"fmt"
	"github.com/rebelit/rpIoT/common"
	"github.com/stianeikeland/go-rpio"
	"strconv"
)

func GpioToggle(number string) error{
	pinNum, err := common.StrToUint8(number)
	if err != nil{
		return err
	}
	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNum)

	pin.Output()

	pin.Toggle()

	return nil
}

func GpioUp(number string) (pinState rpio.State, err error){
	pinNum, err := common.StrToUint8(number)
	if err != nil{
		return 0, err
	}
	if err := rpio.Open(); err != nil {
		return 0, err
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNum)

	pin.PullUp()
	state := pin.Read()

	return state, nil
}

func GpioDown(number string) (pinState rpio.State, err error){
	pinNum, err := common.StrToUint8(number)
	if err != nil{
		return 0, err
	}
	if err := rpio.Open(); err != nil {
		return 0, err
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNum)

	pin.PullDown()
	state := pin.Read()

	return state, nil
}

func ValidateGpioPin(pin string) error{
	//26 max pins to control
	i, err := strconv.Atoi(pin)
	if err != nil{
		return fmt.Errorf("%d is not a number", i)
	}

	if err := common.InRange(i, 2, 26); err != nil{
		return err
	}

	return nil
}
