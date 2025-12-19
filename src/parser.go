package klang

import "fmt"

type Parser struct {
	l         *Lexer
	CurrToken Token
	NextToken Token
	Errors    []string
}

func (p *Parser) nextToken() {
	p.CurrToken = p.NextToken
	p.NextToken = p.l.NextToken()

}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	p.nextToken()
	p.nextToken()
	fmt.Println(p)
	return p
}
