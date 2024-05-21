package parser

type Expr interface {
	ICodegen
	GetType() ExprType
}

var _ Expr = &NumberExpr{}
var _ Expr = &VariableExpr{}
var _ Expr = &BinaryExpr{}
var _ Expr = &CallExpr{}
var _ Expr = &IfExpr{}
var _ Expr = &ForExpr{}
var _ Expr = &AssignExpr{}
var _ Expr = &DeclarationExpr{}

type BaseExpr struct {
	Type ExprType `json:"type"`
}

type NumberExpr struct {
	BaseExpr
	Val float64 `json:"val"`
}

type StringExpr struct {
	BaseExpr
	Val string `json:"val"`
}

type VariableExpr struct {
	BaseExpr
	Name string `json:"name"`
}

type BinaryExpr struct {
	BaseExpr
	Op  OpKind `json:"op"`
	LHS Expr   `json:"lhs"`
	RHS Expr   `json:"rhs"`
}

type BraceExpr struct {
	BaseExpr
	Exprs []Expr `json:"exprs"`
}

type CallExpr struct {
	BaseExpr
	Callee string `json:"callee"`
	Args   []Expr `json:"args"`
}

func NewNumberExpr(val float64) *NumberExpr {
	return &NumberExpr{
		BaseExpr: BaseExpr{Type: EXPR_NUMBER},
		Val:      val,
	}
}

func NewStringExpr(val string) *StringExpr {
	return &StringExpr{
		BaseExpr: BaseExpr{Type: EXPR_STRING},
		Val:      val,
	}
}

func NewVariableExpr(name string) *VariableExpr {
	return &VariableExpr{
		BaseExpr: BaseExpr{Type: EXPR_VARIABLE},
		Name:     name,
	}
}

func NewBinaryExpr(op OpKind, lhs, rhs Expr) *BinaryExpr {
	return &BinaryExpr{
		BaseExpr: BaseExpr{Type: EXPR_BINARY},
		Op:       op,
		LHS:      lhs,
		RHS:      rhs,
	}
}

func NewBraceExpr(exprs []Expr) *BraceExpr {
	return &BraceExpr{
		BaseExpr: BaseExpr{Type: EXPR_BRACE},
		Exprs:    exprs,
	}
}

func NewCallExpr(callee string, args []Expr) *CallExpr {
	return &CallExpr{
		BaseExpr: BaseExpr{Type: EXPR_CALL},
		Callee:   callee,
		Args:     args,
	}
}

func (n *BaseExpr) GetType() ExprType {
	return n.Type
}
