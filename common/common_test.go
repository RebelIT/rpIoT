package common

import (
	"fmt"
	"testing"
)

func TestMetricFormat(t *testing.T){
	fmt.Println("*** Testing metric formatting")
	if err := SendMetric("/bunk", 200, "POST"); err != nil{
		t.Error(err)
	}
}

func TestCommand(t *testing.T){
	fmt.Println("*** Testing local command execution")
	testCmd := "ls"
	testArgs := []string{"$home"}

	resp, err := Cmd(testCmd, testArgs)
	if err != nil{
		t.Error(err)
	}

	expected := "bin  src"
	if resp != expected {
		t.Errorf("cmd response unexpected: got %v want %v",
			resp, expected)
	}
}

func TestInRange(t *testing.T){
	fmt.Println("*** Testing int in range")
	int := 3
	if err := InRange(int, 2, 26); err != nil{
		t.Errorf("inRange unexpected: got %v want nil", err)
	}
}

func TestOutOfRange(t *testing.T){
	fmt.Println("*** Testing int out of range")
	int := 30
	if err := InRange(int, 2, 26); err != nil{
		t.Errorf("inRange unexpected: got %v want nil", err)
	}
}

func TestStringToIntOk(t *testing.T){
	fmt.Println("*** Testing string to int conversion")
	str := "3"

	resp, err := StrToUint8(str);
	if err != nil{
		t.Errorf("conversion failed %v",err)
	}

	expected := uint8(3)
	if resp != expected {
		t.Errorf("conversion unexpected: got %v want %v",
			resp, expected)
	}
}

func TestStringToIntBad(t *testing.T){
	fmt.Println("*** Testing string to int conversion")
	str := "z"

	resp, err := StrToUint8(str);
	if err != nil{
		t.Errorf("conversion failed %v",err)
	}

	expected := uint8(3)
	if resp != expected {
		t.Errorf("conversion unexpected: got %v want %v",
			resp, expected)
	}
}

func TestCheckEnabled(t *testing.T){
	fmt.Println("***Testing reading config file")
	//check for a value I know exists in the default docker testing setup
	ok := "reboot"
	if err := CheckEnabled(ok); err != nil{
		t.Error(err)
	}

	//check for a value i knoe is not present in the default docker testing setup
	bad := "dafuk"
	if err := CheckEnabled(bad); err != nil{
		t.Errorf("This should have errored out because the value does not exist in the config: value %s", bad)
	}
}