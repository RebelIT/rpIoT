package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
)

func command(cmdName string, args []string) error {
	cmd := exec.Command(cmdName, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Start()

	if err != nil {
		return fmt.Errorf("failed starting %s %v: %s: %v", cmdName, args, stderr.String(), err)
	}

	return nil
}

func checkConfig(function string)(value bool, err error){
	c := Config{}
	config, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(config, &c); err != nil{
		errorMsg := errors.New("unable to read Config file")
		return false, errorMsg
	}

	for _, f := range c.Functions{
		if f.Name == function{
			return f.Enabled, nil
		}
	}
	return false, errors.New("function name "+function+" was not found in config file")
}

func validateDisplayAction(state string)(action string, err error){
	if state == "on"{
		action = "1"
	} else if state == "off"{
		action = "0"
	} else{
		return "", errors.New("monitor control action is invalid")
	}
	return action, nil
}

func validateServiceAction(action string) error{
	validated := false
	if action == "start"{
		validated = true
	}
	if action == "stop"{
		validated = true
	}
	if action == "restart" {
		validated = true
	}

	if !validated{
		return errors.New("service control "+action+" is invalid")
	}
	return nil
}