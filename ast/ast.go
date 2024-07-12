package ast

import (
	"strings"
)

type Node interface {
	node()
	String() string
}
type Expression interface {
	Node
	expressionNode()
}
type Statement interface {
	Node
	statementNode()
}

type Type = int
const (
	ILLEGAL Type = iota
	STRING
	INT
)

type Operator byte
const (
	OP_ADD = '+'
	OP_SUB = '-'
)
func (o Operator) String() string {
	return string(o)
}
func (o Operator) Node() {}

type Constant struct {
	Value   any
	Literal string
}

func (c Constant) Type() Type {
	switch c.Value.(type) {
	case string:
		return STRING
	case int:
		return INT
	}
	return ILLEGAL
}
func (c Constant) Node() {}
func (c Constant) String() string {
	return c.Literal
}

type BinaryOperation struct {
	left  Expression
	op    Operator
	right Expression
}

func (b *BinaryOperation) expressionNode() {}
func (b *BinaryOperation) String() string {
	var s strings.Builder

	s.WriteString("(")
	if b.left != nil { s.WriteString(b.left.String()) }
	s.WriteString(" ")
	if b.right != nil { s.WriteString(b.op.String()) }
	s.WriteString(" ")
	s.WriteString(b.right.String())
	s.WriteString(")")

	return s.String()
}

type UnaryOperation struct {
	op  Operator
	exp Expression
}

func (b *UnaryOperation) expressionNode() {}
func (b *UnaryOperation) Node() {}
func (b *UnaryOperation) String() string {
	var s strings.Builder

	s.WriteString("(")
	s.WriteString(b.op.String())
	if b.exp != nil {
		s.WriteString(b.exp.String())
	}
	s.WriteString(")")
	return s.String()
}

type ExprStatment struct {
	Value Expression
}

func (e *ExprStatment) Node() {}
func (e *ExprStatment) String() string {
	if e.Value == nil {
		return ""
	}
	return e.Value.String()
}
func (e *ExprStatment) statementNode() {}

func Expr(e Expression) *ExprStatment {
	return &ExprStatment{ e }
} 