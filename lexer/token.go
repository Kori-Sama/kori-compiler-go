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
	TOKEN_COLON
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACKET
	TOKEN_RBRACKET
	TOKEN_COMMA
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_SLASH
	TOKEN_STAR
	TOKEN_EOF
	// Keyword
	TOKEN_FUNC
	TOKEN_LET
	TOKEN_VAR
	TOKEN_RETURN
	TOKEN_IF
	TOKEN_ELSE_IF
	TOKEN_ELSE
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_FOR
)

type Token struct {
	Kind     TokenKind
	Literal  string
	Line     int
	Location int
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
	if val, ok := tokenNamesMap[t]; ok {
		return val
	}
	return "UNKNOWN"
}

var tokenNamesMap = map[TokenKind]string{
	TOKEN_ILLEGAL:    "ILLEGAL",
	TOKEN_NAME:       "NAME",
	TOKEN_NUMBER:     "NUMBER",
	TOKEN_STRING:     "STRING",
	TOKEN_ASSIGN:     "ASSIGN",
	TOKEN_EQ:         "EQ",
	TOKEN_NOT_EQ:     "NOT_EQ",
	TOKEN_BANG:       "BANG",
	TOKEN_LESS:       "LESS",
	TOKEN_GREATER:    "GREATER",
	TOKEN_LESS_EQ:    "LESS_EQ",
	TOKEN_GREATER_EQ: "GREATER_EQ",
	TOKEN_SEMI:       "SEMI",
	TOKEN_COLON:      "COLON",
	TOKEN_LBRACE:     "LBRACE",
	TOKEN_RBRACE:     "RBRACE",
	TOKEN_LPAREN:     "LPAREN",
	TOKEN_RPAREN:     "RPAREN",
	TOKEN_LBRACKET:   "LBRACKET",
	TOKEN_RBRACKET:   "RBRACKET",
	TOKEN_COMMA:      "COMMA",
	TOKEN_PLUS:       "PLUS",
	TOKEN_MINUS:      "MINUS",
	TOKEN_SLASH:      "SLASH",
	TOKEN_STAR:       "STAR",
	TOKEN_EOF:        "EOF",
	TOKEN_FUNC:       "FUNC",
	TOKEN_LET:        "LET",
	TOKEN_VAR:        "VAR",
	TOKEN_RETURN:     "RETURN",
	TOKEN_IF:         "IF",
	TOKEN_ELSE_IF:    "ELSE_IF",
	TOKEN_ELSE:       "ELSE",
	TOKEN_TRUE:       "TRUE",
	TOKEN_FALSE:      "FALSE",
	TOKEN_FOR:        "FOR",
}
