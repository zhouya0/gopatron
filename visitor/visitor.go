package visitor

import (
	"fmt"
)

type IElement interface {
	Accept(visitor IVisitor) string
}

type Element struct {
}

func (e Element) Accept(visitor IVisitor) string {
	return visitor.visit()
}

type IVisitor interface {
	visit() string
}

type ProductionVisitor struct {
}

func (p ProductionVisitor) visit() string {
	fmt.Println("This is the ProductionVisitor!")
	return "ProductionVisitor"
}

type TestingVisitor struct {
}

func (t TestingVisitor) visit() string {
	fmt.Println("This is the testingVisitor")
	return "TestingVisitor"
}
