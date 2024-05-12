package lexer

type TokenType int

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case NAME:
		return "NAME"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case ASSIGN:
		return "ASSIGN"
	case EQ:
		return "EQ"
	case NOT_EQ:
		return "NOT_EQ"
	case BANG:
		return "BANG"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case LESS_EQ:
		return "LESS_EQ"
	case GREATER_EQ:
		return "GREATER_EQ"
	case SEMI:
		return "SEMI"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

const (
	ILLEGAL TokenType = iota
	NAME
	NUMBER
	STRING
	ASSIGN
	EQ
	NOT_EQ
	BANG
	LESS
	GREATER
	LESS_EQ
	GREATER_EQ
	SEMI
	LBRACE
	RBRACE
	LPAREN
	RPAREN
	COMMA
	PLUS
	MINUS
	SLASH
	STAR
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, l string) *Token {
	return &Token{
		Type:    t,
		Literal: l,
	}
}

func (t *Token) String() string {
	return t.Literal
}
