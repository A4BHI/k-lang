package klang

type Parser struct {
	l         *Lexer
	CurrToken Token
	PeekNext  Token
	Errors    []string
}

func (p *Parser) nextToken() {
	p.CurrToken = p.PeekNext
	p.PeekNext = p.l.NextToken()

}

func (p *Parser) parseStatements() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.CurrToken.Type != EOF{
		stmt := 
	}
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}
