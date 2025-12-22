package klang

import "strconv"

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
	switch p.CurrToken.Type {
	case MAKE:
		return p.parseMakeStatement()
	default:
		return nil
	}
}

func (p *Parser) parseMakeStatement() *MakeStatement {
	stmt := &MakeStatement{}

	p.nextToken()
	if p.CurrToken.Type != IDENT {
		p.Errors = append(p.Errors, "Expected Identifier But Got None")
		return nil
	}

	stmt.Name = &Identifier{Value: p.CurrToken.Literal}

	p.nextToken()
	if p.CurrToken.Type != ASSIGN {
		p.Errors = append(p.Errors, "Expected = after identifier")
		return nil
	}

	p.nextToken()
	switch p.CurrToken.Type {
	case INT:
		v, err := strconv.Atoi(p.CurrToken.Literal)
		if err != nil {
			p.Errors = append(p.Errors, "Error cant convert INT")
			return nil
		}
		stmt.Value = &IntegerLiteral{Value: v}

	default:
		p.Errors = append(p.Errors, "expected expression after =")
		return nil
	}
	if p.PeekNext.Type == SEMICOLON {
		p.nextToken()
	}

	return stmt
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
