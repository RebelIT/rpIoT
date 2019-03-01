package actions

import (
	"github.com/pkg/errors"
	"log"
	"syscall"
	"time"
)

func SystemPower(action string)(err error){
	if err := validatePowerAction(action); err != nil{
		return err
	}

	go reboot()

	return nil
}

func validatePowerAction(action string) error{
	if action != "reboot"{
		return errors.New("power action is invalid")
	}
	return nil
}

func reboot(){
	//best effort no return in goroutine
	time.Sleep(time.Second * 10)

	var linuxRebootMagic1 uintptr = 0xfee1dead
	var linuxRebootMagic2 uintptr = 672274793
	var linuxRebootCmdRestart uintptr = 0x1234567

	_,_, errno := syscall.Syscall(syscall.SYS_REBOOT, linuxRebootMagic1, linuxRebootMagic2, linuxRebootCmdRestart)
	if errno != 0{
		log.Printf("system reboot failed with error code %d", errno)
	}
}
