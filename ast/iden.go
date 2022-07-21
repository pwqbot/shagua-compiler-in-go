package ast

import "compiler/token"

var _ Expression = (*Identifier)(nil)

// TODO(dingwang): Distinguish left and right
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
