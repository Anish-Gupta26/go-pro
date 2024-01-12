package main

import "testing"


func TestInteger(t *testing.T){
	got := Add(2,3)
	want:= 7

	if got != want{
		t.Errorf("Error: Want:%d but got %d",want,got)
	}else {
		t.Log("Success")
	}
}