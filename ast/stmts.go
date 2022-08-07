package ast

import "compiler/token"

var _ Statement = (*LetStatement)(nil)
var _ Statement = (*ReturnStatement)(nil)
var _ Statement = (*ExpressionStatement)(nil)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	return ""
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (ls *ReturnStatement) statementNode() {

}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *ReturnStatement) String() string {
	return ""
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (s *ExpressionStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ExpressionStatement) statementNode() {
}

func (s *ExpressionStatement) String() string {
	return s.Expression.String()
}

type BlockStatement struct {
	Token token.Token // just {
	Statements []Statement
}

func (s *BlockStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *BlockStatement) statementNode() {
}

func (s *BlockStatement) String() string {
	st := "{"
	for _, stmt := range(s.Statements) {
		st += stmt.String()
	}
	st += "}"
	return st
}
