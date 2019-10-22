package factory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("TestSimpleFactory failed!")
	}
}
