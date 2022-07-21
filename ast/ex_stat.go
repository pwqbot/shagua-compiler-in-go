package ast

import "compiler/token"

var _ Expression = (*ExpressionStatement)(nil)

type ExpressionStatement struct {
	t token.Token
}

func (s *ExpressionStatement) TokenLiteral() string {
	return s.t.Literal
}

func (s *ExpressionStatement) expressionNode() {
}
