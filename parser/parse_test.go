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

func TestParseIntegerExpression(t *testing.T) {
	table := [][2]string{
		{
			"5;", "5",
		},
		{
			"-1;", "-1",
		},
		{
			"0;", "0",
		},
	}

	for _, data := range table {
		l := lexer.New(data[0])
		p := New(l)
		program := p.ParseProgram()
		assert.Equal(t, 1, len(program.Statements))
	}
}
