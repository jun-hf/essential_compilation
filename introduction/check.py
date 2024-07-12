import ast

def is_exp(exp):
    match exp:
        case ast.Constant(num):
            if type(num) == int:
                return True
            return False
        case ast.BinOp(left, ast.UAdd, right):
            return is_exp(left) and is_exp(right)
        case ast.BinOp(left, ast.USub, right):
            return is_exp(left) and is_exp(right)
        case ast.UnaryOp(ast.USub, operand):
            return is_exp(operand)
    return False

def is_statement(stmt):
    match stmt:
        case ast.Expr(ast.Call(ast.Name("print"), [exp])):
            return is_exp(exp)
        case ast.Expr(exp):
            return is_exp(exp)
        
    return False

def is_Lint(p):
    match p:
        case ast.Module(stmt):
            return all([is_statement(s) for s in stmt])
    return False


print(is_Lint(ast.Module([ast.Expr(ast.BinOp(ast.Constant(8), ast.UAdd, ast.Constant(9)))])))
print(is_Lint(ast.Module([ast.Expr(ast.BinOp(ast.Constant(8), ast.Mult, ast.Constant(9)))])))