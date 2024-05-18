package parser

import (
	"github.com/Kori-Sama/compiler-go/cerr"
	"github.com/Kori-Sama/compiler-go/lexer"
)

type IfExpr struct {
	Cond Expr `json:"cond"`
	Then Expr `json:"then"`
	Else Expr `json:"else"`
}

func NewIfExpr(cond, then, else_ Expr) *IfExpr {
	return &IfExpr{
		Cond: cond,
		Then: then,
		Else: else_,
	}
}

func (p *Parser) parseIfExpr() (expr Expr) {
	p.nextToken()
	cond := p.parseExpr()
	if cond == nil {
		return nil
	}

	then := p.parseBraceExpr()

	if p.getCurTok().Kind != lexer.TOKEN_ELSE {
		return nil
	}

	p.nextToken()

	else_ := p.parseBraceExpr()

	expr = NewIfExpr(cond, then, else_)

	return expr
}

type ForExpr struct {
	VarName string `json:"var_name"`
	Start   Expr   `json:"start"`
	End     Expr   `json:"end"`
	Step    Expr   `json:"step"`
	Body    Expr   `json:"body"`
}

func NewForExpr(varName string, start, end, step, body Expr) *ForExpr {
	return &ForExpr{
		VarName: varName,
		Start:   start,
		End:     end,
		Step:    step,
		Body:    body,
	}
}

// for let i = 0; i < 10; i = i + 1 { }

func (p *Parser) parseForExpr() (expr Expr) {
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_LET {
		p.Err = cerr.NewParserError("Expected 'let' in for loop", p.getCurTok().Line, p.getCurTok().Location)
		return nil
	}

	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected variable name in for loop", p.getCurTok().Line, p.getCurTok().Location)
		return nil
	}

	varName := p.getCurTok().Literal
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_ASSIGN {
		p.Err = cerr.NewParserError("Expected '=' in for loop", p.getCurTok().Line, p.getCurTok().Location)
		return nil
	}

	p.nextToken()

	start := p.parseExpr()
	if start == nil {
		return nil
	}

	if p.getCurTok().Kind != lexer.TOKEN_SEMI {
		p.Err = cerr.NewParserError("Expected ';' in for loop", p.getCurTok().Line, p.getCurTok().Location)
		return nil
	}

	p.nextToken()

	cond := p.parseExpr()
	if cond == nil {
		return nil
	}

	if p.getCurTok().Kind != lexer.TOKEN_SEMI {
		p.Err = cerr.NewParserError("Expected ';' in for loop", p.getCurTok().Line, p.getCurTok().Location)
		return nil
	}

	p.nextToken()

	step := p.parseAssignExpr()

	p.nextToken()

	body := p.parseBraceExpr()

	expr = NewForExpr(varName, start, cond, step, body)

	return expr
}
