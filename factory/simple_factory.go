package factory

import (
	"fmt"
)

type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// 这里最关键的就是这个返回的类型也是API，所以就意味着这个接口的动态类型是不确定的。
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// 在这样的工厂模式下，我们不难发现，如果要增加一个新的类型，我们除了要写新的结构个方法之外，
// 还要改写工厂函数。这不符合工厂模式的（开放，封闭原则）。
