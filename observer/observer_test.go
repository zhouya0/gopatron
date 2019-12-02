package observer

import (
	"testing"
)

func testObserver(t testing.T) {
	c1 := &ConcreteObserver{1}
	c2 := &ConcreteObserver{2}
	// 注意这里必须c1,c2必须是指针，不然就会出现“has pointer receiver的错”
	// 其实道理就是，我的实现是指针，我传只能传指针；我的实现如果是实体，我传，可以传实体，可以传指针。
	// 指针能找到实体，但是实体找不到指针！！！
	var subject ConcreteSubject
	subject.Attach(c1)
	subject.Attach(c2)
	event := Event{"this is a event!"}
	subject.Notify(&event)
}
