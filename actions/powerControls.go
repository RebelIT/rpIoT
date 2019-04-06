package actions

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

//not to self - syscall reboot magic numbers reference: https://golang.org/pkg/syscall/?GOOS=linux&GOARCH=mips64le

func SystemPower(action string)(err error){
	if err := validatePowerAction(action); err != nil{
		return err
	}

	switch action{
	case "reboot":
		go reboot()
		return nil

	case "shutdown":
		go shutdown()
		return nil

	default:
		return fmt.Errorf("power action is invalid")
	}
}

func validatePowerAction(action string) error{
	actions := []string{
		"reboot",
		"shutdown",
	}
	for _, a := range actions{
		if a == action{
			return nil
		}
	}
	return fmt.Errorf("power action is invalid")
}

func reboot(){
	//best effort no return in goroutine
	time.Sleep(time.Second * 10)

	var linuxRebootMagic1 uintptr = 0xfee1dead
	var linuxRebootMagic2 uintptr = 672274793
	var linuxRebootCmdRestart uintptr = 0x1234567 //LINUX_REBOOT_CMD_RESTART

	_,_, errno := syscall.Syscall(syscall.SYS_REBOOT, linuxRebootMagic1, linuxRebootMagic2, linuxRebootCmdRestart)
	if errno != 0{
		log.Printf("system reboot failed with error code %d", errno)
	}
}

func shutdown(){
	//best effort no return in goroutine
	time.Sleep(time.Second * 10)

	var linuxRebootMagic1 uintptr = 0xfee1dead
	var linuxRebootMagic2 uintptr = 672274793
	var linuxRebootCmdShutdown uintptr = 0x4321fedc //LINUX_REBOOT_CMD_POWER_OFF

	_,_, errno := syscall.Syscall(syscall.SYS_REBOOT, linuxRebootMagic1, linuxRebootMagic2, linuxRebootCmdShutdown)
	if errno != 0{
		log.Printf("system shutdown failed with error code %d", errno)
	}
}
