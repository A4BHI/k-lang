package main

import "fmt"

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
	IF     = "IF"
	ELSE   = "ELSE"
	FOR    = "FOR"
	WHILE  = "WHILE"
	PLUS   = "PLUS"
	MINUS  = "MINUS"
	MULT   = "MULT"
	DIV    = "DIV"
	LBRAC  = "LBRAC"
	RBRAC  = "RBRAC"
	LPARAN = "LPARAN"
	RPARAN = "RPARAN"
	INT    = "INT"
	EOF    = "EOF"
)

var Keywords = map[string]string{
	"make":  MAKE,
	"if":    IF,
	"else":  ELSE,
	"for":   FOR,
	"while": WHILE,
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
	switch l.ch {
	case '+':
		tok = Token{Type: PLUS, Literal: "+"}
	case '-':
		tok = Token{Type: MINUS, Literal: "-"}
	case '*':
		tok = Token{Type: MULT, Literal: "*"}
	case '/':
		tok = Token{Type: DIV, Literal: "/"}
	default:
		if l.ch == 0 {
			tok = Token{Type: EOF, Literal: ""}
		}

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
	input := "+*-/"
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
