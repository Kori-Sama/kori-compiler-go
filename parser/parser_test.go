package parser

import (
	"encoding/json"
	"testing"

	"github.com/Kori-Sama/compiler-go/lexer"
)

var code = "for let i = 1; i < 2; i = i + 1 {1}; \n if i < 2 {2} else {3}"

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
