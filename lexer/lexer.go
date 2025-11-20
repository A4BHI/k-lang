package main

import (
	"fmt"
)

type Token struct {
	Type    string
	Literal string
}

type Lexer struct {
	Input      string
	currentPos int
	nextPos    int
	ch         byte
}

const (
	MAKE   = "MAKE"
	IDENT  = "IDENT"
	FN     = "FN"
	IF     = "IF"
	ELSE   = "ELSE"
	FOR    = "FOR"
	WHILE  = "WHILE"
	PLUS   = "PLUS"
	MINUS  = "MINUS"
	MULT   = "MULT"
	DIV    = "DIV"
	ASSIGN = "ASSIGN"
	LBRAC  = "LBRAC"
	RBRAC  = "RBRAC"
	LPARAN = "LPARAN"
	RPARAN = "RPARAN"
	INT    = "INT"
	EOF    = "EOF"
)

var Keywords = map[string]string{
	"make":  MAKE,
	"fn":    FN,
	"if":    IF,
	"else":  ELSE,
	"for":   FOR,
	"while": WHILE,
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

func (l *Lexer) isDigit() bool {
	if l.ch >= '0' && l.ch <= '9' {
		return true
	}
	return false
}

func (l *Lexer) readDigits() Token {
	var Digit string
	for l.ch >= '0' && l.ch <= '9' {
		Digit = Digit + string(l.ch)
		l.readChar()
	}
	return Token{Type: INT, Literal: Digit}

}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.Input) {

		l.ch = 0
	} else {
		l.ch = l.Input[l.nextPos]
	}

	l.currentPos = l.nextPos
	l.nextPos++

}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhiteSpace()

	if l.isDigit() { //To Read DIGITS

		return l.readDigits()
	}

	if l.ch == 0 {
		return Token{Type: EOF, Literal: ""}
	}

	switch l.ch {
	case '+':
		tok = Token{Type: PLUS, Literal: "+"}
	case '-':
		tok = Token{Type: MINUS, Literal: "-"}
	case '*':
		tok = Token{Type: MULT, Literal: "*"}
	case '/':
		tok = Token{Type: DIV, Literal: "/"}
	case '=':
		tok = Token{Type: ASSIGN, Literal: "="}
	case '(':
		tok = Token{Type: LBRAC, Literal: "("}
	case ')':
		tok = Token{Type: RBRAC, Literal: ")"}
	case '{':
		tok = Token{Type: LPARAN, Literal: "{"}
	case '}':
		tok = Token{Type: RPARAN, Literal: "}"}
	default:

	}

	l.readChar()

	return tok
}

func NewLexer(input string) *Lexer {
	l := Lexer{Input: input}
	l.readChar()
	return &l
}

func main() {
	input := "(+)={*2}-832/"
	l := NewLexer(input)
	for {

		tok := l.NextToken()
		if tok.Type != EOF {
			fmt.Printf("|%-10s -> |%q|\n", tok.Type, tok.Literal)
		} else {
			fmt.Printf("|%-10s -> |%q |\n", tok.Type, tok.Literal)
			break
		}

	}

}
