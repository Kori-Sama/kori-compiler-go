package parser

import (
	"encoding/json"
	"testing"

	"github.com/Kori-Sama/kori-compiler/lexer"
)

var code = `func sort(arr, len) {
  for var i = 0; i < len; i +=1; {
      for var j = i + 1; j < len; j +=1; {
          if arr[i] > arr[j] {
              let tmp = arr[i];
              arr[i] = arr[j];
              arr[j] = tmp;
          }
      }
    }
}`

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
