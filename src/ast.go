package klang

type Identifier struct {
	Value string
}

func (ident *Identifier) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

type Node interface {
	tokLiteral() string
}

type Statements interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}
type CallExpression struct {
	Function   Expression
	Parameters Expression
}

type KLANG struct {
	Statements []Statements
}

type MakeStatement struct {
	Name  *Identifier
	Value Expression
}

func (m MakeStatement) statementNode() {}
