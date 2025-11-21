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
	MAKE      = "MAKE"
	IDENT     = "IDENT"
	FN        = "FN"
	IF        = "IF"
	ELSE      = "ELSE"
	FOR       = "FOR"
	WHILE     = "WHILE"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	MULT      = "MULT"
	DIV       = "DIV"
	ASSIGN    = "ASSIGN"
	LBRAC     = "LBRAC"
	RBRAC     = "RBRAC"
	LPARAN    = "LPARAN"
	RPARAN    = "RPARAN"
	INT       = "INT"
	INCREMENT = "INCREMENT"
	DECREMENT = "DECREMENT"
	POWER     = "POWER"
	NOT       = "NOT"
	NOTEQUAL  = "NOTEQUAL"
	EQUAL     = "EQUAL"
	LTEQUAL   = "LTEQUAL"
	GTEQUAL   = "GTEQUAL"
	LT        = "LT"
	GT        = "GT"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	EOF       = "EOF"
)

var Keywords = map[string]string{
	"make":  MAKE,
	"fn":    FN,
	"if":    IF,
	"else":  ELSE,
	"for":   FOR,
	"while": WHILE,
	"true":  TRUE,
	"false": FALSE,
}

func (l *Lexer) peakCH() byte {
	if l.nextPos >= len(l.Input) {
		return 0
	}

	return l.Input[l.nextPos]

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

func (l *Lexer) isIdent() bool {
	if (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z') || (l.ch == '_') {
		return true
	}
	return false
}

func (l *Lexer) readIdents() Token {
	var word string
	for (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z') || (l.ch == '_') {
		word = word + string(l.ch)
		l.readChar()
	}

	token, ok := Keywords[word]
	if !ok {
		return Token{Type: IDENT, Literal: word}
	}

	return Token{Type: token, Literal: word}
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

	if l.isIdent() {

		return l.readIdents()
	}

	if l.ch == 0 {
		return Token{Type: EOF, Literal: ""}
	}

	switch l.ch {
	case '+':

		if l.peakCH() == '+' {
			tok = Token{Type: INCREMENT, Literal: "++"}
			l.readChar()
		} else {
			tok = Token{Type: PLUS, Literal: "+"}
		}
	case '-':
		if l.peakCH() == '-' {
			tok = Token{Type: DECREMENT, Literal: "--"}
			l.readChar()
		} else {
			tok = Token{Type: MINUS, Literal: "-"}
		}

	case '*':
		if l.peakCH() == '*' {
			tok = Token{Type: POWER, Literal: "**"}
			l.readChar()
		} else {
			tok = Token{Type: MULT, Literal: "*"}
		}
	case '/':
		tok = Token{Type: DIV, Literal: "/"}
	case '=':
		if l.peakCH() == '=' {
			tok = Token{Type: EQUAL, Literal: "=="}
			l.readChar()
		} else {
			tok = Token{Type: ASSIGN, Literal: "="}
		}
	case '!':
		if l.peakCH() == '=' {
			tok = Token{Type: NOTEQUAL, Literal: "!="}
			l.readChar()
		} else {
			tok = Token{Type: NOT, Literal: "!"}
		}
	case '>':
		if l.peakCH() == '=' {
			tok = Token{Type: GTEQUAL, Literal: ">="}
			l.readChar()
		} else {
			tok = Token{Type: GT, Literal: ">"}
		}
	case '<':
		if l.peakCH() == '=' {
			tok = Token{Type: LTEQUAL, Literal: "<="}
			l.readChar()
		} else {
			tok = Token{Type: LT, Literal: "<"}
		}
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

	// input := "(++)={*2}-832/+!=>!<a>=<= make if"
	l := NewLexer("make x = 10 \n if x == 10 {print(true)}")
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
