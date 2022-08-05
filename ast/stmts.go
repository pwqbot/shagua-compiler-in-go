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

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (ls *ReturnStatement) statementNode() {

}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
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
