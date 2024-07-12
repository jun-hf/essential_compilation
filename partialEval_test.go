package main

import (
	"testing"

	"github.com/jun-hf/essential_compilation/ast"
)

func TestPartialEval(t *testing.T) {
	// 1 + (1 + 3)
	p := &ast.Program{
		Body: []ast.Statement{
			&ast.ExprStatment{Value: &ast.BinaryOperation{
				Left: &ast.Constant{Value: 80, Literal: "80"},
				Op: ast.OP_ADD,
				Right: &ast.BinaryOperation{
					Left: &ast.Constant{Value: 67, Literal: "67"},
					Op: ast.OP_SUB,
					Right: &ast.Constant{Value: 50, Literal: "50"},
				},
			}},
		},
	}
	program, err := PartialEval(p)
	if len(err) != 0 {
		t.Error("PartialEval: failed", err)
	}
	if len(program.Body) != 1 {
		t.Errorf("program.Body wrong number of statements: got=%d, expected=%d", len(program.Body), 1)
	}
	exprStat, ok := program.Body[0].(*ast.ExprStatment)
	if !ok {
		t.Errorf("program body wrong type, got=%T, want=%T", program.Body[0], exprStat)
	}
	cons, ok := exprStat.Value.(*ast.Constant)
	if !ok {
		t.Errorf("exprStat wrong type, got=%T, want=%T", exprStat, cons)
	}
	if cons.Value != 97 {
		t.Errorf("constant incorrect value, got=%d, want=%d", cons.Value, 97)
	}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
}