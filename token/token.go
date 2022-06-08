package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	INDENT TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"

	// Operators
	ASSIGN TokenType = "="
	EQ     TokenType = "=="
	NE     TokenType = "!="
	PLUS   TokenType = "+"
	MINUS  TokenType = "-"
	MULTI  TokenType = "*"
	DIVIDE TokenType = "/"
	BANG   TokenType = "!"
	WHAT   TokenType = "?"
	LT     TokenType = "<"
	LE     TokenType = "<="
	GT     TokenType = ">"
	GE     TokenType = ">="

	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LoopUpKeywords(key string) TokenType {
	if tok, ok := keywords[key]; ok {
		return tok
	}
	return INDENT
}
