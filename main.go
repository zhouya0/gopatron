package main

import (
	"flag"
	"fmt"
	"gopattern/strategy"
)

var stra *string = flag.String("type", "a", "input the stategy")
var num1 *int = flag.Int("num1", 1, "input num1")
var num2 *int = flag.Int("num2", 1, "input num2")

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("Hello world")
	com := strategy.Computer{Num1: *num1, Num2: *num2}
	strate := strategy.NewStrategy(*stra)

	com.SetStrategy(strate)
	fmt.Println(com.Do())
}
