package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	want := "Hello"
	res, err := http.Get("http://localhost:8080/health")
	if err != nil {
		log.Println(err)
	}
	var have string

	if err := json.NewDecoder(res.Body).Decode(&have); err != nil {
		log.Println(err)
	}
	if have != want {
		t.Errorf("want %s but found %s", want, have)
	}
}
