package strategy

import (
	"testing"
)

func TestStragety(t *testing.T) {
	s := NewStrategy("m")
	c := Computer{
		Num1: 3,
		Num2: 4,
	}
	c.SetStrategy(s)
	r := c.Do()
	expect := 12
	if r != expect {
		t.Errorf("Unexpected value: %d foun ", r)
	}
}
