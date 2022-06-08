package lexer

import (
	"compiler/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int // point to next position to be read
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipDelim()
	switch l.ch {
	case '=':
		if l.peekRune(1) == "=" {
			l.readRune()
			tok = token.Token{
				Type:    token.EQ,
				Literal: "==",
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekRune(1) == "=" {
			l.readRune()
			tok = token.Token{
				Type:    token.NE,
				Literal: "!=",
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '?':
		tok = newToken(token.WHAT, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '>':
		if l.peekRune(1) == "=" {
			l.readRune()
			tok = token.Token{
				Type:    token.GE,
				Literal: ">=",
			}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '<':
		if l.peekRune(1) == "=" {
			l.readRune()
			tok = token.Token{
				Type:    token.LE,
				Literal: "<=",
			}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '*':
		tok = newToken(token.MULTI, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LoopUpKeywords(tok.Literal)
			return tok
		} else if l.isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readRune()
	return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) peekRune(x int) string {
	if l.readPosition+x-1 >= len(l.input) {
		return ""
	} else {
		return string(l.input[l.readPosition : l.readPosition+x])
	}
}

func (l *Lexer) readIdentifier() string {
	beginPosition := l.position
	for l.isLetter(l.ch) {
		l.readRune()
	}
	return string(l.input[beginPosition:l.position])
}

func (l *Lexer) readNumber() string {
	beginPosition := l.position
	for l.isDigit(l.ch) {
		l.readRune()
	}
	return string(l.input[beginPosition:l.position])
}

func (l *Lexer) isLetter(ch rune) bool {
	return ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch == '_' ||
		ch == '!' ||
		ch == '?'
}

func (l *Lexer) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipDelim() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readRune()
	}
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
