package codegen

import (
	"testing"

	"github.com/Kori-Sama/kori-compiler/lexer"
	"github.com/Kori-Sama/kori-compiler/parser"
)

var code = "func foo() { print(1) }"

func TestCodegenJavascript(t *testing.T) {
	lexer := lexer.NewLexer(&code)
	parser := parser.NewParser(lexer.ParseAll())

	res := parser.Parse()

	for _, ast := range res {
		t.Log(ast)
	}

	t.Log(GenJsCode(res))
}
