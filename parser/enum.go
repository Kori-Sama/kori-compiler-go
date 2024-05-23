package parser

import "github.com/Kori-Sama/kori-compiler/lexer"

type ExprType string

const (
	EXPR_NUMBER       ExprType = "Number"
	EXPR_BOOLEAN      ExprType = "Boolean"
	EXPR_STRING       ExprType = "String"
	EXPR_VARIABLE     ExprType = "Variable"
	EXPR_ARRAY        ExprType = "Array"
	EXPR_BINARY       ExprType = "Binary"
	EXPR_UNARY        ExprType = "Unary"
	EXPR_CALL         ExprType = "Call"
	EXPR_IF           ExprType = "If"
	EXPR_FOR          ExprType = "For"
	EXPR_FOREACH      ExprType = "Foreach"
	EXPR_DECLARATION  ExprType = "Declaration"
	EXPR_ASSIGN       ExprType = "Assign"
	EXPR_RETURN       ExprType = "Return"
	EXPR_BRACE        ExprType = "Brace"
	EXPR_LAMBDA       ExprType = "Lambda"
	EXPR_INDEX        ExprType = "Index"
	EXPR_INDEX_ASSIGN ExprType = "IndexAssign"
)

type OpKind string

const (
	OP_ADD         OpKind = "+"
	OP_SUB         OpKind = "-"
	OP_MUL         OpKind = "*"
	OP_DIV         OpKind = "/"
	OP_LESS        OpKind = "<"
	OP_GREATER     OpKind = ">"
	OP_LESS_EQ     OpKind = "<="
	OP_GREATER_EQ  OpKind = ">="
	OP_EQ          OpKind = "=="
	OP_AND         OpKind = "&"
	OP_OR          OpKind = "|"
	OP_LOGICAL_AND OpKind = "&&"
	OP_LOGICAL_OR  OpKind = "||"
	OP_NOT         OpKind = "!"
	OP_UNKNOWN     OpKind = "UNKNOWN"
)

func isOpKind(tokenKind lexer.TokenKind) bool {
	switch tokenKind {
	case lexer.TOKEN_LBRACKET,
		lexer.TOKEN_LESS, lexer.TOKEN_GREATER,
		lexer.TOKEN_LESS_EQ, lexer.TOKEN_GREATER_EQ,
		lexer.TOKEN_AND, lexer.TOKEN_OR,
		lexer.TOKEN_LOGICAL_AND, lexer.TOKEN_LOGICAL_OR,
		lexer.TOKEN_EQ, lexer.TOKEN_PLUS, lexer.TOKEN_MINUS, lexer.TOKEN_STAR, lexer.TOKEN_SLASH:
		return true
	default:
		return false
	}
}

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
	case lexer.TOKEN_AND:
		return OP_AND
	case lexer.TOKEN_OR:
		return OP_OR
	case lexer.TOKEN_LOGICAL_AND:
		return OP_LOGICAL_AND
	case lexer.TOKEN_LOGICAL_OR:
		return OP_LOGICAL_OR
	default:
		return OP_UNKNOWN
	}
}
