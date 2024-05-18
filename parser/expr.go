package parser

import (
	"fmt"

	"github.com/Kori-Sama/compiler-go/lexer"
)

type Expr interface {
	// String() string
}

var _ Expr = &NumberExpr{}
var _ Expr = &VariableExpr{}
var _ Expr = &BinaryExpr{}
var _ Expr = &CallExpr{}

type NumberExpr struct {
	Val float64 `json:"val"`
}

func (n *NumberExpr) String() string {
	return fmt.Sprintf("Number Expr: [%f]", n.Val)
}

type VariableExpr struct {
	Name string `json:"name"`
}

func (v *VariableExpr) String() string {
	return fmt.Sprintf("Variable Expr: [%s]", v.Name)
}

type OpKind int

func (o OpKind) String() string {
	switch o {
	case OP_ADD:
		return "+"
	case OP_SUB:
		return "-"
	case OP_MUL:
		return "*"
	case OP_DIV:
		return "/"
	case OP_LESS:
		return "<"
	case OP_GREATER:
		return ">"
	case OP_LESS_EQ:
		return "<="
	case OP_GREATER_EQ:
		return ">="
	case OP_EQ:
		return "=="
	default:
		return "UNKNOWN"
	}
}

const (
	OP_ADD OpKind = iota
	OP_SUB
	OP_MUL
	OP_DIV
	OP_LESS
	OP_GREATER
	OP_LESS_EQ
	OP_GREATER_EQ
	OP_EQ
	OP_UNKNOWN
)

func getBinOpKind(kind lexer.TokenKind) OpKind {
	switch kind {
	case lexer.TOKEN_LESS:
		return OP_LESS
	case lexer.TOKEN_GREATER:
		return OP_GREATER
	case lexer.TOKEN_LESS_EQ:
		return OP_LESS_EQ
	case lexer.TOKEN_GREATER_EQ:
		return OP_GREATER_EQ
	case lexer.TOKEN_EQ:
		return OP_EQ
	case lexer.TOKEN_PLUS:
		return OP_ADD
	case lexer.TOKEN_MINUS:
		return OP_SUB
	case lexer.TOKEN_STAR:
		return OP_MUL
	case lexer.TOKEN_SLASH:
		return OP_DIV
	default:
		return OP_UNKNOWN
	}
}

type BinaryExpr struct {
	Op  OpKind `json:"op"`
	LHS Expr   `json:"lhs"`
	RHS Expr   `json:"rhs"`
}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf(
		`Binary Expr: 
		%s
		|----- %s
		|----- %s
		`,
		b.Op, b.LHS, b.RHS)
}

type CallExpr struct {
	Callee string `json:"callee"`
	Args   []Expr `json:"args"`
}

func (c *CallExpr) String() string {
	return fmt.Sprintf("Call Expr: [%s] [%v]", c.Callee, c.Args)
}

func NewNumberExpr(val float64) *NumberExpr {
	return &NumberExpr{
		Val: val,
	}
}

func NewVariableExpr(name string) *VariableExpr {
	return &VariableExpr{
		Name: name,
	}
}

func NewBinaryExpr(op OpKind, lhs, rhs Expr) *BinaryExpr {
	return &BinaryExpr{
		Op:  op,
		LHS: lhs,
		RHS: rhs,
	}
}

func NewCallExpr(callee string, args []Expr) *CallExpr {
	return &CallExpr{
		Callee: callee,
		Args:   args,
	}
}
