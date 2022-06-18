package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"compiler/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	return p
}

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
		return false
	}
}

func newProgramNode() *ast.Program {
	return &ast.Program{
		Statements: make([]ast.Statement, 0),
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := newProgramNode()

	for p.curToken.Type != token.EOF {
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
	}
	return nil

}

func (p *Parser) parsetLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

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

	return nil
}

func (p *Parser) parseReturnStatement() ast.Statement {

	return nil
}

func (p *Parser) parseIfStatement() ast.Statement {
	return nil

}
