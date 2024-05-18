package parser

import (
	"fmt"
	"strconv"

	"github.com/Kori-Sama/compiler-go/cerr"
	"github.com/Kori-Sama/compiler-go/lexer"
)

type Parser struct {
	tokens []*lexer.Token
	curTok int
	Err    *cerr.ParserError
	token  *lexer.Token
}

func (p *Parser) Error(message string) {
	cerr.NewParserError(message, p.token.Line, p.token.Location)
}

func (p *Parser) Expect(what, where string) {
	cerr.NewParserError(fmt.Sprintf("Expect '%s' int %s", what, where), p.token.Line, p.token.Location)
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
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
	case lexer.TOKEN_LET:
		return p.parseInitExpr()
	case lexer.TOKEN_SEMI:
		return nil
	case lexer.TOKEN_EOF:
		return nil
	default:
		p.Err = cerr.NewParserError(fmt.Sprintf("Unknown token '%s' when expecting an expression", tok.Literal), tok.Line, tok.Location)
		return nil
	}
}

func (p *Parser) nextToken() {
	if p.curTok >= len(p.tokens) || p.tokens[p.curTok].Kind == lexer.TOKEN_EOF {
		return
	}
	p.curTok++
	p.token = p.tokens[p.curTok]
}

func (p *Parser) getCurTok() *lexer.Token {
	if p.curTok >= len(p.tokens) {
		return nil
	}
	return p.tokens[p.curTok]
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
	for p.getCurTok().Kind != lexer.TOKEN_RPAREN {
		arg := p.parseExpr()
		if arg == nil {
			return nil
		}
		args = append(args, arg)

		if p.getCurTok().Kind != lexer.TOKEN_COMMA {
			p.Err = cerr.NewParserError("Expected ',' or ')' in arguments", tok.Line, tok.Location)
			return nil
		}

		p.nextToken()
	}

	p.nextToken()
	return NewCallExpr(tok.Literal, args)
}

var binOpPrecedence = map[OpKind]int{
	OP_LESS:       10,
	OP_GREATER:    10,
	OP_LESS_EQ:    10,
	OP_GREATER_EQ: 10,
	OP_EQ:         10,
	OP_ADD:        20,
	OP_SUB:        20,
	OP_MUL:        40,
	OP_DIV:        40,
}

func (p *Parser) parseExpr() Expr {
	lhs := p.parsePrimary()
	if lhs == nil {
		return nil
	}

	return p.parseBinOpRHS(0, lhs)
}

func (p *Parser) parseBinOpRHS(exprPrec int, lhs Expr) Expr {
	for {
		tok := p.getCurTok()
		if tok == nil {
			return lhs
		}

		if tok.Kind != lexer.TOKEN_LESS &&
			tok.Kind != lexer.TOKEN_GREATER &&
			tok.Kind != lexer.TOKEN_LESS_EQ &&
			tok.Kind != lexer.TOKEN_GREATER_EQ &&
			tok.Kind != lexer.TOKEN_EQ &&
			tok.Kind != lexer.TOKEN_PLUS &&
			tok.Kind != lexer.TOKEN_MINUS &&
			tok.Kind != lexer.TOKEN_STAR &&
			tok.Kind != lexer.TOKEN_SLASH {
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

	expr := p.parseExpr()
	if expr == nil {
		return nil
	}

	tok := p.getCurTok()
	if tok.Kind != lexer.TOKEN_RBRACE {
		p.Err = cerr.NewParserError("Expected '}'", tok.Line, tok.Location)
		return nil
	}

	p.nextToken()

	return expr
}

func (p *Parser) parseTopLevelExpr() *FunctionAST {
	expr := p.parseExpr()
	if expr == nil {
		return nil
	}

	proto := NewPrototypeAST("", nil)
	return NewFunctionAST(proto, expr)
}

func (p *Parser) HandleFunction() *FunctionAST {
	fn := p.parseFunction()
	if fn == nil {
		return nil
	}
	return fn
}

func (p *Parser) HandleTopLevel() *FunctionAST {
	fn := p.parseTopLevelExpr()
	if fn == nil {
		return nil
	}
	p.nextToken()
	return fn
}

func (p *Parser) Parse() []*FunctionAST {
	var res []*FunctionAST
	for {
		var fn *FunctionAST
		tok := p.getCurTok()
		if tok == nil {
			goto out
		}

		if p.Err != nil {
			goto out
		}
		switch tok.Kind {
		case lexer.TOKEN_EOF:
			goto out
		case lexer.TOKEN_SEMI:
			p.nextToken()
		case lexer.TOKEN_FUNC:
			fn = p.HandleFunction()
			res = append(res, fn)
		default:
			fn = p.HandleTopLevel()
			res = append(res, fn)
		}
	}
out:
	return res
}
