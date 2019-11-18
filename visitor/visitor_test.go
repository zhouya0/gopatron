package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	e := Element{}
	p := ProductionVisitor{}
	q := TestingVisitor{}
	es := e.Accept(p)
	qs := e.Accept(q)
	if es != "ProductionVisitor" || qs != "TestingVisitor" {
		t.Error("Not returning the right type!")
	}
}
