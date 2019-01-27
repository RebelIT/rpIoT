package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router

}

var routes = Routes{
	//API Endpoints
	Route{"status", "GET", "/api/alive", getStatus},
	Route{"stats", "GET", "/api/system", getSystemStats},
	Route{"powerControl", "POST", "/api/power/{action}", powerAction},
	Route{"packageManager", "POST", "/api/apt/{action}", updateAction},
	Route{"packageManager", "POST", "/api/apt/{action}/{package}", installAction},
	Route{"serviceControl", "POST", "/api/service/{service}/{action}", serviceAction},
	Route{"hdmiControl", "POST", "/api/display/{action}", displayAction},
	Route{"pinControl", "POST", "/api/pullup/{number}", gpioPullUp},
	Route{"pinControl", "POST", "/api/pulldown/{number}", gpioPullDown},
	Route{"pinControl", "POST", "/api/toggle/{number}", gpioSwitch},
}