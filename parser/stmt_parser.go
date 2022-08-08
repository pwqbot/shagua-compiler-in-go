package parser

import (
	"compiler/ast"
	"compiler/token"
)

type StmtParser struct {
	*Parser
}

func (p *StmtParser) parseExpressionStatement(precedence int) ast.Statement {
	stmt := &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: p.exprParser.ParseExpreesion(LOWEST),
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *StmtParser) parsetLetStatement() ast.Statement {
	// let <identifier> = <expression>
	stmt := &ast.LetStatement{
		Token: p.curToken,
		Name:  nil,
		Value: nil,
	}

	print(p.peekToken.Literal)
	if !p.expectPeek(token.IDENT) {
		p.addPeekError(p.peekToken.Type)
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
	stmt.Value = p.exprParser.ParseExpreesion(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *StmtParser) parseReturnStatement() ast.Statement {
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

func (p *StmtParser) parseIfStatement() ast.Statement {
	return nil
}

func (p *StmtParser) parseFunctionStatement() ast.Statement {
	return nil
}

func (p *StmtParser) parseBlockStatement() *ast.BlockStatement {
	b := &ast.BlockStatement{
		Token:      p.curToken,
		Statements: []ast.Statement{},
	}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			b.Statements = append(b.Statements, stmt)
		}
		p.nextToken()
	}

	return b
}
