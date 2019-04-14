package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rebelit/rpIoT/config"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	UPDATE_LOG_DIR = "/var/log/apt/history.log"
)

func SendMetric(uri string, responseCode int, method string ){
	measurement := "rpiot_web_api"
	if config.ApiConfig.StatsdHost == "" {
		//no statsd host configured, log metric and move along...move along...
		log.Printf("[INFO] metric : %s,uri_path=%s,response_code=%s,method=%s", measurement, uri,
			strconv.Itoa(responseCode),method)
		return
	}

	tags := statsd.Tags("uri_path", uri, "response_code", strconv.Itoa(responseCode), "method", method, "app", "rpiot")
	addrOpt := statsd.Address(config.ApiConfig.StatsdHost)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	client, err := statsd.New(addrOpt,fmtOpt,tags)
	if err != nil {
		log.Print(err)
	}
	defer client.Close()

	client.Increment(measurement)
	return
}

func SendSlack (message string){
	if config.ApiConfig.SlackWebhook == ""{
		//do nothing if not defined
		log.Printf("[INFO] slack : no webhook configured")
		return
	}

	msgBody := SlackMsg{}
	msgBody.Username = GetHostname()
	msgBody.Text = message
	msgBody.IconPath = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSVhCe6Dx8fQfJrg6tprtsnnSVxAU6pGqRZMFkaVGy2UOr276lkGw"

	reqBody, err := json.Marshal(msgBody)

	resp, err := http.Post(config.ApiConfig.SlackWebhook, "application/json", bytes.NewReader(reqBody))
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
	fmtOut := strings.Replace(string(out), "\n", "", -1)

	return fmtOut,nil
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

func StrToInt(str string)(number int){
	p, _ := strconv.Atoi(str)
	return p
}

func GetHostname()(hostname string){
	host, err := os.Hostname()
	if err != nil{
		host = "unknown"
	}

	return host
}