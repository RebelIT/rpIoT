package web

import "github.com/rebelit/rpIoT/common"

func validateAuth(username string,token string)(bool){
	c, err := common.ReadConfig()
	if err != nil {
		return false
	}
	if c.DefaultUser == username && c.DefaultUserToken == token{
		return true
	}

	return false
}