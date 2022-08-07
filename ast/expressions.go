package ast

import "compiler/token"

var _ Expression = (*PrefixExpression)(nil)
var _ Expression = (*InfixExpression)(nil)
var _ Expression = (*SuffixExpression)(nil)
var _ Expression = (*Identifier)(nil)
var _ Expression = (*IntegerLiteral)(nil)
var _ Expression = (*Boolean)(nil)
var _ Expression = (*IfExpreesion)(nil)

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
	return "(" + ie.Left.String() + " " + ie.TokenLiteral() + " " + ie.Right.String() + ")"
}

type SuffixExpression struct {
	Token token.Token
	Left  Expression
}

func (p *SuffixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *SuffixExpression) expressionNode() {

}

func (ie *SuffixExpression) String() string {
	return "(" + ie.Left.String() + ie.TokenLiteral() + ")"
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (p *Boolean) TokenLiteral() string {
	return p.Token.Literal
}

func (p *Boolean) expressionNode() {

}

func (ie *Boolean) String() string {
	return "(" + ie.TokenLiteral() + ")"
}

type IfExpreesion struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternatvie *BlockStatement
}

func (expr *IfExpreesion) TokenLiteral() string {
	return expr.Token.Literal
}

func (expr *IfExpreesion) expressionNode() {

}

func (expr *IfExpreesion) String() string {
	s := "if"
	s += " (" + expr.Condition.String() + ")"
	s += expr.Consequence.String() 
	if expr.Alternatvie != nil {
		s += expr.Alternatvie.String()
	}
	return s
}
type FnExpression struct {
	Token token.Token
	Param []Identifier
	Body  BlockStatement
}

func (expr *FnExpression) TokenLiteral() string {
	return expr.Token.Literal
}

func (expr *FnExpression) expressionNode() {

}

func (expr *FnExpression) String() string {
	s := "fn("
	for _, iden := range expr.Param {
		s += iden.TokenLiteral()
		s += ","
	}
	if len(expr.Param) != 0 {
		s = s[:len(s)-1]
	}
	s += ") {"
	s += expr.Body.String()
	s += "}"
	return s
}
