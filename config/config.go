package config

import (
	"encoding/json"
	"log"
	"os"
)

var ApiConfig *Config

func DoApiSettings(path string){
	ApiConfig = setConfiguration(path)
}

func setConfiguration(path string)(* Config){
	c := &Config{}
	loadDefaults(c)
	overrideDefaults(c, path)

	return c
}

func loadDefaults(c *Config){
	c.InstallDir = "/etc/api/"
	c.ApiPort = "8080"
	c.StatsdHost = ""
	c.SlackWebhook = ""
	c.DefaultUser = "admin"
	c.DefaultUserToken = "myPiApi"
	c.AuthEnabled = true

	c.SslEnabled = true
	c.SslCertName = "cert.cert"
	c.SslKeyName = "key.key"

	c.EndpointPower = true
	c.EndpointPackages = true
	c.EndpointService = true
	c.EndpointDisplay = true
	c.EndpointGpio = true
}

func overrideDefaults(c *Config, path string){
	if path == ""{
		log.Printf("[INFO], no configuration file specified")
		return
	}

	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		log.Printf("[ERROR], loading reading configuration file")
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&c)
}