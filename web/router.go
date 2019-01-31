package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	auth := Auth{}
	auth.Populate()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		router.Use(auth.AuthMiddleware)
	}

	return router
}

var routes = Routes{
	//API Endpoints
	Route{"status", "GET", "/api/alive", getStatus},
	Route{"stats", "GET", "/api/system", getSystemStats},
	Route{"powerControl", "POST", "/api/power/{action}", powerAction},
	Route{"packageManager", "POST", "/api/apt/{action}", updateAction},
	Route{"packageManager", "POST", "/api/apt/{package}/{action}", installAction},
	Route{"serviceControl", "POST", "/api/service/{service}/{action}", serviceAction},
	Route{"hdmiControl", "POST", "/api/display/{action}", displayAction},
	Route{"pinControl", "POST", "/api/gpio/{number}/pullup", gpioPullUp},
	Route{"pinControl", "POST", "/api/gpio/{number}/pulldown", gpioPullDown},
	Route{"pinControl", "POST", "/api/gpio/{number}/toggle", gpioSwitch},
}

//************
//Playing with gorilla/mux auth middleware examples.
//plain text for playing with Auth.
//TODO: store this securely somewhere.
func (a *Auth) Populate() {
	a.TokenUsers = make(map[string]string)
	a.TokenUsers["12345"] = "test"
}

func (a *Auth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/alive"{
			//skip auto for health check
			next.ServeHTTP(w, r)
		}else {
			token := r.Header.Get("X-Session-Token")

			if user, found := a.TokenUsers[token]; found {
				// We found the token in our map
				log.Printf("Authenticated user %s\n", user)
				// Pass down the request to the next middleware (or final handler)
				next.ServeHTTP(w, r)
			} else {
				// Write an error and stop the handler chain
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		}
	})
}