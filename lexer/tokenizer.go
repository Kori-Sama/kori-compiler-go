package lexer

type ITokenizer interface {
	Next() *Token
}

type Tokenizer struct {
	Err     *TokenizerError
	content *string
	current int
	line    int
	linePos int
	ch      byte
}

func NewTokenizer(content *string) *Tokenizer {
	return &Tokenizer{
		Err:     nil,
		content: content,
		current: 0,
		line:    1,
		linePos: 0,
	}
}

func (t *Tokenizer) Next() *Token {
	text := t.content

	if t.current >= len(*text) {
		return NewToken(EOF, "")
	}

	t.ch = (*text)[t.current]
	if isWhitespace(t.ch) {
		t.SkipWhitespace()
	}

	if isLetter(t.ch) {
		return t.readName()
	}
	if isDigit(t.ch) {
		return t.readNumber()
	}

	token := t.readSymbol()
	if token.Type == ILLEGAL {
		t.Err = NewTokenizerError("Illegal character", t.line, t.current-t.linePos)
	}
	return token
}

func (t *Tokenizer) PeekToken() *Token {
	current := t.current
	token := t.Next()
	t.current = current
	return token
}

func (t *Tokenizer) SkipWhitespace() {
	text := t.content

	for t.current < len(*text) {
		ch := (*text)[t.current]
		if !isWhitespace(ch) {
			break
		}

		if ch == '\n' {
			t.line++
			t.linePos = t.current + 1
		}

		t.current++
	}
	t.ch = (*text)[t.current]
}

func (t *Tokenizer) readSymbol() *Token {
	t.current++
	switch t.ch {
	case '=':
		if t.peekChar('=') {
			return NewToken(EQ, "==")
		}
		return NewToken(ASSIGN, "=")
	case '!':
		if t.peekChar('=') {
			return NewToken(NOT_EQ, "!=")
		}
		return NewToken(BANG, "!")
	case '<':
		if t.peekChar('=') {
			return NewToken(LESS_EQ, "<=")
		}
		return NewToken(LESS, "<")
	case '>':
		if t.peekChar('=') {
			return NewToken(GREATER_EQ, ">=")
		}
		return NewToken(GREATER, ">")
	case ';':
		return NewToken(SEMI, ";")
	case '{':
		return NewToken(LBRACE, "{")
	case '}':
		return NewToken(RBRACE, "}")
	case '(':
		return NewToken(LPAREN, "(")
	case ')':
		return NewToken(RPAREN, ")")
	case ',':
		return NewToken(COMMA, ",")
	case '+':
		return NewToken(PLUS, "+")
	case '-':
		return NewToken(MINUS, "-")
	case '/':
		return NewToken(SLASH, "/")
	default:
		return NewToken(ILLEGAL, string(t.ch))
	}
}

func (t *Tokenizer) peekChar(expect byte) bool {
	if t.current >= len(*t.content) {
		return false
	}
	ch := (*t.content)[t.current]
	if ch != expect {
		return false
	}
	t.current++
	return true
}

func (t *Tokenizer) readName() *Token {
	text := t.content
	start := t.current

	for t.current < len(*text) {
		ch := (*text)[t.current]
		if !isLetter(ch) {
			break
		}
		t.current++
	}
	return NewToken(NAME, (*text)[start:t.current])
}

func (t *Tokenizer) readNumber() *Token {
	text := t.content
	start := t.current

	for t.current < len(*text) {
		ch := (*text)[t.current]
		if !isDigit(ch) {
			break
		}
		t.current++
	}
	return NewToken(NUMBER, (*text)[start:t.current])
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
