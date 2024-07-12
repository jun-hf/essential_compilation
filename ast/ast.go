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
func (c Constant) node() {}
func (c Constant) String() string {
	return c.Literal
}

type BinaryOperation struct {
	Left  Expression
	Op    Operator
	Right Expression
}

func (b *BinaryOperation) Operator() string {
	return b.Op.String()
}
func (b *BinaryOperation) node() {}
func (b *BinaryOperation) expressionNode() {}
func (b *BinaryOperation) String() string {
	var s strings.Builder

	s.WriteString("(")
	if b.Left != nil { s.WriteString(b.Left.String()) }
	s.WriteString(" ")
	s.WriteString(b.Op.String()) 
	s.WriteString(" ")
	if b.Right != nil { s.WriteString(b.Right.String()) }
	s.WriteString(")")

	return s.String()
}

type UnaryOperation struct {
	Op  Operator
	Exp Expression
}

func (b *UnaryOperation) Operator() string { return b.Op.String() }
func (b *UnaryOperation) expressionNode() {}
func (b *UnaryOperation) node() {}
func (b *UnaryOperation) String() string {
	var s strings.Builder

	s.WriteString("(")
	s.WriteString(b.Op.String())
	if b.Exp != nil {
		s.WriteString(b.Exp.String())
	}
	s.WriteString(")")
	return s.String()
}

type ExprStatment struct {
	Value Expression
}

func (e *ExprStatment) node() {}
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


type Porgram struct {
	Body []Statement
}
