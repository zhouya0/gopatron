package visitor

import (
	"fmt"
)

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (e Element) Accept(visitor IVisitor) {
	visitor.visit()
}

type IVisitor interface {
	visit()
}

type ProductionVisitor struct {
}

func (p ProductionVisitor) visit() {
	fmt.Println("This is the ProductionVisitor!")
}

type TestingVisitor struct {
}

func (t TestingVisitor) visit() {
	fmt.Println("This is the testingVisitor")
}
