package factory

import (
	"testing"
)

// 所以工厂模式都会存在这样一种虚函数？函数本身看上去传进来的都是空的，但实际上使用的时候是不同的接口实现。
func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)

	return op.Result()
}

func TestSimpleFactory(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("TestSimpleFactory failed!")
	}
}

func TestStandardFactory(t *testing.T) {
	var factory OperatorFactory
	factory = CreatePlus{}
	if compute(factory, 1, 2) != 3 {
		t.Error("error with factory method pattern")
	}
	factory, ok := factory.(CreateMinus)
	// 这里我们可以看到其类型是CreatePlus所以不能进行转换。
	if !ok {
		factory = CreateMinus{}
	}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
