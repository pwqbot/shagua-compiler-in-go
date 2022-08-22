package parser

import (
	"compiler/ast"
	"compiler/token"
	"strconv"

	"github.com/golang/glog"
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

	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.IDENT, p.parseIndentifier)
	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseParem)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MULTI, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GE, p.parseInfixExpression)
	p.registerInfix(token.LE, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)

	// NOTE: treat suffix as infix without right expr
	p.registerInfix(token.PLUSPLUS, p.parseSuffixExpression)
	p.registerInfix(token.MINUSMINUS, p.parseSuffixExpression)
	p.registerInfix(token.WHAT, p.parseSuffixExpression)
	return p
}

func (p *ExprParser) ParseExpreesion(precedence int) ast.Expression {
	// NOTE: check if we have a prefixFn associated with curToken
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) &&
		precedence < findPrecedence(p.peekToken.Type) {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()

		leftExp = infix(leftExp)
	}

	glog.Error(leftExp.TokenLiteral())
	return leftExp
}

func (p *ExprParser) registerPrefix(t token.TokenType, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *ExprParser) registerInfix(t token.TokenType, f infixParseFn) {
	p.infixParseFns[t] = f
}

func (p *ExprParser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token: p.curToken,
		Right: nil,
	}

	p.nextToken()
	precedence := findPrecedence(p.curToken.Type)
	expr.Right = p.ParseExpreesion(precedence)

	return expr
}

// <Expr> <op> <Expr>
func (p *ExprParser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token: p.curToken,
		Left:  left,
		Right: nil,
	}

	precedence := findPrecedence(p.curToken.Type)
	p.nextToken()
	expr.Right = p.ParseExpreesion(precedence)

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
	expr := p.ParseExpreesion(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
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

func (p *ExprParser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

func (p *ExprParser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpreesion{
		Token:       p.curToken,
		Condition:   nil,
		Consequence: &ast.BlockStatement{},
		Alternatvie: nil,
	}
	p.expectPeek(token.LPAREN)
	p.nextToken()
	expr.Condition = p.ParseExpreesion(LOWEST)
	p.expectPeek(token.RPAREN)

	p.expectPeek(token.LBRACE)
	expr.Consequence = p.stmtParser.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		p.expectPeek(token.LBRACE)
		expr.Alternatvie = p.stmtParser.parseBlockStatement()
	}

	return expr
}
