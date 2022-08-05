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

	stmtParser     *StmtParser
	exprParser     *ExprParser
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:          l,
		errors:     []error{},
		curToken:   token.Token{},
		peekToken:  token.Token{},
		stmtParser: &StmtParser{},
		exprParser: &ExprParser{},
	}
	p.stmtParser = &StmtParser{
		Parser: p,
	}
	p.exprParser = NewExprParser(p)

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
		return p.stmtParser.parsetLetStatement()
	case token.IF:
		return p.stmtParser.parseIfStatement()
	case token.RETURN:
		return p.stmtParser.parseReturnStatement()
	case token.FUNCTION:
		return p.stmtParser.parseFunctionStatement()
	default:
		return p.stmtParser.parseExpressionStatement(LOWEST)
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

// error handling
func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) addPeekError(t token.TokenType) {
	err := fmt.Errorf("Expect %v, got %v", t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}
