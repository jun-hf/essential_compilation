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
	if e.Op != ast.Operator(ast.STRING) {
		return nil, errors.New("unsupportted unary operation")
	}
	u := &ast.UnaryOperation{Op: e.Op}
	var evalConstant *ast.Constant
	c, ok := e.Exp.(*ast.Constant)
	if !ok {
		eval, err := partialEvalExpression(e.Exp)
		if err != nil {
			return nil, err
		}
		if evalC, ok := eval.(*ast.Constant); ok {
			evalConstant = evalC
		}
		u.Exp = eval
	} else {
		evalConstant = c
	}
	if evalConstant == nil {
		return u, nil
	}
	val, ok := evalConstant.Value.(int)
	if !ok {
		return nil, errors.New("constant is not integer")
	}
	return &ast.Constant{Value: -val, Literal: strconv.Itoa(-val)}, nil
} 

func partialEvalBinaryOperation(b *ast.BinaryOperation) (ast.Expression, error) {
	binOperation := &ast.BinaryOperation{Op: b.Op}
	var leftConstant *ast.Constant
	var rightConstant *ast.Constant
	switch l := b.Left.(type) {
	case *ast.Constant:
		leftConstant = l
	default:
		lexp, err := partialEvalExpression(l)
		if err != nil {
			return nil, err
		}
		if lc, ok := lexp.(*ast.Constant); ok {
			leftConstant = lc
		}
		binOperation.Left = lexp
	}
	switch r := b.Right.(type) {
	case *ast.Constant:
		rightConstant = r
	default:
		rexp, err := partialEvalExpression(r)
		if err != nil {
			return nil, err
		}
		if rc, ok := rexp.(*ast.Constant); ok {
			rightConstant = rc
		}
		binOperation.Right = rexp
	}

	if leftConstant != nil && rightConstant != nil {
		if leftConstant.Type() != rightConstant.Type() {
			return nil, errors.New("type mismatch")
		}
		if leftConstant.Type() == ast.STRING && rightConstant.Type() == ast.STRING {
			if binOperation.Op.String() != "+" {
				return nil, errors.New("invalid binary operation")
			}
			var s strings.Builder
			s.WriteString(leftConstant.String())
			s.WriteString(rightConstant.String())
			return &ast.Constant{Value: s.String(), Literal: s.String()}, nil
		}
		if leftConstant.Type() == ast.INT && rightConstant.Type() == ast.INT {
			lInt, _ := leftConstant.Value.(int)
			rInt, _ := rightConstant.Value.(int)

			total := 0
			switch binOperation.Op.String() {
			case "+":
				total = lInt + rInt
			case "-":
				total = lInt - rInt
			}
			return &ast.Constant{Value: total, Literal: strconv.Itoa(total)}, nil
		}
	}
	return binOperation, nil
}