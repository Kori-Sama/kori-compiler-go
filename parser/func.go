package parser

import (
	"github.com/Kori-Sama/kori-compiler/cerr"
	"github.com/Kori-Sama/kori-compiler/lexer"
)

type PrototypeAST struct {
	Type string   `json:"type"`
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type FunctionAST struct {
	Type  string        `json:"type"`
	Proto *PrototypeAST `json:"proto"`
	Body  Expr          `json:"body"`
}

type LambdaExpr struct {
	BaseExpr
	Proto *PrototypeAST `json:"proto"`
	Body  Expr          `json:"body"`
}

func NewPrototypeAST(name string, args []string) *PrototypeAST {
	return &PrototypeAST{
		Type: "Prototype",
		Name: name,
		Args: args,
	}
}

func NewFunctionAST(proto *PrototypeAST, body Expr) *FunctionAST {
	return &FunctionAST{
		Type:  "Function",
		Proto: proto,
		Body:  body,
	}
}

func NewLambdaExpr(proto *PrototypeAST, body Expr) *LambdaExpr {
	return &LambdaExpr{
		BaseExpr: BaseExpr{Type: EXPR_LAMBDA},
		Proto:    proto,
		Body:     body,
	}
}

func (p *Parser) parseLambdaExpr() Expr {
	p.nextToken()

	proto := p.parsePrototype()
	if proto == nil {
		return nil
	}

	body := p.parseBraceExpr()

	return NewLambdaExpr(proto, body)
}

func (p *Parser) parsePrototype() *PrototypeAST {
	tok := p.getCurTok()
	name := ""
	if tok.Kind == lexer.TOKEN_NAME {
		name = tok.Literal
		p.nextToken()
	}

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
