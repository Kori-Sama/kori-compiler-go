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
			{TOKEN_LET, "let"},
			{TOKEN_NAME, "a"},
			{TOKEN_ASSIGN, "="},
			{TOKEN_NUMBER, "9"},
			{TOKEN_SEMI, ";"}},
	},
	"Two_Lines": {
		"let a = 9 / 9 - 1;\nlet b = 10 * 2;",
		[]Token{
			{TOKEN_LET, "let"},
			{TOKEN_NAME, "a"},
			{TOKEN_ASSIGN, "="},
			{TOKEN_NUMBER, "9"},
			{TOKEN_SLASH, "/"},
			{TOKEN_NUMBER, "9"},
			{TOKEN_MINUS, "-"},
			{TOKEN_NUMBER, "1"},
			{TOKEN_SEMI, ";"},
			{TOKEN_LET, "let"},
			{TOKEN_NAME, "b"},
			{TOKEN_ASSIGN, "="},
			{TOKEN_NUMBER, "10"},
			{TOKEN_STAR, "*"},
			{TOKEN_NUMBER, "2"},
			{TOKEN_SEMI, ";"}},
	},
	"Function": {
		"func add(a, b) {\n    return a + b;\n}",
		[]Token{
			{TOKEN_FUNC, "func"},
			{TOKEN_NAME, "add"},
			{TOKEN_LPAREN, "("},
			{TOKEN_NAME, "a"},
			{TOKEN_COMMA, ","},
			{TOKEN_NAME, "b"},
			{TOKEN_RPAREN, ")"},
			{TOKEN_LBRACE, "{"},
			{TOKEN_RETURN, "return"},
			{TOKEN_NAME, "a"},
			{TOKEN_PLUS, "+"},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_RBRACE, "}"}},
	},

	"Token_With_Double_Operators": {
		"a == b; a <= b; a >= b; a !=b; a < b; a > b;",
		[]Token{
			{TOKEN_NAME, "a"},
			{TOKEN_EQ, "=="},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_NAME, "a"},
			{TOKEN_LESS_EQ, "<="},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_NAME, "a"},
			{TOKEN_GREATER_EQ, ">="},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_NAME, "a"},
			{TOKEN_NOT_EQ, "!="},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_NAME, "a"},
			{TOKEN_LESS, "<"},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"},
			{TOKEN_NAME, "a"},
			{TOKEN_GREATER, ">"},
			{TOKEN_NAME, "b"},
			{TOKEN_SEMI, ";"}},
	},
	"Token_With_Double_Keywords": {
		"if true { return; } else if false { return; } else {}",
		[]Token{
			{TOKEN_IF, "if"},
			{TOKEN_TRUE, "true"},
			{TOKEN_LBRACE, "{"},
			{TOKEN_RETURN, "return"},
			{TOKEN_SEMI, ";"},
			{TOKEN_RBRACE, "}"},
			{TOKEN_ELSE_IF, "else if"},
			{TOKEN_FALSE, "false"},
			{TOKEN_LBRACE, "{"},
			{TOKEN_RETURN, "return"},
			{TOKEN_SEMI, ";"},
			{TOKEN_RBRACE, "}"},
			{TOKEN_ELSE, "else"},
			{TOKEN_LBRACE, "{"},
			{TOKEN_RBRACE, "}"}},
	},
}

func TestNextToken(t *testing.T) {
	for name, test := range validTokens {
		t.Run(name, func(t *testing.T) {
			tok := NewTokenizer(&test.src)

			for _, expected := range test.tok {
				token := tok.Next()
				if token.Kind != expected.Kind {
					t.Errorf("Expected %s, got %s", expected.Kind, token.Kind)
				}
				if token.Literal != expected.Literal {
					t.Errorf("Expected %s, got %s", expected.Literal, token.Literal)
				}
			}
		})
	}
}

var invalidTokensMap = map[string]TokenizerError{
	"let a = 9; !":   {Message: "Illegal character", Line: 0, Location: 11},
	"let a = 9;\n !": {Message: "Illegal character", Line: 1, Location: 1},
}

func TestNextTokenError(t *testing.T) {
	for input, expected := range invalidTokensMap {
		tok := NewTokenizer(&input)

		for {
			token := tok.Next()
			if token.Kind == TOKEN_ILLEGAL || token.Kind == TOKEN_EOF {
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
