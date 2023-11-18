package visitor

import (
	"fmt"
	"strings"
)

/**
Need to define a new operation on an entire type hierarchy
- Given a document model(lists,paragraphs,dynamo_util.), we want to add printing functionality
- Do not want to keep modifying every type in the hierarchy
- Want to have the new functionality separate (SRP)
This approach is often used for traversal
- Alternative to Iterator
- Hierarchy members help you traverse themselves

A pattern where a component(visitor) is allowed to traverse the entire hierarchy of types.
Implemented by propagating a single Accept() method throughout the entire hierarchy.
*/

// (1+2)+3
type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}
type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{
		sb: strings.Builder{},
	}
}

func (ex *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ex.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ex *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ex.sb.WriteRune('(')
	e.left.Accept(ex)
	ex.sb.WriteRune('+')
	e.right.Accept(ex)
	ex.sb.WriteRune(')')
}

func (ex *ExpressionPrinter) String() string {
	return ex.sb.String()
}

var _ ExpressionVisitor = (*ExpressionPrinter)(nil)

/**
Dispatch
Which function to call ?
Single dispatch: depends on name of request and type of receiver
Double dispatch: depends on name of request and type of two receivers(type of visitor, type of element being visited)
*/

func Start() {
	// 1 + ( 2 + 3 )
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			&DoubleExpression{2},
			&DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep.String())
}
