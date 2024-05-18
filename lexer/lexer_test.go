package lexer

import (
	"testing"

	"github.com/Kori-Sama/compiler-go/cerr"
)

var validTokens = map[string](struct {
	src string
	tok []Token
}){
	"One_Line": {
		"let a = 9;",
		[]Token{
			{TOKEN_LET, "let", 0, 0},
			{TOKEN_NAME, "a", 0, 4},
			{TOKEN_ASSIGN, "=", 0, 6},
			{TOKEN_NUMBER, "9", 0, 8},
			{TOKEN_SEMI, ";", 0, 9}},
	},
	"Two_Lines": {
		"let a = 9 / 9 - 1;\nlet b = 10 * 2;",
		[]Token{
			{TOKEN_LET, "let", 0, 0},
			{TOKEN_NAME, "a", 0, 4},
			{TOKEN_ASSIGN, "=", 0, 6},
			{TOKEN_NUMBER, "9", 0, 8},
			{TOKEN_SLASH, "/", 0, 10},
			{TOKEN_NUMBER, "9", 0, 12},
			{TOKEN_MINUS, "-", 0, 14},
			{TOKEN_NUMBER, "1", 0, 16},
			{TOKEN_SEMI, ";", 0, 17},
			{TOKEN_LET, "let", 1, 0},
			{TOKEN_NAME, "b", 1, 4},
			{TOKEN_ASSIGN, "=", 1, 6},
			{TOKEN_NUMBER, "10", 1, 8},
			{TOKEN_STAR, "*", 1, 11},
			{TOKEN_NUMBER, "2", 1, 13},
			{TOKEN_SEMI, ";", 1, 14}},
	},
	"Function": {
		"func add(a, b) {\n    return a + b;\n}",
		[]Token{
			{TOKEN_FUNC, "func", 0, 0},
			{TOKEN_NAME, "add", 0, 5},
			{TOKEN_LPAREN, "(", 0, 8},
			{TOKEN_NAME, "a", 0, 9},
			{TOKEN_COMMA, ",", 0, 10},
			{TOKEN_NAME, "b", 0, 12},
			{TOKEN_RPAREN, ")", 0, 13},
			{TOKEN_LBRACE, "{", 0, 15},
			{TOKEN_RETURN, "return", 1, 4},
			{TOKEN_NAME, "a", 1, 11},
			{TOKEN_PLUS, "+", 1, 13},
			{TOKEN_NAME, "b", 1, 15},
			{TOKEN_SEMI, ";", 1, 16},
			{TOKEN_RBRACE, "}", 2, 0}},
	},

	"Token_With_Double_Operators": {
		"a == b; a <= b; a >= b; a !=b; a < b; a > b;",
		[]Token{
			{TOKEN_NAME, "a", 0, 0},
			{TOKEN_EQ, "==", 0, 2},
			{TOKEN_NAME, "b", 0, 5},
			{TOKEN_SEMI, ";", 0, 6},
			{TOKEN_NAME, "a", 0, 8},
			{TOKEN_LESS_EQ, "<=", 0, 10},
			{TOKEN_NAME, "b", 0, 13},
			{TOKEN_SEMI, ";", 0, 14},
			{TOKEN_NAME, "a", 0, 16},
			{TOKEN_GREATER_EQ, ">=", 0, 18},
			{TOKEN_NAME, "b", 0, 21},
			{TOKEN_SEMI, ";", 0, 22},
			{TOKEN_NAME, "a", 0, 24},
			{TOKEN_NOT_EQ, "!=", 0, 26},
			{TOKEN_NAME, "b", 0, 28},
			{TOKEN_SEMI, ";", 0, 29},
			{TOKEN_NAME, "a", 0, 31},
			{TOKEN_LESS, "<", 0, 33},
			{TOKEN_NAME, "b", 0, 35},
			{TOKEN_SEMI, ";", 0, 36},
			{TOKEN_NAME, "a", 0, 38},
			{TOKEN_GREATER, ">", 0, 40},
			{TOKEN_NAME, "b", 0, 42},
			{TOKEN_SEMI, ";", 0, 43}},
	},
	"Token_With_Double_Keywords": {
		"if true { return; } else if false { return; } else {}",
		[]Token{
			{TOKEN_IF, "if", 0, 0},
			{TOKEN_TRUE, "true", 0, 3},
			{TOKEN_LBRACE, "{", 0, 8},
			{TOKEN_RETURN, "return", 0, 10},
			{TOKEN_SEMI, ";", 0, 16},
			{TOKEN_RBRACE, "}", 0, 18},
			{TOKEN_ELSE_IF, "else if", 0, 20},
			{TOKEN_FALSE, "false", 0, 28},
			{TOKEN_LBRACE, "{", 0, 34},
			{TOKEN_RETURN, "return", 0, 36},
			{TOKEN_SEMI, ";", 0, 42},
			{TOKEN_RBRACE, "}", 0, 44},
			{TOKEN_ELSE, "else", 0, 46},
			{TOKEN_LBRACE, "{", 0, 51},
			{TOKEN_RBRACE, "}", 0, 52}},
	},
	"Token_Array": {
		"[1, 2, 3]",
		[]Token{
			{TOKEN_LBRACKET, "[", 0, 0},
			{TOKEN_NUMBER, "1", 0, 1},
			{TOKEN_COMMA, ",", 0, 2},
			{TOKEN_NUMBER, "2", 0, 4},
			{TOKEN_COMMA, ",", 0, 5},
			{TOKEN_NUMBER, "3", 0, 7},
			{TOKEN_RBRACKET, "]", 0, 8}},
	},
	"Out_Of_Range": {
		"if 1234 {} else {}",
		[]Token{
			{TOKEN_IF, "if", 0, 0},
			{TOKEN_NUMBER, "1234", 0, 3},
			{TOKEN_LBRACE, "{", 0, 8},
			{TOKEN_RBRACE, "}", 0, 9},
			{TOKEN_ELSE, "else", 0, 11},
			{TOKEN_LBRACE, "{", 0, 16},
			{TOKEN_RBRACE, "}", 0, 17}},
	},
}

