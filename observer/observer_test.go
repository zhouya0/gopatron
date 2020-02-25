package observer

import (
	"testing"
)

func InputObeserver(b Obeserver) {
	b.Update(nil)
}

func testObserver(t testing.T) {
	c1 := ConcreteObserver{1}
	c2 := ConcreteObserver{2}
	var subject ConcreteSubject
	// 谁实现了Update方法？是ConcreteObserver还是&ConcreteObserver？
	// subject.Attach(c1) cannot use c1 (variable of type ConcreteObserver) as Obeserver value in argument to subject.Attach: missing method Update
	// 接收体定义了到底是指针对象还是非指针对象来调用方法
	subject.Attach(&c1)
	subject.Attach(&c2)
	event := Event{"this is a event!"}
	// 这里为什么成功？是不是它做了一个自动的转化？
	c1.Update(&event)
	(&c1).Update(&event)
	subject.Notify(&event)
}
