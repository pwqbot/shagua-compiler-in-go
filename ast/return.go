package ast

import "compiler/token"

var _ Statement = (*ReturnStatement)(nil)

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (ls *ReturnStatement) statementNode() {

}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}
