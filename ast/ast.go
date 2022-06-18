package ast

import "compiler/token"

// node in ast
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	ss := ""
	for _, s := range p.Statements {
		ss += s.TokenLiteral() + "\n"
	}
	return ss
}

// TODO: Distinguish left and right
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) expressionNode() {

}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
