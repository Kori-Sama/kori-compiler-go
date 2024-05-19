package parser

import (
	"github.com/Kori-Sama/kori-compiler/cerr"
	"github.com/Kori-Sama/kori-compiler/lexer"
)

type AssignExpr struct {
	BaseExpr
	VarName string `json:"var_name"`
	Expr    Expr   `json:"expr"`
}

func NewAssignExpr(varName string, expr Expr) *AssignExpr {
	return &AssignExpr{
		BaseExpr: BaseExpr{Type: EXPR_ASSIGN},
		VarName:  varName,
		Expr:     expr,
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

type DeclarationExpr struct {
	BaseExpr
	VarName string `json:"var_name"`
	Mutable bool   `json:"mutable"`
	Kind    string `json:"kind"`
	Expr    Expr   `json:"expr"`
}

func NewDeclarationExpr(varName string, mutable bool, expr Expr) *DeclarationExpr {
	var kind string
	if mutable {
		kind = "var"
	} else {
		kind = "let"
	}
	return &DeclarationExpr{
		BaseExpr: BaseExpr{Type: EXPR_DECLARATION},
		VarName:  varName,
		Mutable:  mutable,
		Kind:     kind,
		Expr:     expr,
	}
}

func (p *Parser) parseDeclarationExpr() (expr Expr) {
	tok := p.getCurTok()

	mutable := false

	if tok.Kind == lexer.TOKEN_LET {

	} else if tok.Kind == lexer.TOKEN_VAR {
		mutable = true
	} else {
		p.Err = cerr.NewParserError("Expected 'let' or 'var' in Declaration", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected variable name in Declaration", tok.Line, tok.Location)
		return nil
	}

	varName := p.getCurTok().Literal
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_ASSIGN {
		p.Err = cerr.NewParserError("Expected '=' in Declaration", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	expr = p.parseExpr()
	if expr == nil {
		return nil
	}

	return NewDeclarationExpr(varName, mutable, expr)
}
