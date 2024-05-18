package parser

import (
	"github.com/Kori-Sama/compiler-go/cerr"
	"github.com/Kori-Sama/compiler-go/lexer"
)

type AssignExpr struct {
	VarName string `json:"var_name"`
	Expr    Expr   `json:"expr"`
}

func NewAssignExpr(varName string, expr Expr) *AssignExpr {
	return &AssignExpr{
		VarName: varName,
		Expr:    expr,
	}
}

func (p *Parser) parseAssignExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected variable name in assignment", tok.Line, tok.Location)
		return nil
	}

	varName := tok.Literal
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_ASSIGN {
		p.Err = cerr.NewParserError("Expected '=' in assignment", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	expr = p.parseExpr()
	if expr == nil {
		return nil
	}

	return NewAssignExpr(varName, expr)
}

type InitExpr struct {
	VarName string `json:"var_name"`
	Expr    Expr   `json:"expr"`
}

func NewInitExpr(varName string, expr Expr) *InitExpr {
	return &InitExpr{
		VarName: varName,
		Expr:    expr,
	}
}

func (p *Parser) parseInitExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_LET {
		p.Err = cerr.NewParserError("Expected 'let' in initialization", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected variable name in initialization", tok.Line, tok.Location)
		return nil
	}

	varName := p.getCurTok().Literal
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_ASSIGN {
		p.Err = cerr.NewParserError("Expected '=' in initialization", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	expr = p.parseExpr()
	if expr == nil {
		return nil
	}

	return NewInitExpr(varName, expr)
}
