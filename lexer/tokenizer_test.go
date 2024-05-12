package lexer

import (
	"testing"
)

var validTokens = map[string](struct {
	src string
	tok []Token
}){
	"One_Line": {
		"let a = 9;",
		[]Token{
			{NAME, "let"},
			{NAME, "a"},
			{ASSIGN, "="},
			{NUMBER, "9"},
			{SEMI, ";"}},
	},
	"Two_Lines": {
		"let a = 9;\nlet b = 10;",
		[]Token{
			{NAME, "let"},
			{NAME, "a"},
			{ASSIGN, "="},
			{NUMBER, "9"},
			{SEMI, ";"},
			{NAME, "let"},
			{NAME, "b"},
			{ASSIGN, "="},
			{NUMBER, "10"},
			{SEMI, ";"}},
	},
	"Function": {
		"func add(a, b) {\n    return a + b;\n}",
		[]Token{
			{NAME, "func"},
			{NAME, "add"},
			{LPAREN, "("},
			{NAME, "a"},
			{COMMA, ","},
			{NAME, "b"},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{NAME, "return"},
			{NAME, "a"},
			{PLUS, "+"},
			{NAME, "b"},
			{SEMI, ";"},
			{RBRACE, "}"}},
	},

	"Comparisons": {
		"a == b; a <= b; a >= b; a !=b; a < b; a > b;",
		[]Token{
			{NAME, "a"},
			{EQ, "=="},
			{NAME, "b"},
			{SEMI, ";"},
			{NAME, "a"},
			{LESS_EQ, "<="},
			{NAME, "b"},
			{SEMI, ";"},
			{NAME, "a"},
			{GREATER_EQ, ">="},
			{NAME, "b"},
			{SEMI, ";"},
			{NAME, "a"},
			{NOT_EQ, "!="},
			{NAME, "b"},
			{SEMI, ";"},
			{NAME, "a"},
			{LESS, "<"},
			{NAME, "b"},
			{SEMI, ";"},
			{NAME, "a"},
			{GREATER, ">"},
			{NAME, "b"},
			{SEMI, ";"}},
	},
}

func TestNextToken(t *testing.T) {
	for name, test := range validTokens {
		t.Run(name, func(t *testing.T) {
			tok := NewTokenizer(&test.src)

			for _, expected := range test.tok {
				token := tok.Next()
				if token.Type != expected.Type {
					t.Errorf("Expected %s, got %s", expected.Type, token.Type)
				}
				if token.Literal != expected.Literal {
					t.Errorf("Expected %s, got %s", expected.Literal, token.Literal)
				}
			}
		})
	}
}

var invalidTokensMap = map[string]TokenizerError{
	"let a = 9; !":   {Message: "Illegal character", Line: 1, Location: 12},
	"let a = 9;\n !": {Message: "Illegal character", Line: 2, Location: 2},
}

func TestNextTokenError(t *testing.T) {
	for input, expected := range invalidTokensMap {
		tok := NewTokenizer(&input)

		for {
			token := tok.Next()
			if token.Type == ILLEGAL || token.Type == EOF {
				break
			}
		}

		if tok.Err == nil {
			continue
		}
		if tok.Err.Message != expected.Message {
			t.Errorf("Expected %s, got %s", expected.Message, tok.Err.Message)
		}
		if tok.Err.Line != expected.Line {
			t.Errorf("Expected %d, got %d", expected.Line, tok.Err.Line)
		}
		if tok.Err.Location != expected.Location {
			t.Errorf("Expected %d, got %d", expected.Location, tok.Err.Location)
		}
	}
}
