package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
    let y = 10;
    let foobar = 8234141;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	require.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "let", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok)
	assert.Equal(t, name, letStmt.Name.TokenLiteral())
}

func TestReturnStatements(t *testing.T) {
	input := `return x - 5;
    return 5;
    return x;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	require.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"(x - 5)"},
		{"5"},
		{"x"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testReturnStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testReturnStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "return", s.TokenLiteral())

	letStmt, ok := s.(*ast.ReturnStatement)
	assert.True(t, ok)
	assert.Equal(t, name, letStmt.Value.String())
}
