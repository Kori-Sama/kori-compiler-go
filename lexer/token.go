package lexer

type TokenKind int

const (
	TOKEN_ILLEGAL TokenKind = iota
	TOKEN_NAME
	TOKEN_NUMBER
	TOKEN_STRING
	TOKEN_ASSIGN
	TOKEN_EQ
	TOKEN_NOT_EQ
	TOKEN_BANG
	TOKEN_LESS
	TOKEN_GREATER
	TOKEN_LESS_EQ
	TOKEN_GREATER_EQ
	TOKEN_SEMI
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_COMMA
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_SLASH
	TOKEN_STAR
	TOKEN_EOF
	// Keyword
	TOKEN_FUNC
	TOKEN_LET
	TOKEN_RETURN
	TOKEN_IF
	TOKEN_ELSE_IF
	TOKEN_ELSE
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_FOR
)

type Token struct {
	Kind    TokenKind
	Literal string
}

func NewToken(kind TokenKind, literal string) *Token {
	return &Token{
		Kind:    kind,
		Literal: literal,
	}
}

func (t *Token) String() string {
	return t.Literal
}

func (t *Token) Equals(other *Token) bool {
	return t.Kind == other.Kind && t.Literal == other.Literal
}

func (t TokenKind) String() string {
	switch t {
	case TOKEN_ILLEGAL:
		return "ILLEGAL"
	case TOKEN_NAME:
		return "NAME"
	case TOKEN_NUMBER:
		return "NUMBER"
	case TOKEN_STRING:
		return "STRING"
	case TOKEN_ASSIGN:
		return "ASSIGN"
	case TOKEN_EQ:
		return "EQ"
	case TOKEN_NOT_EQ:
		return "NOT_EQ"
	case TOKEN_BANG:
		return "BANG"
	case TOKEN_LESS:
		return "LESS"
	case TOKEN_GREATER:
		return "GREATER"
	case TOKEN_LESS_EQ:
		return "LESS_EQ"
	case TOKEN_GREATER_EQ:
		return "GREATER_EQ"
	case TOKEN_SEMI:
		return "SEMI"
	case TOKEN_LBRACE:
		return "LBRACE"
	case TOKEN_RBRACE:
		return "RBRACE"
	case TOKEN_LPAREN:
		return "LPAREN"
	case TOKEN_RPAREN:
		return "RPAREN"
	case TOKEN_COMMA:
		return "COMMA"
	case TOKEN_PLUS:
		return "PLUS"
	case TOKEN_MINUS:
		return "MINUS"
	case TOKEN_SLASH:
		return "SLASH"
	case TOKEN_STAR:
		return "STAR"
	case TOKEN_EOF:
		return "EOF"
	case TOKEN_FUNC:
		return "FUNC"
	case TOKEN_LET:
		return "LET"
	case TOKEN_RETURN:
		return "RETURN"
	case TOKEN_IF:
		return "IF"
	case TOKEN_ELSE_IF:
		return "ELSE_IF"
	case TOKEN_ELSE:
		return "ELSE"
	case TOKEN_TRUE:
		return "TRUE"
	case TOKEN_FALSE:
		return "FALSE"
	case TOKEN_FOR:
		return "FOR"
	default:
		return "UNKNOWN"
	}
}
