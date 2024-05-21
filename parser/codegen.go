package parser

import "fmt"

// target: javascript

type ICodegen interface {
	Codegen() string
}

func (n *NumberExpr) Codegen() string {
	return fmt.Sprintf("%f", n.Val)
}

func (n *StringExpr) Codegen() string {
	return fmt.Sprintf(`"%s"`, n.Val)
}

func (n *VariableExpr) Codegen() string {
	return n.Name
}

func (n *BinaryExpr) Codegen() string {
	return fmt.Sprintf("(%s %s %s)", n.LHS.Codegen(), n.Op, n.RHS.Codegen())
}

func (n *CallExpr) Codegen() string {
	args := ""
	for i, arg := range n.Args {
		if i > 0 {
			args += ", "
		}
		if arg == nil {
			continue
		}
		args += arg.Codegen()
	}

	if n.Callee == "println" {
		return fmt.Sprintf("console.log(%s)", args)
	}

	return fmt.Sprintf("%s(%s)", n.Callee, args)
}

func (n *IfExpr) Codegen() string {
	if n.Else == nil {
		return fmt.Sprintf("if (%s) { %s }", n.Cond.Codegen(), n.Then.Codegen())
	}
	return fmt.Sprintf("if (%s) { %s } else { %s }", n.Cond.Codegen(), n.Then.Codegen(), n.Else.Codegen())
}

func (n *ForExpr) Codegen() string {
	return fmt.Sprintf("for (let %s = %s; %s ; %s) { %s }", n.VarName, n.Start.Codegen(), n.End.Codegen(), n.Step.Codegen(), n.Body.Codegen())
}

func (n *AssignExpr) Codegen() string {
	return fmt.Sprintf("%s = %s", n.VarName, n.Expr.Codegen())
}

func (n *DeclarationExpr) Codegen() string {
	kind := "const"
	if n.Mutable {
		kind = "let"
	}
	return fmt.Sprintf("%s %s = %s", kind, n.VarName, n.Expr.Codegen())
}

func (n *BraceExpr) Codegen() string {
	exprs := ""
	for _, expr := range n.Exprs {
		exprs += expr.Codegen() + ";"
	}
	return exprs
}

func (n *ReturnExpr) Codegen() string {
	return fmt.Sprintf("return %s", n.Value.Codegen())
}

func (n *FunctionAST) Codegen() string {
	return fmt.Sprintf("%s { %s }", n.Proto.Codegen(), n.Body.Codegen())
}

func (n *PrototypeAST) Codegen() string {
	args := ""

	for _, arg := range n.Args {
		if len(args) > 0 {
			args += ", "
		}
		args += arg
	}

	return fmt.Sprintf("function %s(%s)", n.Name, args)
}
