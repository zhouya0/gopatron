// +build !linux

package Singleton

import (
	"fmt"
	"sync"
)

func test() {}

// sync.Once可以保证穿进去的函数可以直被执行一次
func do(o *sync.Once, value *int) {
	fmt.Println("Start do")
	o.Do(func() {
		fmt.Println("Doing something once")
		*value = *value + 1
		fmt.Printf("value is %v\n", *value)
	})
	fmt.Println("Do end")
}
