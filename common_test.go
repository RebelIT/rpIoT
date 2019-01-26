package main

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T){
	fmt.Println("***Testing local command execution***")
	testCmd := "ls"
	testArgs := []string{"$home"}

	err := command(testCmd, testArgs)
	if err != nil{
		t.Error(err)
	}
}

func TestValidateDisplayActionOn(t *testing.T){
	fmt.Println("***Testing validate display on conversion***")
	testState := []string{"on", "off"}

	for _, s := range testState{
		val, err := validateDisplayAction(s)
		if err != nil{
			t.Error(err)
		}
		if s == "on" {
			if val != "1" {
				t.Errorf("expected 1 got %s", val)
			}
		}else{
			if val != "0" {
				t.Errorf("expected 1 got %s", val)
			}
		}
	}
}

func TestValidateServiceActions(t *testing.T){
	fmt.Println("***Testing validate service actions***")
	testStates := []string{"start", "stop", "restart"}

	for _, state := range testStates{
		if err := validateServiceAction(state); err != nil{
			t.Error(err)
		}
	}
}

func TestCheckConfig(t *testing.T){
	fmt.Println("***Testing reading config file***")
	//check for a value I know exists in the default docker testing setup
	ok := "reboot"
	val, err := checkConfig(ok)
	if err != nil{
		t.Error(err)
	}
	if !val{
		t.Error("value should have been true")
	}

	//check for a value i knoe is not present in the default docker testing setup
	bad := "dafuk"
	val, err = checkConfig(bad)
	if err == nil{
		t.Errorf("This should have errored out because the value does not exist in the config %s, %v", bad, val)
	}
}

func TestSendingMetric(t *testing.T){
	fmt.Println("***Testing metric formatting***")
	if err := sendMetric("/bunk", 200, "POST"); err != nil{
		t.Error(err)
	}
}