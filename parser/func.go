package parser

import (
	"github.com/Kori-Sama/compiler-go/cerr"
	"github.com/Kori-Sama/compiler-go/lexer"
)

type PrototypeAST struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type FunctionAST struct {
	Proto *PrototypeAST `json:"proto"`
	Body  Expr          `json:"body"`
}

func NewPrototypeAST(name string, args []string) *PrototypeAST {
	return &PrototypeAST{
		Name: name,
		Args: args,
	}
}

func (p *Parser) parsePrototype() *PrototypeAST {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected function name in prototype", tok.Line, tok.Location)
		return nil
	}

	name := tok.Literal
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_LPAREN {
		p.Err = cerr.NewParserError("Expected '(' in prototype", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	args := make([]string, 0)
	for p.getCurTok().Kind != lexer.TOKEN_RPAREN {
		arg := p.getCurTok()
		if arg.Kind != lexer.TOKEN_NAME {
			p.Err = cerr.NewParserError("Expected argument name in prototype", arg.Line, arg.Location)
			return nil
		}
		args = append(args, arg.Literal)
		p.nextToken()

		if p.getCurTok().Kind != lexer.TOKEN_COMMA {
			break
		}
		p.nextToken()
	}

	if p.getCurTok().Kind != lexer.TOKEN_RPAREN {
		p.Err = cerr.NewParserError("Expected ')' in prototype", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	return NewPrototypeAST(name, args)
}

func NewFunctionAST(proto *PrototypeAST, body Expr) *FunctionAST {
	return &FunctionAST{
		Proto: proto,
		Body:  body,
	}
}

func (p *Parser) parseFunction() *FunctionAST {
	p.nextToken()

	proto := p.parsePrototype()
	if proto == nil {
		return nil
	}

	body := p.parseBraceExpr()
	if body == nil {
		return nil
	}

	return NewFunctionAST(proto, body)
}
