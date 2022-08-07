package parser

import (
	"compiler/ast"
	"compiler/token"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
	suffixParseFn func(ast.Expression) ast.Expression
)

type ExprParser struct {
	*Parser
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	suffixParseFns map[token.TokenType]suffixParseFn
}

func NewExprParser(parser *Parser) *ExprParser {
	p := &ExprParser{
		Parser:         parser,
		prefixParseFns: map[token.TokenType]prefixParseFn{},
		infixParseFns:  map[token.TokenType]infixParseFn{},
		suffixParseFns: map[token.TokenType]suffixParseFn{},
	}

	p.registerPrefix(token.IDENT, p.parseIndentifier)
	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseParem)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MULTI, p.parseInfixExpression)

	p.registerSuffix(token.PLUSPLUS, p.parseSuffixExpression)
	p.registerSuffix(token.MINUSMINUS, p.parseSuffixExpression)
	return p
}

func (p *ExprParser) registerPrefix(t token.TokenType, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *ExprParser) registerInfix(t token.TokenType, f infixParseFn) {
	p.infixParseFns[t] = f
}

func (p *ExprParser) registerSuffix(t token.TokenType, f suffixParseFn) {
	p.suffixParseFns[t] = f
}

func (p *ExprParser) parseExpression(precedence int) ast.Expression {
	// NOTE: check if we have a prefixFn associated with curToken
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	suffix := p.suffixParseFns[p.peekToken.Type]
	if suffix != nil {
		println("find")
		p.nextToken()
		leftExp = suffix(leftExp)
	}

	for !p.peekTokenIs(token.SEMICOLON) &&
		precedence < findPrecedence(p.peekToken.Type) {
		p.nextToken()
		infix := p.infixParseFns[p.curToken.Type]
		// println(p.curToken.Type)
		if infix == nil {
			return nil
		}
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *ExprParser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token: p.curToken,
		Right: nil,
	}

	p.nextToken()
	precedence := findPrecedence(p.curToken.Type)
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *ExprParser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token: p.curToken,
		Left:  left,
		Right: nil,
	}

	precedence := findPrecedence(p.curToken.Type)
	p.nextToken()
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *ExprParser) parseSuffixExpression(left ast.Expression) ast.Expression {
	expr := &ast.SuffixExpression{
		Token: p.curToken,
		Left:  left,
	}
	return expr
}

func (p *ExprParser) parseIndentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *ExprParser) parseParem() ast.Expression {
	p.nextToken()
	expr := p.parseExpression(LOWEST)
	if !p.peekTokenIs(token.RPAREN) {
		return nil
	}
	p.nextToken()
	return expr
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
