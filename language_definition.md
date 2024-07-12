# Language

```latex
exp ::= Constant(int) | Call(Name("input_int"), []) | UnaryOp(USub(), exp) | BinOp(exp, Add(), exp) | BinOp(exp, Sub(), exp)

stmt ::= Expr(Call(Name("print"), [exp])) | Expr(exp)

Language ::= Module(stmt*)

```