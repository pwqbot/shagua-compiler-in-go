package parser

import "compiler/token"

const (
	_ int = iota
	LOWEST
	EQUAL
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	NOT
	WEN
	LPAREN
)

var precedencs = map[token.TokenType]int{
	// Operators
	token.ASSIGN:    EQUAL,
	token.EQ:        EQUAL,
	token.NE:        EQUAL,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.MULTI:     PRODUCT,
	token.DIVIDE:    PRODUCT,
	token.BANG:      NOT,
	token.WHAT:      WEN,
	token.LT:        LESSGREATER,
	token.LE:        LESSGREATER,
	token.GT:        LESSGREATER,
	token.GE:        LESSGREATER,
	token.LPAREN:    LPAREN,
	token.RPAREN:    LOWEST,
	token.LBRACE:    LPAREN,
	token.RBRACE:    LOWEST,
	token.COMMA:     LPAREN,
	token.SEMICOLON: LOWEST,
}

func findPrecedence(t token.TokenType) int {
	if precedence, ok := precedencs[t]; ok {
		return precedence
	} else {
		return LOWEST
	}
}
