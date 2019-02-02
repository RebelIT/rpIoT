package common

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/alexcesaro/statsd.v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

const APPDIR = "/etc/api/"
const CONFIG = "api_config.json"

func SendMetric(uri string, responseCode int, method string ) error{
	c, err := ReadConfig()
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

func Cmd(cmdName string, args []string) (cmdOut string, err error) {
	out, err := exec.Command(cmdName, args...).Output()
	if err != nil {
		return string(out), err
	}

	return string(out),nil
}

func InRange(i, min, max int) error {
	if (i >= min) && (i <= max) {
		return nil
	} else {
		return errors.New("out of range")
	}
}

func StrToUint8(number string)(uint8, error){
	rawInt, err := strconv.Atoi(number) //convert the string number to int
	if err != nil{
		return 0, err
	}
	uInt := uint8(rawInt) //convert the int to uint8
	return uInt, nil
}

func CheckEnabled(function string)(err error){
	c, err := ReadConfig()
	if err != nil{
		return err
	}

	enabled := false
	found := false

	for _, f := range c.Functions{
		if f.Name == function{
			if f.Enabled{
				enabled = true
				found = true
			}
		}
	}
	if !found{
		return errors.New("function name "+function+" was not found in config file")
	}
	if !enabled{
		return errors.New("API function disabled")
	}

	return nil
}

func ReadConfig()(Config, error){
	c := Config{}
	config, err := ioutil.ReadFile(APPDIR+CONFIG)
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal(config, &c); err != nil{
		errorMsg := errors.New("unable to read Config file")
		return c, errorMsg
	}
	return c, nil
}