package main

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpGetStatus(t *testing.T){
	fmt.Println("Testing GET /api/alive")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/alive", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")

	getStatus(w,r)

	//have to trim whatespace - I don't know why this is happening yet...
	resp := strings.Replace(w.Body.String(), "\n", "", -1)
	expected := `{"namespace":"/api/alive","message":"alive"}`
	if resp != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}