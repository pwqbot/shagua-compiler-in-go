package eval

import (
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testEval(s string) object.Object {
	l := lexer.New(s)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func TestBangExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"!true", false},
		{"!false", true},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testBooleanObject(t, e, tt.expect)
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect interface{}
	}{
		{"if (true) { 10 }", int64(10)},
		// {"if (false) { 10 }", nil},
		// {"if (1) { 10 }", 10},
		// {"if (1 < 2) { 10 }", 10},
		// {"if (1 > 2) { 10 }", nil},
		// {"if (1 < 2) { 10 } else { 20 }", 10},
		// {"if (1 > 2) { 10 } else { 20 }", 20},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		switch expect := tt.expect.(type) {
		case int64:
			testIntegerObject(t, e, expect)
		case bool:
			testBooleanObject(t, e, expect)
		case nil:
			testNullObject(t, e)
		default:
			glog.Fatal("type not found")
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) {
	assert.Equal(t, obj, object.NULL)
}

func TestInfixBooleanExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"5 == 5", true},
		{"-6 > 1", false},
		{"6 > 1", true},
		{"-6 < -1", true},
		{"6 < -1", false},
		{"-1 <= -1", true},
		{"1 <= -1", false},
		{"-1 >= -1", true},
		{"-1 >= 1", false},
		{"true == true", true},
		{"false == false", true},
		{"(1 > 5) == false", true},
		{"(1 < 5) == true", true},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testBooleanObject(t, e, tt.expect)
	}
}

func TestInfixMathematicExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{"5 + 6", 11},
		{"6 - 1", 5},
		{"-6 -1", -7},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testIntegerObject(t, e, tt.expect)
	}
}

func TestPrefixMinuxExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
		{"-5", -5},
		{"-6", -6},
	}

	for _, tt := range tests {
		e := testEval(tt.input)
		testIntegerObject(t, e, tt.expect)
	}
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect int64
	}{
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
		input  string
		expect bool
	}{
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

	assert.Equal(t, valueExpect, r.Value)
}

func testIntegerObject(t *testing.T, obj object.Object, valueExpect int64) {
	r, ok := obj.(*object.Integer)
	require.True(t, ok)

	assert.Equal(t, valueExpect, r.Value)
}

func TestLetStatements(t *testing.T) {
}

func TestReturnStatements(t *testing.T) {

}
