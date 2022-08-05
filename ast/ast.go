package ast

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
