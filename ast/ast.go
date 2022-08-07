package ast

// node in ast
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
	String() string
}

type Expression interface {
	Node
	String() string
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
