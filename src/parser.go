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

func (p *Parser) parseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.CurrToken.Type != EOF {

		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()

	}

	return program
}

func (p *Parser) parseStatement() Statement {
	return nil
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
