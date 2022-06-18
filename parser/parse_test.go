package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10; 
    ley foobar = 8234141;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

    require.Equal(t, len(program.Statements), 3)

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
	assert.Equal(t, s.TokenLiteral(), "let")
	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok)
    assert.Equal(t, letStmt.TokenLiteral(), name)
}
