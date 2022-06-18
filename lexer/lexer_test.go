package lexer

import (
	"compiler/token"
	"testing"
)

func TestNextToken_BasicToken(t *testing.T) {
	// Arrange
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Type,
				tt.expectedType)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Literal,
				tt.expectedLiteral)
		}
	}
}

func TestNextToken_Identifier(t *testing.T) {
	// Arrange
	input := `let abc = 512
        let add = fn(x, y) {
            x + y
        }
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "abc"},
		{token.ASSIGN, "="},
		{token.INT, "512"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Type,
				tt.expectedType)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Literal,
				tt.expectedLiteral)
		}
	}
}

func TestNextToken_Peek(t *testing.T) {
	// Arrange
	input := `a == b
    a != b
    a <= b
    a>=b`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "a"},
		{token.EQ, "=="},
		{token.IDENT, "b"},
		{token.IDENT, "a"},
		{token.NE, "!="},
		{token.IDENT, "b"},
		{token.IDENT, "a"},
		{token.LE, "<="},
		{token.IDENT, "b"},
		{token.IDENT, "a"},
		{token.GE, ">="},
		{token.IDENT, "b"},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Type,
				tt.expectedType)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test %d error, got %s, expect %s", i, token.Literal,
				tt.expectedLiteral)
		}
	}
}
