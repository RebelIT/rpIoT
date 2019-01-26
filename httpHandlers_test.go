package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpGetStatus(t *testing.T){
	fmt.Println("***Testing GET /api/alive***")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/alive", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")

	getStatus(w,r)

	if w.Result().StatusCode != 200{
		t.Errorf("handler returned non-200 response")
	}

	//have to trim whatespace in response body string - I don't know why this is happening yet...
	resp := strings.Replace(w.Body.String(), "\n", "", -1)
	expected := `{"namespace":"/api/alive","message":"alive"}`
	if resp != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestHttpGetSystem(t *testing.T){
	fmt.Println("***Testing GET /api/system***")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/system", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")

	getSystemStats(w,r)

	fmt.Printf("%+v", w.Body)
	if w.Result().StatusCode != 200{
		t.Errorf("handler returned non-200 response")
	}
}

//Expect a non-200 return code in the tests because the tests are expected to run in docker.
func TestHttpPostPower(t *testing.T){
	actions := []string{"reboot", "shutdown"}

	fmt.Printf("***Testing POST /api/power/{%s}***\n", actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/power/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		powerAction(w,r)

		if w.Result().StatusCode != 500{
			t.Errorf("handler returned non-500 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/power/`+a+`","message":"failed starting `+a+` [+1]: : exec: \"`+a+`\": executable file not found in $PATH"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

func TestHttpPostYum(t *testing.T){
	actions := []string{"update"}

	fmt.Printf("***Testing POST /api/apt/{%s}***\n",actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/apt/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		yumAction(w,r)

		if w.Result().StatusCode != 200{
			t.Errorf("handler returned non-200 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/apt/`+a+`","message":"success"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

func TestHttpPostService(t *testing.T){
	service := "redis"
	actions := []string{"stop", "start", "restart"}

	fmt.Printf("***Testing POST /api/service/{%s}/{%s}***\n",service, actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/service/"+service+"/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a, "service": service})

		serviceAction(w,r)

		if w.Result().StatusCode != 200{
			t.Errorf("handler returned non-200 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/service/`+service+`/`+a+`","message":"success"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

//Expect a non-200 return code in the tests because the tests are expected to run in docker.
func TestHttpPostDisplay(t *testing.T){
	actions := []string{"on", "off"}

	fmt.Printf("***Testing POST /api/display/{%s}***\n", actions)

	for _, a := range actions{
		i := ""
		if a == "on"{
			i = "1"
		} else {
			i = "0"
		}
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/display/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		displayAction(w,r)

		if w.Result().StatusCode != 500{
			t.Errorf("handler returned non-500 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/display/`+a+`","message":"failed starting vcgencmd [display_power `+i+`]: : exec: \"vcgencmd\": executable file not found in $PATH"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

//Expect a non-200 return code in the tests because the tests are expected to run in docker.
func TestHttpPostPinUp(t *testing.T){
	actions := []string{"1", "3"}

	fmt.Printf("***Testing POST /api/pullup/{%s}***\n", actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/pullup/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		gpioPullUp(w,r)

		if w.Result().StatusCode != 500{
			t.Errorf("handler returned non-500 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/pullup/`+a+`","message":"strconv.Atoi: parsing \"\": invalid syntax"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

//Expect a non-200 return code in the tests because the tests are expected to run in docker.
func TestHttpPostPinDown(t *testing.T){
	actions := []string{"1", "3"}

	fmt.Printf("***Testing POST /api/pulldown/{%s}***\n", actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/pulldown/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		gpioPullUp(w,r)

		if w.Result().StatusCode != 500{
			t.Errorf("handler returned non-500 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/pulldown/`+a+`","message":"strconv.Atoi: parsing \"\": invalid syntax"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}

//Expect a non-200 return code in the tests because the tests are expected to run in docker.
func TestHttpPostPinToggle(t *testing.T){
	actions := []string{"1", "3"}

	fmt.Printf("***Testing POST /api/toggle/{%s}***\n", actions)

	for _, a := range actions{
		w := httptest.NewRecorder()

		r := httptest.NewRequest("POST", "/api/toggle/"+a, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r = mux.SetURLVars(r, map[string]string{"action": a})

		gpioPullUp(w,r)

		if w.Result().StatusCode != 500{
			t.Errorf("handler returned non-500 response, are you testing in Docker???")
		}

		//have to trim whatespace in response body string - I don't know why this is happening yet...
		resp := strings.Replace(w.Body.String(), "\n", "", -1)
		expected := `{"namespace":"/api/toggle/`+a+`","message":"strconv.Atoi: parsing \"\": invalid syntax"}`
		if resp != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	}
}