package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/alexcesaro/statsd.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func SendSlack (message string){
	c, err := ReadConfig()
	if err != nil{
		log.Printf("[ERROR] slack alert : %s", err)
		return
	}

	if c.SlackWebhook == ""{
		//do nothing if not defined
		return
	}

	msgBody := SlackMsg{}
	msgBody.Username = GetHostname()
	msgBody.Text = message
	msgBody.IconPath = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSVhCe6Dx8fQfJrg6tprtsnnSVxAU6pGqRZMFkaVGy2UOr276lkGw"

	reqBody, err := json.Marshal(msgBody)

	resp, err := http.Post(c.SlackWebhook, "application/json", bytes.NewReader(reqBody))
	if err != nil{
		log.Printf("[ERROR] slack alert : %s", err)
	}

	if resp.StatusCode != 200 {
		log.Printf("[ERROR] slack alert: %s\n", fmt.Errorf("slack returned a non 200 response"))
		return
	}

	log.Printf("[INFO] slack alert sent: %s\n", message)
	return
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
		return fmt.Errorf("out of range")
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
		return fmt.Errorf("function name "+function+" was not found in config file")
	}
	if !enabled{
		return fmt.Errorf("API function disabled")
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
		errorMsg := fmt.Errorf("unable to read Config file")
		return c, errorMsg
	}
	return c, nil
}

func GetHostname()(hostname string){
	host, err := os.Hostname()
	if err != nil{
		host = "unknown"
	}

	return host
}