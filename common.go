package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stianeikeland/go-rpio"
	"gopkg.in/alexcesaro/statsd.v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
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
	c, err := readConfig()
	if err != nil{
		return false, err
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

func readConfig()(Config, error){
	c := Config{}
	config, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal(config, &c); err != nil{
		errorMsg := errors.New("unable to read Config file")
		return c, errorMsg
	}
	return c, nil
}

func sendMetric(uri string, responseCode int, method string ) error{
	c, err := readConfig()
	if err != nil{
		log.Printf("[ERROR] metric : %s", err)
		return err
	}
	measurement := "rpiot_web_api"
	if c.Statsd_host == "" {
		//no statsd host configured, log metric and move along...move along...
		log.Printf("[INFO] metric : %s,uri_path=%s,response_code=%s,method=%s", measurement, uri,
			strconv.Itoa(responseCode),method)
		return err
	}

	tags := statsd.Tags("uri_path", uri, "response_code", strconv.Itoa(responseCode), "method", method)
	addrOpt := statsd.Address(c.Statsd_host)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	client, err := statsd.New(addrOpt,fmtOpt,tags)
	if err != nil {
		log.Print(err)
	}
	defer client.Close()

	client.Increment(measurement)
	return err
}

func strToUint8(number string)(uint8, error){
	rawInt, err := strconv.Atoi(number) //convert the string number to int
	if err != nil{
		return 0, err
	}
	uInt := uint8(rawInt) //convert the int to uint8
	return uInt, nil
}

func gpioToggle(number string) error{
	pinNum, err := strToUint8(number)
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

func gpioUp(number string) (pinState rpio.State, err error){
	pinNum, err := strToUint8(number)
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

func gpioDown(number string) (pinState rpio.State, err error){
	pinNum, err := strToUint8(number)
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