package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetResponse(t *testing.T) {
	val := reflect.ValueOf(GetResponse())
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Println(val.Type().Field(i).Tag.Get("json"))
	}

	if true {
		t.Error("Something went wrong")
	}
}
