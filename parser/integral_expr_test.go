package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseIntegerExpression(t *testing.T) {
	table := []struct {
		input string
		value int
	}{
		{
			"5;", 5,
		},
		{
			"0;", 0,
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		assert.True(t, ok)

		intLiteral, ok := expr.Expression.(*ast.IntegerLiteral)
		assert.True(t, ok)

		assert.Equal(t, int64(data.value), intLiteral.Value)
	}
}

func TestParsePrefixExpression(t *testing.T) {
	table := []struct {
		input string
		op    string
		value int
	}{
		{
			"!5;", "!", 5,
		},
		{
			"-1;", "-", 1,
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		assert.True(t, ok)

		prefixExpression, ok := expr.Expression.(*ast.PrefixExpression)
		assert.True(t, ok)

		intLiteral, ok := prefixExpression.Right.(*ast.IntegerLiteral)
		assert.True(t, ok)
		assert.Equal(t, int64(data.value), intLiteral.Value)
		assert.Equal(t, data.op, prefixExpression.TokenLiteral())
	}
}

func TestParseInfixExpression(t *testing.T) {
	table := []struct {
		input  string
		lvalue int
		op     string
		rvalue int
	}{
		{
			"5 + 5;", 5, "+", 5,
		},
		{
			"1 - 1;", 1, "-", 1,
		},
		{
			"22 * 22;", 22, "*", 22,
		},
		{
			"0 / 0;", 0, "/", 0,
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		assert.True(t, ok)

		infixExpression, ok := expr.Expression.(*ast.InfixExpression)
		assert.True(t, ok)

		LIntLiteral, ok := infixExpression.Left.(*ast.IntegerLiteral)
		assert.True(t, ok)
		RIntLiteral, ok := infixExpression.Right.(*ast.IntegerLiteral)
		assert.True(t, ok)

		assert.Equal(t, int64(data.lvalue), LIntLiteral.Value)
		assert.Equal(t, int64(data.rvalue), RIntLiteral.Value)
		assert.Equal(t, data.op, infixExpression.TokenLiteral())
	}
}

func TestPrecedenceOperator(t *testing.T) {
	table := []struct {
		input  string
		expect string
	}{
		{
			"5 + 5;", "(5 + 5)",
		},
		{
			"1 - 1 * 10;", "(1 - (1 * 10))",
		},
		{
			"22 * 22 / 2;", "((22 * 22) / 2)",
		},
		{
			"1 + 0 / 0 - 5 * 0;", "((1 + (0 / 0)) - (5 * 0))",
		},
		{
			"(1 + 2) * 3;", "((1 + 2) * 3)",
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		assert.True(t, ok)

		assert.Equal(t, data.expect, expr.Expression.String())
	}
}
