package strategy

import (
	"fmt"
)

type Strategier interface {
	Compute(num1, num2 int) int
}

type Division struct{}

func (p Division) Compute(num1, num2 int) int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
			return
		}
	}()

	if num2 == 0 {
		panic("nums2 must not be 0!")
	}

	return num1 / num2
}

type Subtraction struct{}

func (s Subtraction) Compute(num1, num2 int) int {
	return num1 - num2
}

type Addition struct{}

func (a Addition) Compute(num1, num2 int) int {
	return num1 + num2
}

type Multiplication struct{}

func (m Multiplication) Compute(num1, num2 int) int {
	return num1 * num2
}

func NewStrategy(t string) (res Strategier) {
	switch t {
	case "s":
		res = Subtraction{}
	case "m":
		res = Multiplication{}
	case "d":
		res = Division{}
	case "a":
		fallthrough
	default:
		res = Addition{}
	}
	return res
}

type Computer struct {
	Num1, Num2 int
	strate     Strategier
}

func (c *Computer) SetStrategy(strate Strategier) {
	c.strate = strate
}

func (c Computer) Do() int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
		}
	}()

	if c.strate == nil {
		panic("Strategier is null")
	}

	return c.strate.Compute(c.Num1, c.Num2)
}
