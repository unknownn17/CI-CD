package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func Test_api(t *testing.T){
	want:="Hello"
	res,err:=http.Get("http://localhost:8080/health")
	if err!=nil{
		t.Error(err)
	}
	var have string
	if err:=json.NewDecoder(res.Body).Decode(&have);err!=nil{
		t.Error(err)
	}
	if want!=have{
		t.Error(err)
	}
}
