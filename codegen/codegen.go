package codegen

import (
	"errors"

	"github.com/Kori-Sama/kori-compiler/parser"
)

func GenJsCode(asts []*parser.FunctionAST) (target string, err error) {
	hasMain := false
	for _, ast := range asts {
		if ast == nil {
			continue
		}
		target += ast.Codegen() + "\n"

		if ast.Proto.Name == "main" {
			hasMain = true
		}
	}

	if !hasMain {
		return "", errors.New("no main function found")
	}

	return target + "\nmain();\n", nil
}
