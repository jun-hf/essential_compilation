package main

import (
	"github.com/jun-hf/essential_compilation/ast"
)

func IsExpression(node ast.Node) bool {
	switch n := node.(type) {
	case ast.Constant:
		if n.Type() == ast.STRING || n.Type() == ast.INT {
			return true
		}
		return false
	case *ast.BinaryOperation:
		if n.Operator() == string(ast.OP_ADD) || n.Operator() == string(ast.OP_SUB) {
			return IsExpression(n.Left) && IsExpression(n.Right)
		}
		return false
	case *ast.UnaryOperation:
		if n.Operator() == string(ast.OP_SUB) {
			return IsExpression(n.Exp)
		}
	}

	return false
}

func IsStatement(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.ExprStatment:
		return IsExpression(n.Value)
	}
	return false
}

func IsLanguageINT(p ast.Porgram) bool {
	for _, stmt := range p.Body {
		if !IsStatement(stmt) {
			return false
		}
	}
	return true
}