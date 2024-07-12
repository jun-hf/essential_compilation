from ast import *

def partial_Evaluator(node: Module):
    match node:
        case Module(body):
            return Module(partial_body(body))
    return Module([])


def partial_body(body):
    stmt_list = []
    for s in body:
        stmt_list.append(partial_stmt(s))
    return stmt_list

def partial_stmt(s: stmt):
    match s:
        case Expr(exp):
            return Expr(partial_exp(exp))
        case Expr(Call(Name("print", [args]))):
            return Expr(Call(Name("print", [partial_exp(args)])))
    return Expr()

def partial_exp(exp: Expression):
    match exp:
        case BinOp(left, Add(), right):
            return pe_add(left, right)
        case BinOp(left, Sub(), right):
            return pe_sub(left, right)
        case UnaryOp(USub(), Exp):
            return pe_unSub(Exp)
        case Constant(num):
            return exp
        
def pe_add(left, right):
    match (left, right):
        case (Constant(n), Constant(n2)):
            return Constant(n + n2)
    return BinOp(left, Add(), right)

def pe_sub(left, right):
    match (left, right):
        case (Constant(n), Constant(n2)):
            return Constant(n - n2)
    return BinOp(left, Sub(), right)
    
def pe_unSub(operand):
    match operand:
        case Constant(num):
            return Constant(-num)
    return operand