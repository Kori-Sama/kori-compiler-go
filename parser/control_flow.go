package parser

import (
	"github.com/Kori-Sama/kori-compiler/cerr"
	"github.com/Kori-Sama/kori-compiler/lexer"
)

type IfExpr struct {
	BaseExpr
	Cond Expr `json:"cond"`
	Then Expr `json:"then"`
	Else Expr `json:"else"`
}

func NewIfExpr(cond, then, else_ Expr) *IfExpr {
	return &IfExpr{
		BaseExpr: BaseExpr{Type: EXPR_IF},
		Cond:     cond,
		Then:     then,
		Else:     else_,
	}
}

func (p *Parser) parseIfExpr() (expr Expr) {
	p.nextToken()
	cond := p.parseExpr()
	if cond == nil {
		return nil
	}

	then := p.parseBraceExpr()

	var else_ Expr
	if p.getCurTok().Kind == lexer.TOKEN_ELSE {

		p.nextToken()
		else_ = p.parseBraceExpr()
	}
	expr = NewIfExpr(cond, then, else_)

	return expr
}

type ForExpr struct {
	BaseExpr
	VarName string `json:"var_name"`
	Start   Expr   `json:"start"`
	End     Expr   `json:"end"`
	Step    Expr   `json:"step"`
	Body    Expr   `json:"body"`
}

func NewForExpr(varName string, start, end, step, body Expr) *ForExpr {
	return &ForExpr{
		BaseExpr: BaseExpr{Type: EXPR_FOR},
		VarName:  varName,
		Start:    start,
		End:      end,
		Step:     step,
		Body:     body,
	}
}

// for let i = 0; i < 10; i = i + 1 { }

func (p *Parser) parseForExpr() (expr Expr) {
	p.nextToken()

	if p.getCurTok().Kind == lexer.TOKEN_LBRACE {
		body := p.parseBraceExpr()
		expr = NewForExpr("", nil, nil, nil, body)
		return expr
	}

	if p.getCurTok().Kind != lexer.TOKEN_VAR {
		p.Err = cerr.NewParserError("Expected 'var' in for loop", p.getCurTok().Line, p.getCurTok().Location)
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

type ReturnExpr struct {
	BaseExpr
	Value Expr `json:"value"`
}

func NewReturnExpr(value Expr) *ReturnExpr {
	return &ReturnExpr{
		BaseExpr: BaseExpr{Type: EXPR_RETURN},
		Value:    value,
	}
}

func (p *Parser) parseReturnExpr() (expr Expr) {
	p.nextToken()

	value := p.parseExpr()

	expr = NewReturnExpr(value)

	return expr
}
