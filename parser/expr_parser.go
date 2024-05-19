package parser

import (
	"fmt"
	"strconv"

	"github.com/Kori-Sama/kori-compiler/cerr"
	"github.com/Kori-Sama/kori-compiler/lexer"
)

func (p *Parser) parseExpr() Expr {
	lhs := p.parsePrimary()
	if lhs == nil {
		return nil
	}

	return p.parseBinOpRHS(0, lhs)
}

func (p *Parser) parsePrimary() Expr {
	tok := p.getCurTok()
	switch tok.Kind {
	case lexer.TOKEN_NUMBER:
		return p.parseNumberExpr()
	case lexer.TOKEN_LPAREN:
		return p.parseParenExpr()
	case lexer.TOKEN_NAME:
		return p.parseIdentifierExpr()
	case lexer.TOKEN_LBRACE:
		return p.parseBraceExpr()
	case lexer.TOKEN_IF:
		return p.parseIfExpr()
	case lexer.TOKEN_FOR:
		return p.parseForExpr()
	case lexer.TOKEN_LET, lexer.TOKEN_VAR:
		return p.parseDeclarationExpr()
	case lexer.TOKEN_RETURN:
		return p.parseReturnExpr()
	case lexer.TOKEN_SEMI:
		return nil
	case lexer.TOKEN_EOF:
		return nil
	default:
		p.Err = cerr.NewParserError(fmt.Sprintf("Unknown token '%s' when expecting an expression", tok.Literal), tok.Line, tok.Location)
		return nil
	}
}

func (p *Parser) parseBinOpRHS(exprPrec int, lhs Expr) Expr {
	for {
		tok := p.getCurTok()
		if tok == nil {
			return lhs
		}

		if !isOpKind(tok.Kind) {
			return lhs
		}

		binOp := getBinOpKind(tok.Kind)
		if binOp == OP_UNKNOWN {
			return lhs
		}

		prec := binOpPrecedence[binOp]
		if prec < exprPrec {
			return lhs
		}

		p.nextToken()

		rhs := p.parsePrimary()
		if rhs == nil {
			return nil
		}

		nextPrec := binOpPrecedence[getBinOpKind(p.getCurTok().Kind)]
		if prec < nextPrec {
			rhs = p.parseBinOpRHS(prec+1, rhs)
			if rhs == nil {
				return nil
			}
		}

		lhs = NewBinaryExpr(binOp, lhs, rhs)
	}
}

func (p *Parser) parseBraceExpr() Expr {
	p.nextToken()

	if p.getCurTok().Kind == lexer.TOKEN_RBRACE {
		p.nextToken()
		return nil
	}

	var exprs []Expr
	for {
		if p.getCurTok().Kind == lexer.TOKEN_RBRACE {
			break
		}

		expr := p.parseExpr()
		if expr == nil {
			return nil
		}

		exprs = append(exprs, expr)

		if p.getCurTok().Kind == lexer.TOKEN_RBRACE {
			break
		}

		if p.getCurTok().Kind != lexer.TOKEN_SEMI {
			p.Error("Expected ';' or '}' in block")
		}

		if p.getCurTok().Kind == lexer.TOKEN_SEMI {
			p.nextToken()
		}
	}

	p.nextToken()
	return NewBraceExpr(exprs)
}

func (p *Parser) parseNumberExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_NUMBER {
		p.Err = cerr.NewParserError("Expected number", tok.Line, tok.Location)
		return nil
	}

	val, err := strconv.ParseFloat(tok.Literal, 64)

	if err != nil {
		p.Err = cerr.NewParserError("Could not parse number", tok.Line, tok.Location)
		return nil
	}
	expr = NewNumberExpr(val)

	p.nextToken()

	return expr
}

func (p *Parser) parseParenExpr() (expr Expr) {
	p.nextToken()

	expr = p.parseExpr()
	if expr == nil {
		return nil
	}

	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_RPAREN {
		p.Err = cerr.NewParserError("Expected ')'", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	return expr
}

func (p *Parser) parseIdentifierExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected identifier", tok.Line, tok.Location)
		return nil
	}
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_LPAREN {
		return NewVariableExpr(tok.Literal)
	}

	p.nextToken()

	args := make([]Expr, 0)
	if p.getCurTok().Kind == lexer.TOKEN_RPAREN {
		p.nextToken()
		if p.getCurTok().Kind == lexer.TOKEN_SEMI {
			p.nextToken()
		}
		return NewCallExpr(tok.Literal, args)
	}
	for {
		arg := p.parseExpr()
		if arg == nil {
			return nil
		}
		args = append(args, arg)

		if p.getCurTok().Kind == lexer.TOKEN_RPAREN {
			break
		}

		if p.getCurTok().Kind != lexer.TOKEN_COMMA {
			p.Err = cerr.NewParserError("Expected ',' or ')' in arguments", tok.Line, tok.Location)
			return nil
		}

		p.nextToken()
	}

	p.nextToken()
	return NewCallExpr(tok.Literal, args)
}
