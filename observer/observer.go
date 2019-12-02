package observer

import (
	"fmt"
)

type Event struct {
	data string
}

type Subject interface {
	Attach(Obeserver)
	Detach(Obeserver)
	Notify(*Event)
}

type Obeserver interface {
	Update(*Event)
}

type ConcreteSubject struct {
	Observers map[Obeserver]string
}

func (c *ConcreteSubject) Attach(ob Obeserver) {
	c.Observers[ob] = ""
}

func (c *ConcreteSubject) Detach(ob Obeserver) {
	delete(c.Observers, ob)
}

func (c *ConcreteSubject) Notify(e *Event) {
	for obserber := range c.Observers {
		obserber.Update(e)
	}
}

type ConcreteObserver struct {
	Id int
}

func (c *ConcreteObserver) Update(e *Event) {
	fmt.Printf("The %d observer has update %s", c.Id, e.data)
}
