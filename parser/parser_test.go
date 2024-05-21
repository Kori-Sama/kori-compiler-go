package parser

import (
	"encoding/json"
	"testing"

	"github.com/Kori-Sama/kori-compiler/lexer"
)

var code = "func foo() { let str = \"123\"; }"

func TestParser(t *testing.T) {
	lexer := lexer.NewLexer(&code)
	parser := NewParser(lexer.ParseAll())
	res := parser.Parse()

	if parser.Err != nil {
		t.Error(parser.Err)
	}

	ast, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(ast))
}
