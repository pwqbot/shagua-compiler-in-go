package eval

import (
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testEval(s string) object.Object{
	l := lexer.New(s)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input string
		expect int64
	} {
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testIntegerObject(t, e, tt.expect)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input string
		expect bool
	} {
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testBooleanObject(t, e, tt.expect)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, valueExpect bool) {
	r, ok := obj.(*object.Boolean)
	require.True(t, ok)

	assert.Equal(t, r.Value, valueExpect)
}

func testIntegerObject(t *testing.T, obj object.Object, valueExpect int64) {
	r, ok := obj.(*object.Integer)
	require.True(t, ok)

	assert.Equal(t, r.Value, valueExpect)
}

func TestLetStatements(t *testing.T) {
}

func TestReturnStatements(t *testing.T) {

}
