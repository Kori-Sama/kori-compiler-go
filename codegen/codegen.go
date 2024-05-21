package codegen

import (
	"errors"

	"github.com/Kori-Sama/kori-compiler/parser"
)

func GenJsCode(asts []*parser.FunctionAST) (target string, err error) {
	if hasRepeatedFunc(asts) {
		return "", errors.New("repeated function found")
	}

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

func hasRepeatedFunc(asts []*parser.FunctionAST) bool {
	funcs := make(map[string]bool)
	for _, ast := range asts {
		if ast == nil {
			continue
		}
		if _, ok := funcs[ast.Proto.Name]; ok {
			return true
		}
		funcs[ast.Proto.Name] = true
	}
	return false
}
