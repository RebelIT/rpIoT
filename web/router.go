package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebelit/rpIoT/config"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		if config.ApiConfig.AuthEnabled{	//auth enabled use auth middleware function
			router.Use(AuthMiddleware)
		}
	}

	return router
}

var routes = Routes{
	//API Endpoints
	Route{"status", "GET", "/api/alive", getStatus},
	Route{"stats", "GET", "/api/system", getSystemStats},
	Route{"powerControl", "POST", "/api/power/{action}", powerAction},
	Route{"packageManager", "GET", "/api/apt", getUpdates},
	Route{"packageManager", "POST", "/api/apt/{action}", updateAction},
	Route{"packageManager", "POST", "/api/apt/{package}/{action}", installAction},
	Route{"serviceControl", "POST", "/api/service/{service}/{action}", serviceAction},
	Route{"hdmiControl", "GET", "/api/display", displayGet},
	Route{"hdmiControl", "POST", "/api/display/{action}", displayAction},
	Route{"pinControl", "POST", "/api/gpio/{number}/pullup", gpioPullUp},
	Route{"pinControl", "POST", "/api/gpio/{number}/pulldown", gpioPullDown},
	Route{"pinControl", "POST", "/api/gpio/{number}/toggle", gpioSwitch},
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/alive"{
			//skip auth for health check
			next.ServeHTTP(w, r)
		}else {
			hdrToken := r.Header.Get("X-API-Token")
			hdrUser := r.Header.Get("X-API-User")
			if validateAuth(hdrUser,hdrToken){
				// Pass down the request to the next handler
				next.ServeHTTP(w, r)
			} else{
				resp := Response{}
				resp.Namespace = string(r.URL.Path)
				returnUnauthorized(w, r, resp, fmt.Errorf("nope... unauthorized :("))
			}
		}
	})
}