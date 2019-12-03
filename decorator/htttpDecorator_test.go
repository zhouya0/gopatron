package Decorator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHttpDecorator(t *testing.T) {
	testFunc := WithServerHeader(hello)
	typeFunc := reflect.TypeOf(testFunc)
	typeString := fmt.Sprint(typeFunc)
	if typeString != "http.HandlerFunc" {
		t.Error("Error happens!")
	}
}
