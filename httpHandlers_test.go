package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHttpGetStatus(t *testing.T){
	fmt.Println("Testing GET /api/alive")

	req, err := http.NewRequest("GET", "/api/alive", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getStatus)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Code != 200{
		t.Errorf("handler returned a non-200 response")
	}

}