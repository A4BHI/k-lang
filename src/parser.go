package klang

import (
	"strconv"
)

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

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.CurrToken.Type != EOF {

		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()

	}

	return program
}

func (p *Parser) ParseStatement() Statement {
	switch p.CurrToken.Type {
	case MAKE:
		return p.parseMakeStatement()
	default:
		return p.parseExpressionStatement()
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
	expr := p.parseExpression()
	if expr == nil {
		p.Errors = append(p.Errors, "Unknown Expression")
		return nil
	}

	stmt.Value = expr
	if p.PeekNext.Type == SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{}
	stmt.Expression = p.parseExpression()
	if p.PeekNext.Type == SEMICOLON {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression() Expression {
	switch p.CurrToken.Type {
	case INT:
		val, err := strconv.Atoi(p.CurrToken.Literal)
		if err != nil {
			p.Errors = append(p.Errors, "Could not parse integer")
			return nil
		}

		return &IntegerLiteral{Value: val}
	case IDENT:
		if p.PeekNext.Type == LBRAC {
			return p.parseFunctionCall()

		}
		return &Identifier{Value: p.CurrToken.Literal}

	default:
		return nil

	}

}

func (p *Parser) parseFunctionCall() Expression {

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
