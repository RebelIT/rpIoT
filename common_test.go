package main

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T){
	fmt.Println("Testing local command execution")
	testCmd := "ls"
	testArgs := []string{"$home"}

	err := command(testCmd, testArgs)
	if err != nil{
		t.Error(err)
	}
}

func TestValidateDisplayActionOn(t *testing.T){
	fmt.Println("Testing validate display on conversion")
	testState := "on"

	val, err := validateDisplayAction(testState)
	if err != nil{
		t.Error(err)
	}

	if val != "1"{
		t.Errorf("expected 1 got %s", val)
	}
}

func TestValidateDisplayActionOff(t *testing.T){
	fmt.Println("Testing validate display off conversion")
	testState := "off"

	val, err := validateDisplayAction(testState)
	if err != nil{
		t.Error(err)
	}

	if val != "0"{
		t.Errorf("expected 0 got %s", val)
	}
}

func TestValidateServiceActions(t *testing.T){
	fmt.Println("Testing validate service actions")
	testStates := []string{"start", "stop", "restart"}

	for _, state := range testStates{
		if err := validateServiceAction(state); err != nil{
			t.Error(err)
		}
	}
}
