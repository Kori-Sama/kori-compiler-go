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
	case lexer.TOKEN_TRUE, lexer.TOKEN_FALSE:
		return p.parseBooleanExpr()
	case lexer.TOKEN_STRING:
		return p.parseStringExpr()
	case lexer.TOKEN_LBRACKET:
		return p.parseArrayExpr()
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
	case lexer.TOKEN_FUNC:
		return p.parseLambdaExpr()
	case lexer.TOKEN_RETURN:
		return p.parseReturnExpr()
	case lexer.TOKEN_BANG:
		return p.parseUnaryExpr()
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

func (p *Parser) parseUnaryExpr() Expr {
	var op OpKind
	if p.getCurTok().Kind == lexer.TOKEN_BANG {
		op = OP_NOT
	}

	p.nextToken()

	expr := p.parsePrimary()
	if expr == nil {
		return nil
	}

	return NewUnaryExpr(op, expr)
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

func (p *Parser) parseBooleanExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_TRUE && tok.Kind != lexer.TOKEN_FALSE {
		p.Err = cerr.NewParserError("Expected boolean", tok.Line, tok.Location)
		return nil
	}

	val := tok.Kind == lexer.TOKEN_TRUE
	expr = NewBooleanExpr(val)

	p.nextToken()

	return expr
}

func (p *Parser) parseStringExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_STRING {
		p.Err = cerr.NewParserError("Expected string", tok.Line, tok.Location)
		return nil
	}

	expr = NewStringExpr(tok.Literal)

	p.nextToken()

	return expr
}

func (p *Parser) parseArrayExpr() (expr Expr) {
	p.nextToken()

	if p.getCurTok().Kind == lexer.TOKEN_RBRACKET {
		p.nextToken()
		return nil
	}

	var values []Expr
	for {
		if p.getCurTok().Kind == lexer.TOKEN_RBRACKET {
			break
		}

		value := p.parseExpr()
		if value == nil {
			return nil
		}

		values = append(values, value)

		if p.getCurTok().Kind == lexer.TOKEN_RBRACKET {
			break
		}

		if p.getCurTok().Kind != lexer.TOKEN_COMMA {
			p.Err = cerr.NewParserError("Expected ',' or ']' in array", p.getCurTok().Line, p.getCurTok().Location)
			return nil
		}

		p.nextToken()
	}

	p.nextToken()
	return NewArrayExpr(values)
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
	if p.peekExpect(1, lexer.TOKEN_ASSIGN) {
		return p.parseAssignExpr()
	}

	if p.peekExpect(1, lexer.TOKEN_LBRACKET) {
		return p.parseIndexExpr()
	}

	if p.peekExpect(1, lexer.TOKEN_PLUS_EQ) ||
		p.peekExpect(1, lexer.TOKEN_MINUS_EQ) ||
		p.peekExpect(1, lexer.TOKEN_STAR_EQ) ||
		p.peekExpect(1, lexer.TOKEN_SLASH_EQ) {
		return p.parseAssignOpExpr()
	}

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

func (p *Parser) parseIndexExpr() (expr Expr) {
	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_NAME {
		p.Err = cerr.NewParserError("Expected identifier", tok.Line, tok.Location)
		return nil
	}
	p.nextToken()

	if p.getCurTok().Kind != lexer.TOKEN_LBRACKET {
		p.Err = cerr.NewParserError("Expected '['", tok.Line, tok.Location)
		return nil
	}
	p.nextToken()

	index := p.parseExpr()
	if index == nil {
		return nil
	}

	if p.getCurTok().Kind != lexer.TOKEN_RBRACKET {
		p.Err = cerr.NewParserError("Expected ']'", tok.Line, tok.Location)
		return nil
	}
	p.nextToken()

	if p.getCurTok().Kind == lexer.TOKEN_ASSIGN {
		p.nextToken()
		value := p.parseExpr()
		if value == nil {
			return nil
		}

		return NewIndexAssignExpr(tok.Literal, index, value)
	}

	return NewIndexExpr(tok.Literal, index)
}
