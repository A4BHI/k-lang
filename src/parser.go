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

const (
	LOWEST = iota
	SUM
	PRODUCT
	PREFIX
)

var precedence = map[string]int{
	PLUS:  SUM,
	MINUS: SUM,
	MULT:  PRODUCT,
	DIV:   PRODUCT,
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
	var l Expression
	switch p.CurrToken.Type {
	case INT:
		val, err := strconv.Atoi(p.CurrToken.Literal)
		if err != nil {
			p.Errors = append(p.Errors, "Could not parse integer")
			return nil
		}
		l = &IntegerLiteral{Value: val}

	case IDENT:
		if p.PeekNext.Type == LBRAC {
			l = p.parseFunctionCall()

		} else {
			l = &Identifier{Value: p.CurrToken.Literal}

		}

	case MINUS, NOT:
		l = p.parsePrefixExpression()

	default:
		return nil

	}

	if p.PeekNext.Type == SEMICOLON {
		return l
	}
	var infix InfixExpression
	for p.PeekNext.Type != SEMICOLON {
		if p.currPrecedence() < p.peekPrecedence() {
			infix.Left = l
			p.nextToken()
			infix.Operator = p.CurrToken.Type

		}
	}

	return l

}

func (p *Parser) peekPrecedence() int {
	return precedence[p.PeekNext.Type]
}

func (p *Parser) currPrecedence() int {
	return precedence[p.CurrToken.Type]
}

func (p *Parser) parseFunctionCall() Expression {
	call := &FunctionCall{}

	call.Function = &Identifier{Value: p.CurrToken.Literal}
	p.nextToken()
	p.nextToken()

	call.Arguments = []Expression{}

	if p.PeekNext.Type != RBRAC {
		args := p.parseExpression()
		call.Arguments = append(call.Arguments, args)

		for p.PeekNext.Type == COMMA {
			p.nextToken()
			p.nextToken()

			args := p.parseExpression()
			call.Arguments = append(call.Arguments, args)
		}

	}

	if p.PeekNext.Type == RBRAC {
		p.nextToken()
	}
	return call

}

func (p *Parser) parsePrefixExpression() Expression {
	pre := &PrefixExpression{}
	pre.Operator = p.CurrToken.Literal
	p.nextToken()
	pre.Expression = p.parseExpression()
	return pre

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