func TestNextToken(t *testing.T) {
	for name, test := range validTokens {
		t.Run(name, func(t *testing.T) {
			lexer := NewLexer(&test.src)

			if name == "Two_Lines" {
				t.Log("test")
			}

			for _, expected := range test.tok {
				token := lexer.Next()
				if token.Kind != expected.Kind {
					t.Errorf("Expected %s, got %s", expected.Kind, token.Kind)
				}
				if token.Literal != expected.Literal {
					t.Errorf("Expected %s, got %s", expected.Literal, token.Literal)
				}

				if token.Line != expected.Line {
					t.Errorf("Expected Line %d, got %d", expected.Line, token.Line)
				}

				if token.Location != expected.Location {
					t.Errorf("Expected Location %d, got %d", expected.Location, token.Location)
				}
			}
		})
	}
}

var invalidTokensMap = map[string]cerr.LexerError{
	"let a = 9; !":   {Message: "Illegal character", Line: 0, Location: 11},
	"let a = 9;\n !": {Message: "Illegal character", Line: 1, Location: 1},
}

func TestNextTokenError(t *testing.T) {
	for input, expected := range invalidTokensMap {
		lexer := NewLexer(&input)

		for {
			token := lexer.Next()
			if token.Kind == TOKEN_ILLEGAL || token.Kind == TOKEN_EOF {
				break
			}
		}

		if lexer.Err == nil {
			continue
		}
		if lexer.Err.Message != expected.Message {
			t.Errorf("Expected %s, got %s", expected.Message, lexer.Err.Message)
		}
		if lexer.Err.Line != expected.Line {
			t.Errorf("Expected %d, got %d", expected.Line, lexer.Err.Line)
		}
		if lexer.Err.Location != expected.Location {
			t.Errorf("Expected %d, got %d", expected.Location, lexer.Err.Location)
		}
	}
}
