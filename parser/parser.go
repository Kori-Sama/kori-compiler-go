package parser

import (
	"fmt"

	"github.com/Kori-Sama/kori-compiler/cerr"
	"github.com/Kori-Sama/kori-compiler/lexer"
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

	fn.Proto.Name = "TopLevel"

	return fn
}

func (p *Parser) parseTopLevelExpr() *FunctionAST {
	expr := p.parseExpr()
	if expr == nil {
		return nil
	}

	proto := NewPrototypeAST("", nil)
	return NewFunctionAST(proto, expr)
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
			continue
			// fn = p.HandleTopLevel()
			// res = append(res, fn)
		}
	}
out:
	return res
}
