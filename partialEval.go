package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jun-hf/essential_compilation/ast"
)

func PartialEval(p *ast.Program) (*ast.Program, []error) {
	if !IsLanguageINT(p) {
		return nil, []error{errors.New("program is not a Language")}
	}
	stmts := make([]ast.Statement, len(p.Body))
	errorList := []error{}
	for i, stmt := range p.Body {
		st, err := partialEvalStatement(stmt)
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		stmts[i] = st
	}
	return &ast.Program{Body: stmts}, errorList
}

func partialEvalStatement(stmt ast.Statement) (ast.Statement, error) {
	if stmt == nil {
		return nil, fmt.Errorf("nil passed into partialEvalStatement")
	}
	switch s := stmt.(type) {
	case *ast.ExprStatment:
		e, err := partialEvalExpression(s.Value)
		if err != nil {
			return nil, err
		}
		return &ast.ExprStatment{Value: e}, nil
	}
	return nil, fmt.Errorf("%s is not a valid statement", stmt)
}

func partialEvalExpression(exp ast.Expression) (ast.Expression, error) {
	if exp == nil {
		return nil, errors.New("nil passsed to partialEvalExpression")
	}
	switch e := exp.(type) {
	case *ast.Constant:
		return exp, nil
	case *ast.BinaryOperation:
		return partialEvalBinaryOperation(e)
	case *ast.UnaryOperation:
		return partialEvalUnaryOperation(e)
	}

	return exp, errors.New("exppression not the language")
}

func partialEvalUnaryOperation(e *ast.UnaryOperation) (ast.Expression, error) {
	if e.Op != ast.OP_SUB {
		return nil, errors.New("unsupportted unary operation")
	}

	eval, err := partialEvalExpression(e.Exp)
	if err != nil {
		return nil, err
	}

	if c, ok := eval.(*ast.Constant); ok {
		val, ok := c.Value.(int)
		if !ok {
			return nil, errors.New("constant is not integer")
		}
		return &ast.Constant{Value: -val, Literal: strconv.Itoa(-val)}, nil
	}

	return &ast.UnaryOperation{Op: e.Op, Exp: eval}, nil
}

func partialEvalBinaryOperation(b *ast.BinaryOperation) (ast.Expression, error) {
	lEval, err := partialEvalExpression(b.Left)
	if err != nil {
		return nil, err
	}
	rEval, err := partialEvalExpression(b.Right)
	if err != nil {
		return nil, err
	}

	lConstant, lIsConstant := lEval.(*ast.Constant)
	rConstant, rIsConstant := rEval.(*ast.Constant)
	switch {
	case lIsConstant && rIsConstant:
		return evalBinaryOperation(b.Op, lConstant, rConstant)
	case !lIsConstant && rIsConstant:
		return &ast.BinaryOperation{Left: lEval, Op: b.Op, Right: rConstant}, nil
	case lIsConstant && !rIsConstant:
		return &ast.BinaryOperation{Left: lConstant, Op: b.Op, Right: rEval}, nil
	}
	return &ast.BinaryOperation{Left: lEval, Op: b.Op, Right: rEval}, nil
}

func evalBinaryOperation(op ast.Operator, left, right *ast.Constant) (*ast.Constant, error) {
	if left.Type() != right.Type() {
		return nil, errors.New("type mismatch")
	}

	switch left.Type() {
	case ast.STRING:
		if op.String() != "+" { 
			return nil, errors.New("invalid binary operation on string type")
		}
		var s strings.Builder
		s.WriteString(left.String())
		s.WriteString(right.String())
		return &ast.Constant{Value: s.String(), Literal: s.String()}, nil
	case ast.INT:
		lInt, lOk := left.Value.(int)
		rInt, rOk := right.Value.(int)
		if !lOk || !rOk {
			return nil, errors.New("ast.Constant: inconsistent type and Value")
		}
		var total int
		switch op.String() {
		case "+":
			total = lInt + rInt
		case "-":
			total = lInt - rInt
		}
		return &ast.Constant{Value: total, Literal: strconv.Itoa(total)}, nil
	}
	return nil, errors.New("unsupported constant type")
}