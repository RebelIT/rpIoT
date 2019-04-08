package actions

import (
	"github.com/rebelit/rpIoT/common"
	"github.com/stianeikeland/go-rpio"
)

func GpioStatus(number int) (pinState Pin, error error){
	if err := rpio.Open(); err != nil {
		return Pin{}, err
	}
	defer rpio.Close()

	pin := rpio.Pin(uint8(number))

	state := Pin{}
	state.BcmPin = number
	state.State = pin.Read()

	return state, nil
}

func GpioAllStatus() (pinStates GpioStates, error error){
	pinMap := GpioStates{}

	if err := rpio.Open(); err != nil {
		return GpioStates{}, err
	}
	defer rpio.Close()

	for i := 2; i <= 26; i++ {
		p := Pin{}

		pin := rpio.Pin(uint8(i))
		p.BcmPin = i
		p.State = pin.Read()
		pinMap.Pins = append(pinMap.Pins, p)
	}

	return pinMap,nil
}

func GpioToggle(number int) (pinState Pin, error error){
	if err := rpio.Open(); err != nil {
		return Pin{},err
	}
	defer rpio.Close()

	pin := rpio.Pin(uint8(number))
	pin.Output()
	pin.Toggle()

	state := Pin{}
	state.BcmPin = number
	state.State = pin.Read()

	return state, nil
}

func GpioUp(number int) (pinState Pin, err error){
	if err := rpio.Open(); err != nil {
		return Pin{}, err
	}
	defer rpio.Close()

	pin := rpio.Pin(uint8(number))
	pin.PullUp()

	state := Pin{}
	state.BcmPin = number
	state.State = pin.Read()

	return state, nil
}

func GpioDown(number int) (pinState Pin, err error){
	if err := rpio.Open(); err != nil {
		return Pin{}, err
	}
	defer rpio.Close()

	pin := rpio.Pin(uint8(number))
	pin.PullDown()

	state := Pin{}
	state.BcmPin = number
	state.State = pin.Read()

	return state, nil
}

func ValidateGpioPin(pin int) error{
	//26 max pins to control
	if err := common.InRange(pin, 2, 29); err != nil{
		return err
	}

	return nil
}
