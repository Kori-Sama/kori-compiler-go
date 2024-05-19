package lexer

import "github.com/Kori-Sama/kori-compiler/cerr"

type Lexer struct {
	Err     *cerr.LexerError
	content *string
	current int
	line    int
	linePos int
	ch      byte
}

func NewLexer(content *string) *Lexer {
	return &Lexer{
		Err:     nil,
		content: content,
		current: 0,
		line:    0,
		linePos: 0,
	}
}

func (l *Lexer) ParseAll() []*Token {
	tokens := make([]*Token, 0)

	for {
		token := l.Next()
		tokens = append(tokens, token)
		if token.Kind == TOKEN_EOF || token.Kind == TOKEN_ILLEGAL {
			break
		}
	}

	return tokens
}

func (l *Lexer) Next() *Token {
	text := l.content

	var token *Token

	if l.current >= len(*text) {
		token = NewToken(TOKEN_EOF, "")
		return token
	}

	l.ch = (*text)[l.current]
	if isWhitespace(l.ch) {
		l.SkipWhitespace()
	}

	line := l.line
	pos := l.current - l.linePos

	if isLetter(l.ch) {
		token = l.readName()
	} else if isDigit(l.ch) {
		token = l.readNumber()
	} else {
		token = l.readSymbol()
		if token.Kind == TOKEN_ILLEGAL {
			l.Err = cerr.NewLexerError("Illegal character", l.line, l.current-l.linePos)
		}
	}

	token.Line = line
	token.Location = pos
	return token
}

func (l *Lexer) readSymbol() *Token {
	l.current++
	switch l.ch {
	case '=':
		if l.peekChar('=') {
			return NewToken(TOKEN_EQ, "==")
		}
		return NewToken(TOKEN_ASSIGN, "=")
	case '!':
		if l.peekChar('=') {
			return NewToken(TOKEN_NOT_EQ, "!=")
		}
		return NewToken(TOKEN_BANG, "!")
	case '<':
		if l.peekChar('=') {
			return NewToken(TOKEN_LESS_EQ, "<=")
		}
		return NewToken(TOKEN_LESS, "<")
	case '>':
		if l.peekChar('=') {
			return NewToken(TOKEN_GREATER_EQ, ">=")
		}
		return NewToken(TOKEN_GREATER, ">")
	case ';':
		return NewToken(TOKEN_SEMI, ";")
	case ':':
		return NewToken(TOKEN_COLON, ":")
	case '{':
		return NewToken(TOKEN_LBRACE, "{")
	case '}':
		return NewToken(TOKEN_RBRACE, "}")
	case '[':
		return NewToken(TOKEN_LBRACKET, "[")
	case ']':
		return NewToken(TOKEN_RBRACKET, "]")
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
	case '*':
		return NewToken(TOKEN_STAR, "*")
	default:
		return NewToken(TOKEN_ILLEGAL, string(l.ch))
	}
}

func (l *Lexer) readKeyword(name string) *Token {
	switch name {
	case "let":
		return NewToken(TOKEN_LET, "let")
	case "var":
		return NewToken(TOKEN_VAR, "var")
	case "func":
		return NewToken(TOKEN_FUNC, "func")
	case "return":
		return NewToken(TOKEN_RETURN, "return")
	case "if":
		return NewToken(TOKEN_IF, "if")
	case "else":
		if l.PeekToken(TOKEN_IF) {
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

func (l *Lexer) readName() *Token {
	text := l.content
	start := l.current

	for l.current < len(*text) {
		ch := (*text)[l.current]
		if !isLetter(ch) {
			break
		}
		l.current++
	}

	name := (*text)[start:l.current]

	if keywordToken := l.readKeyword(name); keywordToken != nil {
		return keywordToken
	}

	return NewToken(TOKEN_NAME, name)
}

func (l *Lexer) readNumber() *Token {
	text := l.content
	start := l.current

	for l.current < len(*text) {
		ch := (*text)[l.current]
		if !isDigit(ch) {
			break
		}
		l.current++
	}
	return NewToken(TOKEN_NUMBER, (*text)[start:l.current])
}

func (l *Lexer) SkipWhitespace() {
	text := l.content

	for l.current < len(*text) {
		ch := (*text)[l.current]
		if !isWhitespace(ch) {
			break
		}

		if ch == '\n' {
			l.line++
			l.linePos = l.current + 1
		}

		l.current++
	}
	l.ch = (*text)[l.current]
}

func (l *Lexer) PeekToken(expect TokenKind) bool {
	start := l.current
	token := l.Next()
	if token.Kind != expect {
		l.current = start
		return false
	}
	return true
}

func (l *Lexer) peekChar(expect byte) bool {
	if l.current >= len(*l.content) {
		return false
	}
	ch := (*l.content)[l.current]
	if ch != expect {
		return false
	}
	l.current++
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
