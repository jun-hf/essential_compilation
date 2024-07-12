from ast import *

def input_int() -> int:
    # entering illegal characters may cause exception,
    # but we won't worry about that
    x = int(input())
    # clamp to 64 bit signed number, emulating behavior of C's scanf
    return x

read = Call(Name("input_int"), [])
ast1_1 = BinOp(read, Add(), UnaryOp(USub(), Constant(8)))


def intrepret_expression(node: Expression):
    match node:
        case Constant(num):
            return num
        case BinOp(left, Sub(), right):
            l = intrepret_expression(left)
            r = intrepret_expression(right)
            return l - r
        case BinOp(left, Add(), right):
            l = intrepret_expression(left)
            r = intrepret_expression(right)
            return l + r
        case UnaryOp(USub(), exp):
            e = intrepret_expression(exp)
            return -e

def interpret_statement(node: stmt):
    match node:
        case Expr(Call(Name("print"), [exp])):
            print(intrepret_expression(exp))
        case Expr(exp):
            return intrepret_expression(exp)

def intrepret_Lint(p: Module):
    match p:
        case Module(body):
            for s in body:
                interpret_statement(s)

intrepret_Lint(Module([Expr(Call(Name("print"), [BinOp(Constant(10), Add(), UnaryOp(USub(), Constant(32)))]))]))


