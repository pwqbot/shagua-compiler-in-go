package ast

import "compiler/token"

var _ Expression = (*Identifier)(nil)
var _ Expression = (*IntegerLiteral)(nil)
var _ Expression = (*PrefixExpression)(nil)
var _ Expression = (*InfixExpression)(nil)

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

func (i *Identifier) String() string {
	return i.TokenLiteral()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) String() string {
	return i.TokenLiteral()
}

type PrefixExpression struct {
	Token token.Token
	Right Expression
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) expressionNode() {

}

func (ie *PrefixExpression) String() string {
	return "(" + ie.TokenLiteral() + ie.Right.String() + ")"
}

type InfixExpression struct {
	Token token.Token
	Left  Expression
	Right Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) expressionNode() {

}

func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " "+ ie.TokenLiteral() + " " +  ie.Right.String() + ")"
}
