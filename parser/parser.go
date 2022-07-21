package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"compiler/token"
	"fmt"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// take token from lexer, then parse to ast
type Parser struct {
	l      *lexer.Lexer
	errors []error

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		errors:         []error{},
		curToken:       token.Token{},
		peekToken:      token.Token{},
		prefixParseFns: map[token.TokenType]prefixParseFn{},
		infixParseFns:  map[token.TokenType]infixParseFn{},
	}
	p.registerPrefix(token.IDENT, p.parseIndentifier)
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: make([]ast.Statement, 0),
	}

	for p.curToken.Type != token.EOF {
		// progrom made up of statements
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parsetLetStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.FUNCTION:
		return p.parseFunctionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	return nil
}

func (p *Parser) parsetLetStatement() ast.Statement {
	// let <identifier> = <expression>
	stmt := &ast.LetStatement{
		Token: p.curToken,
		Name:  nil,
		Value: nil,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// do noting with assign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO(dingwang): parse expression
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return nil
}

func (p *Parser) parseReturnStatement() ast.Statement {
	// return <expression>
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
		Value: nil,
	}
	p.nextToken()
	// TODO(dingwang): parse expression
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseIfStatement() ast.Statement {
	return nil
}

func (p *Parser) parseFunctionStatement() ast.Statement {
	return nil
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseIndentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
} 

// some helper junctions
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addPeekError(t)
		return false
	}
}

func (p *Parser) registerPrefix(t token.TokenType, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *Parser) registerInfix(t token.TokenType, f infixParseFn) {
	p.infixParseFns[t] = f
}

// error handling
func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) addPeekError(t token.TokenType) {
	err := fmt.Errorf("Expect %v, got %v", t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}
