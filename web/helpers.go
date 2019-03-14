package web

import (
	"github.com/rebelit/rpIoT/config"
)

func validateAuth(username string,token string)(bool){
	if config.ApiConfig.DefaultUser == username && config.ApiConfig.DefaultUserToken == token{
		return true
	}

	return false
}