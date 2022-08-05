package parser

import (
	"compiler/ast"
	"compiler/token"
	"strconv"
)

type ExprParser struct {
	*Parser
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewExprParser(parser *Parser) *ExprParser {
	p := &ExprParser{
		Parser:         parser,
		prefixParseFns: map[token.TokenType]prefixParseFn{},
		infixParseFns:  map[token.TokenType]infixParseFn{},
	}
	p.registerPrefix(token.IDENT, p.parseIndentifier)
	p.registerPrefix(token.INT, p.parseInteger)
	return p
}

func (p *ExprParser) parsePrefixExpression() ast.Expression {
	// pf := ast.PrefixExpression{
	// 	Token: p.curToken,
	// 	Right: nil,
	// }
	fn := p.prefixParseFns[p.curToken.Type]
	ex := fn()
	precedence := findPrecedence(p.curToken.Type)
	p.parseExpression(precedence)
	return ex

}

func (p *ExprParser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *ExprParser) parseIndentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *ExprParser) parseInteger() ast.Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		p.addPeekError(p.curToken.Type)
	}

	return &ast.IntegerLiteral{
		Token: p.curToken,
		Value: value,
	}
}

func (p *ExprParser) registerPrefix(t token.TokenType, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *ExprParser) registerInfix(t token.TokenType, f infixParseFn) {
	p.infixParseFns[t] = f
}
