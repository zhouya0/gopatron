package Singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	m := GetInstance()
	m.Set("Crisp")
	fmt.Printf("the name is %s", m.name)
	nm := GetInstance()
	fmt.Printf("agian the nam is %s", nm.name)
	if m.name != nm.name {
		t.Error("Singleton not quilified!")
	}
}
