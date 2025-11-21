package klang

type Identifier struct {
	Value string
}

func (ident *Identifier) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

func (in IntegerLiteral) expressionNode() {}

type Node interface {
	tokLiteral() string
}

type Statement interface {
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

type InfinixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

type KLANG struct {
	Statements []Statement
}

type MakeStatement struct {
	Name  *Identifier
	Value Expression
}

func (m MakeStatement) statementNode() {}

type IfStatement struct {
	Condition      Expression
	TrueStatement  []Statement
	FalseStatement []Statement
}

func (ifs IfStatement) statementNode() {}
