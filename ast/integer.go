package ast

import "compiler/token"

var _ Expression = (*IntegerLiteral)(nil)

type IntegerLiteral struct {
	token token.Token
}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.token.Literal
}

func (i *IntegerLiteral) expressionNode() {

}
