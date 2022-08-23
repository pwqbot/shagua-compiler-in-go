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
		require.True(t, ok)

		intLiteral, ok := expr.Expression.(*ast.IntegerLiteral)
		require.True(t, ok)

		assert.Equal(t, int64(data.value), intLiteral.Value)
	}
}

func TestParseBooleanExpression(t *testing.T) {
	table := []struct {
		input string
		value bool
	}{
		{
			"true;", true,
		},
		{
			"false", false,
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		require.True(t, ok)

		intLiteral, ok := expr.Expression.(*ast.Boolean)
		require.True(t, ok)

		assert.Equal(t, data.value, intLiteral.Value)
	}
}

func TestIfExpreesion(t *testing.T) {
	table := []struct {
		input  string
		expect string
	}{
		{
			"if (x < y) { x }", "if ((x < y)){x}",
		},
		{
			"if (x < y) { x } else { y }", "if ((x < y)){x}{y}",
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, []error{}, p.Errors())
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		require.True(t, ok)

		intLiteral, ok := expr.Expression.(*ast.IfExpreesion)
		require.True(t, ok)

		assert.Equal(t, data.expect, intLiteral.String())
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
		require.True(t, ok)

		prefixExpression, ok := expr.Expression.(*ast.PrefixExpression)
		require.True(t, ok)

		intLiteral, ok := prefixExpression.Right.(*ast.IntegerLiteral)
		require.True(t, ok)
		assert.Equal(t, int64(data.value), intLiteral.Value)
		assert.Equal(t, data.op, prefixExpression.TokenLiteral())
	}
}

func TestParseSuffixExpression(t *testing.T) {
	table := []struct {
		input  string
		expect string
	}{
		{
			"5++;", "(5++)",
		},
		{
			"1--;", "(1--)",
		},
	}

	for _, data := range table {
		// println(data)
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		require.True(t, ok)

		suffixExpression, ok := expr.Expression.(*ast.SuffixExpression)
		require.True(t, ok)
		assert.Equal(t, data.expect, suffixExpression.String())
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
		l := lexer.New(data.input)
		p := New(l)
		program := p.ParseProgram()
		require.Equal(t, 1, len(program.Statements))

		expr, ok := (program.Statements[0]).(*ast.ExpressionStatement)
		assert.True(t, ok)

		infixExpression, ok := expr.Expression.(*ast.InfixExpression)
		assert.True(t, ok, expr.Expression.String())

		LIntLiteral, ok := infixExpression.Left.(*ast.IntegerLiteral)
		assert.True(t, ok, infixExpression.Left.String())
		RIntLiteral, ok := infixExpression.Right.(*ast.IntegerLiteral)
		assert.True(t, ok, infixExpression.Right.String())

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
		{
			"1-- + 2++;", "((1--) + (2++))",
		},
		{
			"1 + 2 < 3;", "((1 + 2) < 3)",
		},
		{
			"1 + 2 <= 3;", "((1 + 2) <= 3)",
		},
		{
			"1 + 2 >= 3;", "((1 + 2) >= 3)",
		},
		{
			"1 + 2 > 3;", "((1 + 2) > 3)",
		},
		{
			"1 + (2 > 3);", "(1 + (2 > 3))",
		},
		{
			"5 > 4 == 3 < 4;", "((5 > 4) == (3 < 4))",
		},
		// {
		// 	"1 +-+ 2;", "",
		// },
		{
			"(1-- + 2 * 3) + 2++;", "(((1--) + (2 * 3)) + (2++))",
		},
		{
			"-5 + 5", "((-5) + 5)",
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
