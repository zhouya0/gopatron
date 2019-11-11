package Singleton

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncOnce(t *testing.T) {
	o := sync.Once{}
	test := 1
	fmt.Printf("start value is %v\n", test)
	go do(&o, &test)
	go do(&o, &test)
	time.Sleep(time.Second * 2)
	fmt.Printf("final value is %v\n", test)
	if test != 2 {
		t.Error("value is not expected!")
	}
}
