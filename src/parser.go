package klang

type Parser struct {
	l         *Lexer
	currToken Token
	nextToken Token
	Errors    []string
}

func (p *Parser) nexttoken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()

}
