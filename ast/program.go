package ast

var _ Node = (*Program)(nil)

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
