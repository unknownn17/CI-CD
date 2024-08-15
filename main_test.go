package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_api(t *testing.T) {
	want := "Hello"
	
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	GetHealth(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var have string
	if err := json.NewDecoder(rr.Body).Decode(&have); err != nil {
		t.Error(err)
	}
	if want != have {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}