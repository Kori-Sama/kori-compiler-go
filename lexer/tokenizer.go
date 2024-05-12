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
		line:    0,
		linePos: 0,
	}
}

func (t *Tokenizer) Next() *Token {
	text := t.content

	if t.current >= len(*text) {
		return NewToken(TOKEN_EOF, "")
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
	if token.Kind == TOKEN_ILLEGAL {
		t.Err = NewTokenizerError("Illegal character", t.line, t.current-t.linePos)
	}
	return token
}

func (t *Tokenizer) readSymbol() *Token {
	t.current++
	switch t.ch {
	case '=':
		if t.peekChar('=') {
			return NewToken(TOKEN_EQ, "==")
		}
		return NewToken(TOKEN_ASSIGN, "=")
	case '!':
		if t.peekChar('=') {
			return NewToken(TOKEN_NOT_EQ, "!=")
		}
		return NewToken(TOKEN_BANG, "!")
	case '<':
		if t.peekChar('=') {
			return NewToken(TOKEN_LESS_EQ, "<=")
		}
		return NewToken(TOKEN_LESS, "<")
	case '>':
		if t.peekChar('=') {
			return NewToken(TOKEN_GREATER_EQ, ">=")
		}
		return NewToken(TOKEN_GREATER, ">")
	case ';':
		return NewToken(TOKEN_SEMI, ";")
	case '{':
		return NewToken(TOKEN_LBRACE, "{")
	case '}':
		return NewToken(TOKEN_RBRACE, "}")
	case '(':
		return NewToken(TOKEN_LPAREN, "(")
	case ')':
		return NewToken(TOKEN_RPAREN, ")")
	case ',':
		return NewToken(TOKEN_COMMA, ",")
	case '+':
		return NewToken(TOKEN_PLUS, "+")
	case '-':
		return NewToken(TOKEN_MINUS, "-")
	case '/':
		return NewToken(TOKEN_SLASH, "/")
	default:
		return NewToken(TOKEN_ILLEGAL, string(t.ch))
	}
}

func (t *Tokenizer) readKeyword(name string) *Token {
	switch name {
	case "let":
		return NewToken(TOKEN_LET, "let")
	case "func":
		return NewToken(TOKEN_FUNC, "func")
	case "return":
		return NewToken(TOKEN_RETURN, "return")
	case "if":
		return NewToken(TOKEN_IF, "if")
	case "else":
		if t.PeekToken(TOKEN_IF) {
			return NewToken(TOKEN_ELSE_IF, "else if")
		}
		return NewToken(TOKEN_ELSE, "else")
	case "true":
		return NewToken(TOKEN_TRUE, "true")
	case "false":
		return NewToken(TOKEN_FALSE, "false")
	case "for":
		return NewToken(TOKEN_FOR, "for")
	default:
		return nil
	}
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

	name := (*text)[start:t.current]

	if keywordToken := t.readKeyword(name); keywordToken != nil {
		return keywordToken
	}

	return NewToken(TOKEN_NAME, name)
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
	return NewToken(TOKEN_NUMBER, (*text)[start:t.current])
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
			t.linePos = t.current
		}

		t.current++
	}
	t.ch = (*text)[t.current]
}

func (t *Tokenizer) PeekToken(expect TokenKind) bool {
	start := t.current
	token := t.Next()
	if token.Kind != expect {
		t.current = start
		return false
	}
	return true
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

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
