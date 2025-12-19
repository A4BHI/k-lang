package klang

type Parser struct {
	l         *Lexer
	currToken Token
	nextToken Token
}
